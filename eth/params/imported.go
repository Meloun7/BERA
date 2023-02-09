// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package params

import "github.com/ethereum/go-ethereum/params"

type (
	ChainConfig = params.ChainConfig
	Rules       = params.Rules
)

var (
	// `RefundQuotient` is the refund quotient parameter.
	RefundQuotient = params.RefundQuotient
	// `RefundQuotientEIP3529` is the refund quotient parameter for EIP-3529.
	RefundQuotientEIP3529 = params.RefundQuotientEIP3529
	// `TxAccessListAddressGas` is the cost of an address for a transaction with an access list.
	TxAccessListAddressGas = params.TxAccessListAddressGas
	// `TxAccessListAddressGasEIP2930` is the cost of an address for a transaction with an access
	// list.
	TxAccessListStorageKeyGas = params.TxAccessListStorageKeyGas
	// `TxDataNonZeroGasFrontier` is the cost of a non-zero byte of data for a transaction.
	TxDataNonZeroGasFrontier = params.TxDataNonZeroGasFrontier
	// `TxDataNonZeroGasEIP2028` is the cost of a non-zero byte of data for a transaction.
	TxDataNonZeroGasEIP2028 = params.TxDataNonZeroGasEIP2028
	// `TxDataZeroGas` is the cost of a zero byte of data or code for a transaction.
	TxDataZeroGas = params.TxDataZeroGas
	// `TxGasContractCreation` is the amount of gas that is refunded for a contract creation
	// transaction.
	TxGasContractCreation = params.TxGasContractCreation
	// `TxGas` is the amount of gas that is refunded for a transaction.
	TxGas = params.TxGas
)
