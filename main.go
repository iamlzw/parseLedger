package main

import (
	"bytes"
	"github.com/hyperledger/fabric/common/tools/protolator"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"flag"
	"fmt"
	"github.com/hyperledger/fabric/common/ledger/blkstorage"
	"github.com/hyperledger/fabric/common/ledger/util/leveldbhelper"
	"github.com/hyperledger/fabric/common/metrics/disabled"
	"github.com/lifegoeson/parseblockfiles/utils"
	"os"
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

	// 定义命令行参数
	//inputPath := flag.String("in", "", "Input Path")
	//outputPath := flag.String("out", "", "Output Path")

	chlName := flag.String("channelName", "", "通道名称")

	start := flag.Int64("start", 0, "区块结束的高度")

	end := flag.Int64("end", 0, "区块开始的高度")

	all := flag.Bool("all", false, "是否解析通道的所有区块")

	// 解析命令行参数
	flag.Parse()

	conf := &utils.Conf{BlockStorageDir: "chains", MaxBlockfileSize: 100000000000}
	indexConfig := &blkstorage.IndexConfig{AttrsToIndex: attrsToIndex}
	p := leveldbhelper.NewProvider(&leveldbhelper.Conf{DBPath: conf.GetIndexDir()})
	// create stats instance at provider level and pass to newFsBlockStore
	stats := utils.NewStats( &disabled.Provider{})
	provider := utils.FsBlockstoreProvider{Conf: conf, IndexConfig: indexConfig, LeveldbProvider: p, Stats: stats}

	blkStore, err := provider.OpenBlockStore(*chlName)

	if blkStore == nil{
		fmt.Errorf("获取blkStore为空")
		return
	}
	if err != nil  {
		fmt.Print("error")
	}
	//
	blockInfo, err := blkStore.GetBlockchainInfo()
	if err != nil{
		fmt.Print("error")
	}

	if blockInfo == nil{
		fmt.Errorf("获取blockInfo为空")
		return
	}

	//fmt.Println("当前通道区块高度: ",blockInfo.Height)
	var st,ed int64

	//是否解析通道所有的区块数据
	if !*all {
		st = *start
		ed = *end
	}else {
		st = 0
		ed = int64(blockInfo.Height)
	}

	//
	for i := st; i < ed ; i++ {
		//fmt.Println(blockInfo.Height)
		block, err := blkStore.RetrieveBlockByNumber(uint64(i))

		buf := new (bytes.Buffer)
		err = protolator.DeepMarshalJSON(buf, block)
		if err != nil {
			fmt.Errorf("malformed block contents: %s", err)
		}

		blockFilesPath := filepath.Join("blockfiles",*chlName)

		err = CreateDirIfNotExists(blockFilesPath)

		if err !=nil {
			fmt.Println("获取blockfiles路径失败")
			return
		}

		//fmt.Println(blockFilesPath)

		filename := blockFilesPath+"/"+*chlName+"_"+strconv.FormatInt(i,10)+".json"
		err = ioutil.WriteFile(filename,buf.Bytes(),0644)
		if err != nil{
			fmt.Println("write to file failure:",err)
		}
	}

}

func IsDirExists(path string) bool {
	// 获取文件或文件夹的信息
	info, err := os.Stat(path)

	// 如果发生错误，并且错误是因为文件或目录不存在
	if os.IsNotExist(err) {
		return false
	}

	// 如果没有发生错误，判断它是否是一个目录
	return info.IsDir()
}

func CreateDirIfNotExists(path string) error {
	if !IsDirExists(path) {
		// 如果目录不存在，则创建该目录
		err := os.MkdirAll(path, os.ModePerm) // os.ModePerm 表示 0777 权限
		if err != nil {
			return err
		}
		fmt.Printf("Directory %s created successfully\n", path)
	}
	return nil
}
