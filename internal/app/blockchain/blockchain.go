package blockchain

type Blockchain struct {
	Blocks []*Block
}

func NewBlockChain(blocks []*Block) *Blockchain {
	return &Blockchain{Blocks: blocks}
}

func (b *Blockchain) Add(data []byte) {
	
}
