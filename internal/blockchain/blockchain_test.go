package blockchain

import (
	"testing"
)

func TestBlockchain(t *testing.T) {
	bc := NewBlockchain()

	test := []string{
		"Send 1 BTC to Ivan",
		"Send 2 more BTC to Ivan",
	}

	for _, v := range test {
		bc.AddBlock(v)
	}

	for i := range test {
		if string(bc.blocks[i+1].Data) != test[i] {
			t.Fatalf("Expected: %s, Got: %s", test[i], string(bc.blocks[i].Data))
		}
	}
}
