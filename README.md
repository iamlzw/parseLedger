# parse-fabric-block-files
a tool to read and parse hyperledger fabric block files
解析fabric账本文件的工具

## 克隆仓库

```
# git clone https://github.com/iamlzw/parse-fabric-block-files.git
```

## 复制账本以及索引文件

```
# docker cp peer0.org1.example.com:/var/hyperledger/production/ledgersData/chains $PWD
#删除原来的账本以及索引文件
# rm -rf ${工具路径}/chains/*
将账本复制到工具相应目录下
# cp -r chains/* ${工具路径}/chains/
``` 
## 修改ledger id
```
# cd ${工具路径}/readLedger
# vim main.go
```
将channel id改为你相应的channel id
```
blkStore, err := provider.OpenBlockStore("mychannel")
```
## 执行

```
# cd ${工具路径}/readLedger
# go run main.go
```
解析出的区块数据在blockfiles目录下

