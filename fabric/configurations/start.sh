#!/bin/bash

echo "开始下载镜像..."
./download-dockerimages.sh
echo "下载镜像结束"

echo "开始生成证书..."
cryptogen generate --config ./crypto-config.yaml

echo "生成证书结束"

echo "开始生成创世块和通道..."
mkdir channel-artifacts
configtxgen -profile GenGenesis -outputBlock ./channel-artifacts/genesis.block
configtxgen -profile GenChannel -channelID mychannel -outputCreateChannelTx ./channel-artifacts/channel.tx
echo "生成创世块和通道结束"

echo "开始启动docker-compose..."
docker-compose -f ./docker-compose-cli.yaml up -d
echo "启动docker-compose结束"

sleep 5

echo "开始创建通道"
docker exec cli peer channel create -o orderer.example.com:7050 -c mychannel -f ./channel-artifacts/channel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
echo "创建通道结束"

echo "开始将当前peer节点加入通道"
docker exec cli peer channel join -b mychannel.block
echo "当前peer节点加入通道结束"

echo "开始安装链码"
docker exec cli peer chaincode install -n example02 -p github.com/hyperledger/fabric/examples/chaincode/go -v 1.0.0
echo "链码安装结束"

echo "开始初始化链码"
docker exec cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -c '{"Args":["init","a","200","b","300"]}' -n example02 -P "OR ('Org1.member','Org2.member')" -v 1.0.0 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
echo "链码初始化结束"