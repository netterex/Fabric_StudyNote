#!/bin/bash

#echo "开始下载镜像..."
##./download-dockerimages.sh
#echo "下载镜像结束"

#echo "开始生成证书..."
#cryptogen generate --config ./crypto-config.yaml
#echo "生成证书结束"
#
#echo "开始生成创世块和通道..."
#mkdir channel-artifacts
#configtxgen -profile GenGenesis -outputBlock ./channel-artifacts/genesis.block
#configtxgen -profile GenChannel -channelID policechannel -outputCreateChannelTx ./channel-artifacts/policechannel.tx
#configtxgen -profile GenChannel -channelID housemngchannel -outputCreateChannelTx ./channel-artifacts/housemngchannel.tx
#configtxgen -profile GenChannel -channelID creditchannel -outputCreateChannelTx ./channel-artifacts/creditchannel.tx
#echo "生成创世块和通道结束"
#
#echo "开始启动docker-compose..."
#docker-compose -f ./docker-compose-cli.yaml up -d
#echo "启动docker-compose结束"
#
#sleep 5
#
#echo "开始创建通道"
#docker exec cli_police peer channel create -o orderer.example.com:7050 -c policechannel -f ./channel-artifacts/policechannel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
#docker exec cli_housemng peer channel create -o orderer.example.com:7050 -c housemngchannel -f ./channel-artifacts/housemngchannel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
#docker exec cli_credit peer channel create -o orderer.example.com:7050 -c creditchannel -f ./channel-artifacts/creditchannel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
#echo "创建通道结束"
#
#echo "拷贝生成的block文件到宿主机"
#docker cp cli_police:/opt/gopath/src/github.com/hyperledger/fabric/peer/policechannel.block /home/netter/go/workspace/src/fabric-houserental/fabric-network
#docker cp cli_housemng:/opt/gopath/src/github.com/hyperledger/fabric/peer/housemngchannel.block /home/netter/go/workspace/src/fabric-houserental/fabric-network
#docker cp cli_credit:/opt/gopath/src/github.com/hyperledger/fabric/peer/creditchannel.block /home/netter/go/workspace/src/fabric-houserental/fabric-network
#
#echo "拷贝block文件到容器"
#docker cp /home/netter/go/workspace/src/fabric-houserental/fabric-network/policechannel.block cli_police1:/opt/gopath/src/github.com/hyperledger/fabric/peer
#docker cp /home/netter/go/workspace/src/fabric-houserental/fabric-network/housemngchannel.block cli_housemng1:/opt/gopath/src/github.com/hyperledger/fabric/peer
#docker cp /home/netter/go/workspace/src/fabric-houserental/fabric-network/creditchannel.block cli_credit1:/opt/gopath/src/github.com/hyperledger/fabric/peer
#echo "拷贝结束"
#
#
#echo "开始将当前peer节点加入通道"
#docker exec cli_police peer channel join -b policechannel.block
#docker exec cli_police1 peer channel join -b policechannel.block
#docker exec cli_housemng peer channel join -b housemngchannel.block
#docker exec cli_housemng1 peer channel join -b housemngchannel.block
#docker exec cli_credit peer channel join -b creditchannel.block
#docker exec cli_credit1 peer channel join -b creditchannel.block
#echo "当前peer节点加入通道结束"

echo "开始安装链码"
docker exec cli_police peer chaincode install -n policeCc -v 1.0 -p github.com/hyperledger/fabric/chaincode/go/police
docker exec cli_police1 peer chaincode install -n policeCc -v 1.0 -p github.com/hyperledger/fabric/chaincode/go/police
docker exec cli_housemng peer chaincode install -n housemngCc -v 1.0 -p github.com/hyperledger/fabric/chaincode/go/housemng
docker exec cli_housemng1 peer chaincode install -n housemngCc -v 1.0 -p github.com/hyperledger/fabric/chaincode/go/housemng
docker exec cli_credit peer chaincode install -n creditCc -v 1.0 -p github.com/hyperledger/fabric/chaincode/go/credit
docker exec cli_credit1 peer chaincode install -n creditCc -v 1.0 -p github.com/hyperledger/fabric/chaincode/go/credit
docker exec cli_credit peer chaincode install -n contractCc -v 1.0 -p github.com/hyperledger/fabric/chaincode/go/contract
docker exec cli_credit1 peer chaincode install -n contractCc -v 1.0 -p github.com/hyperledger/fabric/chaincode/go/contract
docker exec cli_credit peer chaincode install -n transactionCc -v 1.0 -p github.com/hyperledger/fabric/chaincode/go/transaction
docker exec cli_credit1 peer chaincode install -n transactionCc -v 1.0 -p github.com/hyperledger/fabric/chaincode/go/transaction
docker exec cli_credit peer chaincode install -n userCc -v 1.0 -p github.com/hyperledger/fabric/chaincode/go/user
docker exec cli_credit1 peer chaincode install -n userCc -v 1.0 -p github.com/hyperledger/fabric/chaincode/go/user
echo "链码安装结束"

# 只需初始化一个容器（即一个peer节点）即可
echo "开始初始化链码"
docker exec cli_police peer chaincode instantiate -o orderer.example.com:7050 -C policechannel -n policeCc -P "OR ('PoliceMSP.member')" -c '{"Args":["init","641111111111111111","zs","18","xxx"]}' -v 1.0 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
docker exec cli_housemng peer chaincode instantiate -o orderer.example.com:7050 -C housemngchannel -n housemngCc -P "OR ('HousemngMSP.member')" -c '{"Args":["init","00000001","zs","xxx","shangpin"]}' -v 1.0 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
docker exec cli_credit peer chaincode instantiate -o orderer.example.com:7050 -C creditchannel -n creditCc -P "OR ('CreditMSP.member')" -c '{"Args":["init","641111111111111111","A"]}' -v 1.0 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
docker exec cli_credit peer chaincode instantiate -o orderer.example.com:7050 -C creditchannel -n contractCc -P "OR ('CreditMSP.member')" -c '{"Args":["init","000001","641111111111111111","00000001","1523467a7befeee"]}' -v 1.0 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
docker exec cli_credit peer chaincode instantiate -o orderer.example.com:7050 -C creditchannel -n transactionCc -P "OR ('CreditMSP.member')" -c '{"Args":["init","0000000001","4","641111111111111112","641111111111111113","ls","ww","00000002","1000","2023-12-12 16:38:33","000002","test"]}' -v 1.0 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
docker exec cli_credit peer chaincode instantiate -o orderer.example.com:7050 -C creditchannel -n userCc -P "OR ('CreditMSP.member')" -c '{"Args":["init","4","user","123456","18","female"]}' -v 1.0 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
echo "链码初始化结束"
