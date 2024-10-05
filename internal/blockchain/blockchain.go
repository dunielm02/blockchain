package blockchain

import (
	"iter"
	"log"

	bolt "go.etcd.io/bbolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocksBucket"

type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

func (bc *Blockchain) AddBlock(data string) {
	var lastBlockHash []byte

	bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastBlockHash = b.Get([]byte{'l'})

		return nil
	})

	newBlock := NewBlock(data, lastBlockHash)

	bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		err := b.Put(newBlock.Hash, newBlock.Serialize())

		if err != nil {
			return err
		}

		err = b.Put([]byte("l"), newBlock.Hash)
		bc.tip = newBlock.Hash

		return err
	})

}

func (bc *Blockchain) GetByHash(hash []byte) *Block {
	var ret *Block

	bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		encoded := b.Get(hash)

		ret = DeserializeBlock(encoded)

		return nil
	})

	return ret
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, _ = tx.CreateBucket([]byte(blocksBucket))
			b.Put(genesis.Hash, genesis.Serialize())
			b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}

func (bc *Blockchain) Close() error {
	return bc.db.Close()
}

func (bc *Blockchain) Iterator() iter.Seq[*Block] {
	return func(yield func(*Block) bool) {
		currentBlock := bc.GetByHash(bc.tip)

		for string(currentBlock.Data) != "Genesis Block" {
			if !yield(currentBlock) {
				return
			}

			currentBlock = bc.GetByHash(currentBlock.PrevBlockHash)
		}
	}
}
