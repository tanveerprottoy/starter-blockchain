package blockchain

import (
	"bytes"

	"github.com/tanveerprottoy/starter-blockchain/pkg/cryptopkg"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func NewBlock(data []byte, prevHash []byte) *Block {
	b := &Block{[]byte{}, data, prevHash}
	b.DeriveHash()
	return b
}

func (b *Block) DeriveHash() {
	// i (info) bytes.Join together slice of bytes
	// bytes.Join first param accepts a two dimensional slice of byte
	// combine with an empty slice of byte
	i := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	// h (hash) generated
	h := cryptopkg.SHA256Sum(i)
	// the current block has can be set now
	// [:] operator makes a slice from an array
	// syntax for array to slice convert [startIdx:endIdx(excluding)]
	b.Hash = h[:]
}
