package blockchain

import (
	"fmt"
	"testing"
	"time"
)

func TestBlockchain(t *testing.T) {
	now := time.Now()

	bc := NewBlockchain()

	test := []string{
		"Send 1 BTC to Ivan",
		"Send 2 more BTC to Ivan",
	}

	for _, v := range test {
		bc.AddBlock(v)
	}

	bc.Close()

	bc = NewBlockchain()
	var i = len(test) - 1
	for block := range bc.Iterator() {
		if string(block.Data) != test[i] {
			t.Fatalf("expected: %s, got: %s\n", test[i], string(block.Data))
		}
		i--
	}

	fmt.Println(time.Since(now))
}
