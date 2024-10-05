package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var printChain = &cobra.Command{
	Use:   "printchain",
	Short: "Prints all the block in the blockchain.",
	Long:  `Prints all the block in the blockchain.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		for block := range bc.Iterator() {
			fmt.Println(string(block.Data))
		}
	},
}
