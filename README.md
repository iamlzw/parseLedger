# ReadLedger
用于解析hyperledger fabric peer或者orderer账本文件的工具

## 复制账本以及索引文件
```
# docker cp peer0.org1.example.com:/var/hyperledger/production/ledgersData $PWD
将账本复制到工具相应目录下
# cp -r ledgersData/chains ${工具路径}/
```
## 编译
```
# go build -o parse main.go
```
## 执行,--channelName 要解析的通道名称,--start 开始的区块高度,--end 结束的区块高度
```
$ ./parse --channelName mychannel --start 0 --end 6
```
解析出的区块数据在blockfiles目录下

