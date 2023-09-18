// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package core

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core/vm"

	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/types"
)

// ChainWriter defines methods that are used to perform state and block transitions.
type ChainWriter interface {
	// Prepare prepares the chain for a new block. This method is called before the first tx in
	// the block.
	Prepare(context.Context, uint64)
	// ProcessTransaction processes the given transaction and returns the receipt after applying
	// the state transition. This method is called for each tx in the block.
	ProcessTransaction(context.Context, *types.Transaction) (*ExecutionResult, error)
	// Finalize is called after the last tx in the block.
	Finalize(context.Context) error
	// SendTx sends the given transaction to the tx pool.
	SendTx(ctx context.Context, signedTx *types.Transaction) error
}

// =========================================================================
// Block Processing
// =========================================================================

// Prepare prepares the blockchain for processing a new block at the given height.
func (bc *blockchain) Prepare(ctx context.Context, number uint64) {
	// Prepare the State, Block, Configuration, Gas, and Historical plugins for the block.
	bc.sp.Prepare(ctx)
	bc.bp.Prepare(ctx)
	bc.cp.Prepare(ctx)
	bc.gp.Prepare(ctx)

	if bc.hp != nil {
		bc.hp.Prepare(ctx)
	}

	coinbase, timestamp := bc.bp.GetNewBlockMetadata(number)
	chainConfig := bc.Config()

	// Build the new block header.
	parent := bc.CurrentFinalBlock()
	if number >= 1 && parent == nil {
		parent = bc.GetHeaderByNumber(number - 1)
	}

	// Polaris does not set Ethereum state root (Root), mix hash (MixDigest), extra data (Extra),
	// and block nonce (Nonce) on the new header.
	header := &types.Header{
		// Used in Polaris.
		ParentHash: parent.Hash(),
		Coinbase:   coinbase,
		Number:     new(big.Int).SetUint64(number),
		GasLimit:   bc.gp.BlockGasLimit(),
		Time:       timestamp,
	}

	// TODO: Settable in PrepareProposal.
	// Set the extra field.
	if /*len(w.extra) != 0*/ true {
		header.Extra = nil
	}

	// Set the randomness field from the beacon chain if it's available.
	// TODO: Settable in PrepareProposal.
	if /*genParams.random != (common.Hash{})*/ true {
		// header.MixDigest = genParams.random
		header.MixDigest = common.Hash{}
	}

	// Apply EIP-1559.
	// TODO: Move to PrepareProposal.
	if chainConfig.IsLondon(header.Number) {
		header.BaseFee = eip1559.CalcBaseFee(chainConfig, parent)
		// On switchover.
		// TODO: implement.
		// if !chainConfig.IsLondon(parent.Number) {
		// 	parentGasLimit := parent.GasLimit * chainConfig.ElasticityMultiplier()
		// 	header.GasLimit = core.CalcGasLimit(parentGasLimit, bc.gp.BlockGasLimit())
		// }
	}

	// Apply EIP-4844, EIP-4788.
	// TODO: Move to PrepareProposal.
	if chainConfig.IsCancun(header.Number, header.Time) {
		var excessBlobGas uint64
		if chainConfig.IsCancun(parent.Number, parent.Time) {
			excessBlobGas = eip4844.CalcExcessBlobGas(*parent.ExcessBlobGas, *parent.BlobGasUsed)
		} else {
			// For the first post-fork block, both parent.data_gas_used and parent.excess_data_gas are evaluated as 0
			excessBlobGas = eip4844.CalcExcessBlobGas(0, 0)
		}
		header.BlobGasUsed = new(uint64)
		header.ExcessBlobGas = &excessBlobGas
		header.ParentBeaconRoot = &common.Hash{}
	}
	bc.logger.Info("preparing evm block", "seal_hash", header.Hash())

	// We update the base fee in the txpool to the next base fee.
	// TODO: Move to PrepareProposal.
	bc.tp.SetBaseFee(header.BaseFee)

	// Prepare the State Processor, StateDB and the EVM for the block.
	// TODO: Move to PrepareProposal.
	bc.processor.Prepare(
		bc.GetEVM(ctx, vm.TxContext{}, bc.statedb, header, bc.vmConfig),
		header,
	)
}

// ProcessTransaction processes the given transaction and returns the receipt.
func (bc *blockchain) ProcessTransaction(ctx context.Context, tx *types.Transaction) (*ExecutionResult, error) {
	bc.logger.Debug("processing evm transaction", "tx_hash", tx.Hash())

	// Reset the Gas and State plugins for the tx.
	bc.gp.Reset(ctx) // TODO: may not need this.
	bc.sp.Reset(ctx)

	return bc.processor.ProcessTransaction(ctx, tx)
}

// Finalize finalizes the current block.
func (bc *blockchain) Finalize(ctx context.Context) error {
	block, receipts, logs, err := bc.processor.Finalize(ctx)
	if err != nil {
		return err
	}

	blockHash, blockNum := block.Hash(), block.Number().Uint64()
	bc.logger.Info("finalizing evm block", "block_hash", blockHash.Hex(), "num_txs", len(receipts))

	// store the block header on the host chain
	err = bc.bp.StoreHeader(block.Header())
	if err != nil {
		bc.logger.Error("failed to store block header", "err", err)
		return err
	}

	// store the block, receipts, and txs on the host chain if historical plugin is supported
	if bc.hp != nil {
		if err = bc.hp.StoreBlock(block); err != nil {
			bc.logger.Error("failed to store block", "err", err)
			return err
		}
		if err = bc.hp.StoreReceipts(blockHash, receipts); err != nil {
			bc.logger.Error("failed to store receipts", "err", err)
			return err
		}
		if err = bc.hp.StoreTransactions(blockNum, blockHash, block.Transactions()); err != nil {
			bc.logger.Error("failed to store transactions", "err", err)
			return err
		}
	}

	// mark the current block, receipts, and logs
	if block != nil {
		bc.currentBlock.Store(block)
		bc.finalizedBlock.Store(block)

		// Todo: nuke these caches.
		bc.blockNumCache.Add(blockNum, block)
		bc.blockHashCache.Add(blockHash, block)

		// Todo: nuke these caches.
		for txIndex, tx := range block.Transactions() {
			bc.txLookupCache.Add(
				tx.Hash(),
				&types.TxLookupEntry{
					Tx:        tx,
					TxIndex:   uint64(txIndex),
					BlockNum:  blockNum,
					BlockHash: blockHash,
				},
			)
		}
	}
	if receipts != nil {
		bc.currentReceipts.Store(receipts)
		// Todo: nuke this cache.
		bc.receiptsCache.Add(blockHash, receipts)
	}
	if logs != nil {
		bc.pendingLogsFeed.Send(logs)
		bc.currentLogs.Store(logs)
		if len(logs) > 0 {
			bc.logsFeed.Send(logs)
		}
	}

	// Send chain events.
	bc.chainFeed.Send(ChainEvent{Block: block, Hash: blockHash, Logs: logs})
	bc.chainHeadFeed.Send(ChainHeadEvent{Block: block})

	return nil
}

func (bc *blockchain) SendTx(_ context.Context, signedTx *types.Transaction) error {
	return bc.tp.SendTx(signedTx)
}
