# parseLedger
用于解析hyperledger fabric peer或者orderer账本文件的工具

## 编译
```
# go build -o parse main.go
```

## 复制账本以及索引文件,需要确保账本的目录结构为chains/{chains,index}
### peer
```
$ docker cp peer0.org1.example.com:/var/hyperledger/production/ledgersData/chains $PWD
将账本复制到工具相应目录下
$ cp -r chains ${工具路径}/
$ ls chains
chains
index
```
### orderer
```
$ docker cp orderer.example.com:/var/hyperledger/production/orderer $PWD
在当前目录下创建chains目录
mkdir chains
将账本复制到工具相应目录下
$ cp -r orderer/* ${工具路径}/chains/
$ ls chains
chains
index
```


## 执行,--channel 要解析的通道名称,--start 开始的区块高度,--end 结束的区块高度
```
$ ./parse --channel mychannel --start 0 --end 6
```
解析出的区块数据在blockfiles目录下

