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

package debug_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"pkg.berachain.dev/stargazer/lib/errors/debug"
)

func TestDebug(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "lib/errors/debug")
}

func Hello() error {
	return nil
}

type GoodBye struct{}

func (g GoodBye) GoodBye() error {
	return nil
}

func (g *GoodBye) GoodByePtr() error {
	return nil
}

type GoodBye2 struct{}

func (g GoodBye2) GoodBye2() error {
	return nil
}

var _ = Describe("TestFnName", func() {
	It("should return the name of the function", func() {
		Expect(debug.GetFnName(Hello)).Should(Equal("Hello"))
	})

	It("should return the name of the function for struct functions", func() {
		gb := GoodBye{}
		Expect(debug.GetFnName(gb.GoodBye)).Should(Equal("GoodBye-fm"))
		Expect(debug.GetFnName(gb.GoodByePtr)).Should(Equal("GoodByePtr-fm"))
		Expect(debug.GetFnName(GoodBye2{}.GoodBye2)).Should(Equal("GoodBye2-fm"))
	})
})
