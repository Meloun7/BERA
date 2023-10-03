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

package evm

import (
	"encoding/json"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"pkg.berachain.dev/polaris/eth/core"
)

// DefaultGenesis returns default genesis state as raw bytes for the evm
// module.
func (AppModuleBasic) DefaultGenesis(_ codec.JSONCodec) json.RawMessage {
	ethGen := core.DefaultGenesis
	rawGenesis, err := ethGen.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return rawGenesis
}

// ValidateGenesis performs genesis state validation for the evm module.
func (AppModuleBasic) ValidateGenesis(_ codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	ethGen := new(core.Genesis)
	if err := ethGen.UnmarshalJSON(bz); err != nil {
		return err
	}

	// TODO: this code actually needs to be removed as it will cause issues.
	// for address, account := range ethGen.Alloc {
	// 	if ethGen.Config.IsEIP155(big.NewInt(0)) && account.Code != nil && account.Nonce == 0 {
	// 		// NOTE: EIP 161 was released at the same block as EIP 155.
	// 		return fmt.Errorf(
	// 			"EIP-161 requires an account with code (%s) to have nonce of at least 1, given (0)",
	// 			address.Hex(),
	// 		)
	// 	}
	// }

	return nil
}

// InitGenesis performs genesis initialization for the evm module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, _ codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var ethGen core.Genesis
	if err := ethGen.UnmarshalJSON(data); err != nil {
		panic(err)
	}

	if err := am.keeper.InitGenesis(ctx, &ethGen); err != nil {
		panic(err)
	}
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the evm
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, _ codec.JSONCodec) json.RawMessage {
	ethGen := am.keeper.ExportGenesis(ctx)
	ethGenBz, err := ethGen.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return ethGenBz
}
