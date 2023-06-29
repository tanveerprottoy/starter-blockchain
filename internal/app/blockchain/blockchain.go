package blockchain

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain(blocks []*Block) *BlockChain {
	return &BlockChain{Blocks: blocks}
}

func (b *BlockChain) Add(data []byte) {
	
}
