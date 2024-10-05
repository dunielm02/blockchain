package main

import (
	"github.com/spf13/cobra"
)

var addBlock = &cobra.Command{
	Use:   "addblock",
	Short: "Adds a block to the blockchain.",
	Long:  `Adds a block to the blockchain.`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		bc.AddBlock(args[0])
	},
}
