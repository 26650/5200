package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// 이런게 있고, 앞으로 많아질거다
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// 나중엔 더 나은걸로 할거다 지금은 그냥 블록을 참조해서 갖고 있다
type BlockChain struct {
	blocks []*Block
}

// derive -> hash는 data, prevhash 등 여러가지를 합쳐서 만들어진다
func (b *Block) DeriveHash() {
	//data, prevhash byte를 합치기 위해서 join한다.
	//b.Data, b.PrevHash, 빈 바이트 슬라이스(구분자)를 요소로 가지는 슬라이스를 생성
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	//sha256은 너무 단순한 알고리즘. 그래서 앞으로 바꿀 것
	hash := sha256.Sum256(info)
	//sum256은 array고, byte[]는 바이트 슬라이스
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// 그냥 링크드 리스트다.
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

// 첫번째 블록은 만들어야 한다.
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// 그냥 첫번째 블록을 넣은, 링크드 리스트를 반환한다. 이게 블록체인
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block After genesis")
	chain.AddBlock("Second Block After genesis")
	chain.AddBlock("Third Block After genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Prev Hash %x\n", block.PrevHash)
		fmt.Printf("Prev Hash %s\n", block.Data)
		fmt.Printf("hash: %x\n", block.Hash)
	}
}
