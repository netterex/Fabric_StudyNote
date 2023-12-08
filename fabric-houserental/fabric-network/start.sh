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
configtxgen -profile GenChannel -channelID policechannel -outputCreateChannelTx ./channel-artifacts/policechannel.tx
configtxgen -profile GenChannel -channelID housemngchannel -outputCreateChannelTx ./channel-artifacts/housemngchannel.tx
configtxgen -profile GenChannel -channelID creditchannel -outputCreateChannelTx ./channel-artifacts/creditchannel.tx
echo "生成创世块和通道结束"

echo "开始启动docker-compose..."
docker-compose -f ./docker-compose-cli.yaml up -d
echo "启动docker-compose结束"

sleep 5

echo "开始创建通道"
docker exec cli_police peer channel create -o orderer.example.com:7050 -c policechannel -f ./channel-artifacts/policechannel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
docker exec cli_housemng peer channel create -o orderer.example.com:7050 -c housemngchannel -f ./channel-artifacts/housemngchannel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
docker exec cli_ciredit peer channel create -o orderer.example.com:7050 -c creditchannel -f ./channel-artifacts/creditchannel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
echo "创建通道结束"

echo "开始将当前peer节点加入通道"
docker exec cli_police peer channel join -b policechannel.block
docker exec cli_housemng peer channel join -b housemngchannel.block
docker exec cli_credit peer channel join -b creditchannel.block
echo "当前peer节点加入通道结束"

# todo 切换peer1节点并加入通道

#echo "开始安装链码"
#docker exec cli_police peer chaincode install -n police -p github.com/hyperledger/fabric/chaincode/go/police -v 1.0
#docker exec cli_housemng peer chaincode install -n housemng -p github.com/hyperledger/fabric/chaincode/go/housemng -v 1.0
#docker exec cli_credit peer chaincode install -n credit -p github.com/hyperledger/fabric/chaincode/go/credit -v 1.0
#echo "链码安装结束"
#
#echo "开始初始化链码"
#docker exec cli_police peer chaincode instantiate -o orderer.example.com:7050 -C policechannel -c '{"Args":["init","a","200","b","300"]}' -n police -P "OR ('PoliceOrg.member')" -v 1.0 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
#docker exec cli_housemng peer chaincode instantiate -o orderer.example.com:7050 -C housemngchannel -c '{"Args":["init","a","200","b","300"]}' -n housemng -P "OR ('HousemngOrg.member')" -v 1.0 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
#docker exec cli_credit peer chaincode instantiate -o orderer.example.com:7050 -C creditchannel -c '{"Args":["init","a","200","b","300"]}' -n credit -P "OR ('CreditOrg.member')" -v 1.0 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
#echo "链码初始化结束"