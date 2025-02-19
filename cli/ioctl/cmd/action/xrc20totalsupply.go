// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package action

import (
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
)

// xrc20TotalSupplyCmd represents total supply of the contract
var xrc20TotalSupplyCmd = &cobra.Command{
	Use:   "totalSupply -c ALIAS|CONTRACT_ADDRESS",
	Short: "Get total supply",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		bytecode, err := xrc20ABI.Pack("totalSupply")
		if err != nil {
			return err
		}
		contract, err := xrc20Contract()
		if err != nil {
			return err
		}
		output, err := read(contract, bytecode)
		if err == nil {
			fmt.Println("Raw output:", output)
			decimal, _ := new(big.Int).SetString(output, 16)
			fmt.Printf("Output in decimal: %d\n", decimal)
		}
		return err
	},
}
