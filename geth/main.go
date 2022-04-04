package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main() {
	// create an eth client
	url := fmt.Sprintf("%s://%s:%d", "http", "192.168.43.121", 8545)
	client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}
	// create an indexer
	indexer := &indexer{
		client,
	}
	// start the indexer
	indexer.Start(3611553, 3611553)
}

var ctx = context.TODO()

type indexer struct {
	client *ethclient.Client
}

// 宣告一個indexer方法
func (indexer *indexer) Start(from int64, to int64) error {
	start := big.NewInt(from)
	end := big.NewInt(to)

	for i := new(big.Int).Set(start); i.Cmp(end) <= 0; i.Add(i, big.NewInt(1)) {
		// get block by number
		block, err := indexer.client.BlockByNumber(ctx, i)
		if err != nil {
			return err
		}
		fmt.Println("block hash data: " + block.Hash().String())

		// parse block header
		//blockHeader := indexer.ParseBlockHeader(block)
		//fmt.Println("block data: " + block.TxHash().String())

		// parse transactions
		for _, tx := range block.Transactions() {
			fmt.Println("tx data: " + tx.Hash().String())
			//transaction, receipt, err := indexer.ParseTransaction(tx, block.Number())
			//if err != nil {
			//	return err
			//}
		}
	}
	return nil
}
