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

package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// WrappedEthereumTransaction defines a Cosmos SDK message for Ethereum transactions.
var _ sdk.Msg = (*WrappedEthereumTransaction)(nil)

// WrapTx sets the transaction data from an `coretypes.Transaction`.
func WrapTx(tx *coretypes.Transaction) (*WrappedEthereumTransaction, error) {
	bz, err := tx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to wrap transaction: %w", err)
	}

	return &WrappedEthereumTransaction{
		Data: bz,
	}, nil
}

// AsTransaction extracts the transaction as an `coretypes.Transaction`.
func (etr *WrappedEthereumTransaction) UnwrapTx() *coretypes.Transaction {
	tx := new(coretypes.Transaction)
	if err := tx.UnmarshalBinary(etr.Data); err != nil {
		return nil
	}
	return tx
}
