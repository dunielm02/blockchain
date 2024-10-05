package main

import (
	"fmt"

	"github.com/dunielm02/blockchain/internal/blockchain"
	"github.com/spf13/cobra"
)

var printChain = &cobra.Command{
	Use:   "printchain",
	Short: "Prints all the block in the blockchain.",
	Long:  `Prints all the block in the blockchain.`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		for block := range bc.Iterator() {
			fmt.Printf("%-*s %x\n", 11, "Prev. Hash:", block.PrevBlockHash)
			fmt.Printf("%-*s %s\n", 11, "Data:", string(block.Data))
			fmt.Printf("%-*s %x\n", 11, "Hash:", block.Hash)
			pow := blockchain.NewProofOfWork(block)
			fmt.Printf("%-*s %v\n", 11, "PoW:", pow.Validate())
			fmt.Println()
		}
		
	},
}
