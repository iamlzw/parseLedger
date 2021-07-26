package main

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/common/ledger/blkstorage"
	"github.com/hyperledger/fabric/common/ledger/util/leveldbhelper"
	"github.com/hyperledger/fabric/common/metrics/disabled"
	"github.com/hyperledger/fabric/common/tools/protolator"
	"io/ioutil"
	"strconv"
)

var attrsToIndex = []blkstorage.IndexableAttr{
	blkstorage.IndexableAttrBlockHash,
	blkstorage.IndexableAttrBlockNum,
	blkstorage.IndexableAttrTxID,
	blkstorage.IndexableAttrBlockNumTranNum,
	blkstorage.IndexableAttrBlockTxID,
	blkstorage.IndexableAttrTxValidationCode,
}
func main(){
	conf := &Conf{"chains",100000000000}
	indexConfig := &blkstorage.IndexConfig{AttrsToIndex: attrsToIndex}
	p := leveldbhelper.NewProvider(&leveldbhelper.Conf{DBPath: conf.getIndexDir()})
	// create stats instance at provider level and pass to newFsBlockStore
	stats := newStats( &disabled.Provider{})
	provider := &FsBlockstoreProvider{conf, indexConfig, p, stats}

	blkStore, err := provider.OpenBlockStore("mychannel")
	if err != nil  {
		fmt.Print("error")
	}

	blockInfo, err := blkStore.GetBlockchainInfo()
	if err != nil{
		fmt.Print("error")
	}

	for i := uint64(0) ; i < blockInfo.Height ; i++ {
		fmt.Println(blockInfo.Height)
		block, err := blkStore.RetrieveBlockByNumber(i)

		buf := new (bytes.Buffer)
		err = protolator.DeepMarshalJSON(buf, block)
		if err != nil {
			fmt.Errorf("malformed block contents: %s", err)
		}
		filename := "blockfiles/channel2_"+strconv.FormatInt(int64(i),10)+".json"
		err = ioutil.WriteFile(filename,buf.Bytes(),0644)
		if err != nil{
			fmt.Println("write to file failure:",err)
		}
	}

}
