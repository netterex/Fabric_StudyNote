#  1 fabric环境搭建

## 1.1 python2.7安装

安装net-tools，配置ip地址。

```shell
	sudo apt-get install net-tools
	ifconfig
```

更新apt-get。

```shell
	sudo apt-get update
```

安装python2.7。

```shell
	sudo apt-get install python2.7
```

安装pip。

```shell
	sudo apt-get install python-pip
```

更新pip。

```shell
	sudo pip install --upgrade pip
```

## 1.2 go安装

安装vim。

```shell
	sudo apt-get install vim
```

安装go，在路径 /home/用户/ 下创建go文件夹，将对应版本的go解压在本目录中，如1.15.11。

```shell
	cd /home/用户/
	mkdir go
	cd go/
	# wget url 从网上下载文件
	# tar命令
	# -c 建立打包文件
    # -t 查看打包文件的内容含有哪些文件
    # -x 解打包或解压缩的功能，可以搭配-C(大写)在特定目录解开
    # -j 通过bzip2的支持进行压缩/解压缩，此时文件最好为*.tar.bz2
    # -z 通过gzip的支持进行压缩/解压缩，此时文件最好为*.tar.gz
    # -v 在压缩/解压缩的过程中，将正在处理的文件名显示出来
    # -f 后面跟处理后文件的全名称（路径+文件名+后缀名）
    # -C 这个选项用在解压缩，若要在特定目录解压缩，可以使用这个选项
    # -p 保留备份数据的原本权限与属性，常用于备份(-c)重要的配置文件
	# tar -xvf file解压
	# tar -cvf file压缩
	tar -xvf go1.15.11.linux-amd64.tar.gz
```

配置go环境变量。

```shell
	# vim命令
	# i 进入输入模式
	# x 删除当前光标所在处的字符
	# a 进入插入模式，在光标一个位置开始输入文本
	# o 在当前行的下方插入一个新行，并进入插入模式
	# dd 删除当前行
	# yy 复制当前行
	# u 撤销上一次操作
	# esc 返回普通模式
	# :wq 保存并退出编辑器
	# :q! 强制退出编辑器，不保存修改
	vim ~/.bashrc
	export GOROOT=/home/用户/go/go
	export GOPATH=/home/用户/go/workspace
	export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
	source ~/.bashrc
	go version
	go env
```

## 1.3 docker安装

安装git，用来拉取fabric代码。

```shell
	sudo apt-get install git
```

安装docker。

```shell
	sudo apt-get install docker
	sudo apt-get install docker.io
	# 验证，会提示权限不够
	docker version
```

创建docker用户组，并添加当前用户。

```shell
	sudo groupadd docker
	sudo gpasswd -a 用户 docker
	# 重启docker服务
	systemctl restart docker
	# 切换至docker用户组
	newgrp docker
	# 验证
	docker version
```

从[Releases · docker/compose (github.com)](https://github.com/docker/compose/releases)下载docker-compose的包到 /home/用户 路径下。

```shell
	sudo mv ~/docker-compose /usr/local/bin/docker-compose
	# 修改权限
	chmod +x /usr/local/bin/docker-compose
```

## 1.4 fabric安装

在GOPATH下新建一个目录，用于存放fabric有关的源码。

```shell
	# bin 可执行文件
	# pkg 编译后文件
	# src 源码
	mkdir bin pkg src
```

进入 src 目录中，通过git命令拉取fabric项目源码。

```shell
	cd src
	mkdir github.com
	cd github.com
	mkdir hyperledger
	cd hyperledger
	git clone https://github.com/hyperledger/fabric.git
```

切换分支。

```shell
	cd fabric
	git checkout v1.0.0
```

fabric的docker镜像下载。

```shell
	# examples是1.0分支下才有的
	cd examples/e2e_cli
	# 批量下载镜像，可能下载不完整或版本不对
	source download-dockerimages.sh -c x86_64-1.0.0 -f x86_64-1.0.0
	# 逐个下载镜像
	docker pull hyperledger/fabric-tools:x86_64-1.0.0
	docker pull hyperledger/fabric-orderer:x86_64-1.0.0
	docker pull hyperledger/fabric-peer:x86_64-1.0.0
	docker pull hyperledger/fabric-couchdb:x86_64-1.0.0
	docker pull hyperledger/fabric-kafka:x86_64-1.0.0
	docker pull hyperledger/fabric-ca:x86_64-1.0.0
	docker pull hyperledger/fabric-ccenv:x86_64-1.0.0
	docker pull hyperledger/fabric-baseimage:x86_64-0.4.7
	docker pull hyperledger/fabric-baseos:x86_64-0.4.7
	docker pull hyperledger/fabric-javaenv:x86_64-1.0.0
	docker pull hyperledger/fabric-zookeeper:x86_64-1.0.0
	docker pull hyperledger/fabric-membersrvc:latest
	# 修改标签为latest
	docker tag hyperledger/fabric-tools:x86_64-1.0.0 hyperledger/fabric-tools
	docker tag hyperledger/fabric-orderer:x86_64-1.0.0 hyperledger/fabric-orderer
	docker tag hyperledger/fabric-peer:x86_64-1.0.0 hyperledger/fabric-peer
	docker tag hyperledger/fabric-couchdb:x86_64-1.0.0 hyperledger/fabric-couchdb
	docker tag hyperledger/fabric-kafka:x86_64-1.0.0 hyperledger/fabric-kafka
	docker tag hyperledger/fabric-ca:x86_64-1.0.0 hyperledger/fabric-ca
	docker tag hyperledger/fabric-ccenv:x86_64-1.0.0 hyperledger/fabric-ccenv
	docker tag hyperledger/fabric-baseimage:x86_64-0.4.7 hyperledger/fabric-baseimage
	docker tag hyperledger/fabric-baseos:x86_64-0.4.7 hyperledger/fabric-baseos
	docker tag hyperledger/fabric-javaenv:x86_64-1.0.0 hyperledger/fabric-javaenv
	docker tag hyperledger/fabric-zookeeper:x86_64-1.0.0 hyperledger/fabric-zookeeper
	# 验证
	docker images
```

## 1.5 验证

需要将fabric项目放在 github.com/hyperledger/ 下。

```shell
	cd src/github.com/hyperledger/fabric/examples/e2e_cli
	# 启动fabric网络
	./network_setup.sh up
```

查看是否生成相关配置文件，并复制文件到 workspace/bin/ 路径中。

```shell
	cd src/github.com/hyperledger/fabric/release/linux-amd64/bin
	ll
	cp configtxgen /home/用户/go/workspace/bin
	cp configtxlator /home/用户/go/workspace/bin
	cp cryptogen /home/用户/go/workspace/bin
	cp orderer /home/用户/go/workspace/bin
	cp peer /home/用户/go/workspace/bin
	# 给 bin 路径权限
	chmod 775 -R /home/用户/go/workspace/bin
```

# 2 fabric项目结构

## 2.1 文件结构

1. **bccsp**：加密标准以及算法的实现，包括加密、签名、验签服务。
2. **bddtests**：bdd测试框架相关的测试。
3. **common**：通用功能以及一些通用的代码实现，包括日志、错误、工具包等，主要包括项目全局的功能性代码。
4. **core**：核心代码模块，其中包括权限控制、chaincode模块、committer、endorser、ledger、policy等核心功能的代码实现。
5. **discovery**：为客户端程序提供服务发现的功能。
6. **docs**：文档。
7. **events**：为客户端提供事件订阅的功能。
8. **examples**：案例。
9. **gossip**：信息传播的模块，为fabric在节点间达成最终一致性。
10. **idemix**：零知识证明，无需提供私有数据即可证明，包括用户（User）、发行者（Issuer）、验证者（Verifier）。
11. **integration**：代码集成。
12. **gotools**：用于编译fabric。
13. **msp**：提供成员服务。
14. **orderer**：排序节点模块。
15. **peer**：peer节点。
16. **proposals**：存放相关提案。
17. **protos**：存放Protocol buffer消息。
18. **release_notes**：各个版本的更新日志。
19. **sampleconfig**：相关样例配置文件。
20. **scripts**：存放相关脚本文件。
21. **unit-test**：单元测试（testenv）。
22. **swagger**：接口文档。
23. **tools**：工具。
24. **vagrant**：创建虚拟机，自动化配置和安装开发环境。

## 2.2 核心代码

1. **bccsp包**：实现对加解密算法和机制的支持。
2. **common包**：一些通用的模块。
3. **core包**：包含核心代码。
4. **events包**：为客户端提供事件订阅的功能。
5. **examples包**：示例。
6. **gossip包**：信息传播的模块。
7. **msp包**：提供成员服务。
8. **order包**：order服务相关的入口和框架代码。
9. **peer包**：peer的入口和框架代码。
10. **protos包**：包括各种协议和消息的protobuf定义文件和生成的go文件。

## 2.3 核心模块

**系统模块**：会以守护进程的方式在后台运行，不会中断。

**工具模块**：负责证书文件、区块链创世块、通道创世块等文件的生成，不参与系统的运行。

- **peer**：系统模块，peer节点，负责存储区块链数据，维护链码。
- **order**：系统模块，排序、打包交易，并提交给peer节点。
- **cryptogen**：工具模块，组织和证书生成模块。
- **configtxgen**：工具模块，区块和交易生成模块。
- **configtxlator**：工具模块，区块和交易解析模块。

# 3 配置文件和生成证书

新建一个项目目录。

```shell
	cd src
	mkdir hyperledger
	cd hyperledger
	cryptogen showtemplate > crypto-config.yaml
```

## 3.1 yaml格式介绍

yaml

- 大小写敏感。
- ‘#’表示注释。
- 不允许使用tab键，只能使用空格键。

## 3.2 配置信息

crypto-config.yaml 文件

```yaml
OrdererOrgs: # order组织配置项，有多个，属于数据
  - Name: Orderer # 组织名称
    Domain: example.com # 根域名
    Specs:
      - Hostname: orderer # 二级域名，orderer.example.com
PeerOrgs: # peer组织配置项
  - Name: Org1 # 组织名称
    Domain: org1.example.com
    EnableNode0Us: # 是否支持node.js
    Template: # 模板
      Count: 1 # 生成的peer节点数量，Hostname: peer0,peer1,peer2等，如peer0.org1.example.com
    Users: # 创建的普通用户，管理员会自动生成
      Count: 1 # 用户数
  - Name: Org2
    Domain: org2.example.com
    Template:
      Count: 1
    Users:
      Count: 1
```

Specs和Template的区别：

- Specs可以指定二级域名，Template会自动在上级域名前面按顺序加上peer0、peer1这种二级域名。
- 可以互换使用，可以同时存在。

## 3.3 生成证书

根据 config.yaml 生成证书。

```shell
	cd src/hyperledger/
	# 此时会生成一个 crypto-config 目录，包含密钥材料
	cryptogen generate --config crypto-config.yaml
```

# 4 创世块及通道文件的生成

## 4.1 配置信息

configtx.yaml 文件

```yaml
Organizations: # 组织，下面有排序节点和peer节点
    # orderer 组织，排序节点 
    - &OrdererOrg # 最好与cryptogen生成的排序节点名称一致
      Name: OrdererMSP # 排序节点名称
      ID: OrdererMSP # 排序节点ID，orderer.yaml的 LocalMSPID 必须与该ID相同
      # MSP的路径
      MSPDir: crypto-config/ordererOrganizations/example.com/msp

    # peer 组织， peer节点，可能有多个
    - &Org1 # 最好与cryptogen生成的peer节点名称一致
      Name: Org1MSP
      ID: Org1MSP
      MSPDir: crypto-config/peerOrganizations/org1.example.com/msp
      AnchorPeers:  # 锚节点
        - Host: peer0.org1.example.com # 节点域名，只能有一个锚节点，组织下任意一个节点选为锚节点
          Port: 7051 # 默认端口

    - &Org2
      Name: Org2MSP
      ID: Org2MSP
      MSPDir: crypto-config/peerOrganizations/org2.example.com/msp
      AnchorPeers: 
        - Host: peer0.org2.example.com
          Port: 7051
          
Orderer: &OrdererDefaults # 默认即可
    OrdererType: solo # 共识机制，1.0.0只包含solo和kafka

    Addresses: #排序节点的域名
        - orderer.example.com:7050
    BatchTimeout: 2s # 批次，多长时间产生新的区块
    BatchSize: # 批次大小
        MaxMessageCount: 100 # 交易的最大数量，达到后产生新的区块，100条左右
        AbsoluteMaxBytes: 64 MB # 数据量达到后产生新的区块，32或64MB左右
        PreferredMaxBytes: 512 KB # 建议的交易数值大小，512KB即可
    MaxChannels: 0 # 最大通道数
    Kafka: # kafka信息
        Brokers:
            - 127.0.0.1:9092
    Organizations:
    
Profiles:
    # 创世块命令的配置
    # 命令名称GenGennesis和GenChannel可以更改
    GenGenesis:
      Orderer: # Orderer配置项
        <<: *OrdererDefaults # &OrdererDefaults
        Organizations: # orderer组织
          - *OrdererOrg # orderer组织引用
      
      Consortiums: # 联盟
        SampleConsortium: 
          Organizations: # Peer节点组织
            - *Org1
            - *Org2

    # 通道命令的配置
    GenChannel:
      # 与上面Consortiums中的SampleConsortium保持一致
      Consortium: SampleConsortium
      Application:
        <<: *ApplicationDefaults
        Organizations: # peer节点
          - *Org1
          - *Org2
```

注意：

- 节点的id和name最好保持一致。
- 锚节点只能有一个。
- key: value 中，: 和 value 之间有空格
  - 如 Host: peer0.org1.example.com
         Port: 7051
  - Host 和 Port 应对齐。

## 4.2 生成创世块和通道文件

将编辑好的 configtx.yaml 文件上传到 crypto-config.yaml 文件同目录下。

在该目录中执行下述命令。

```shell
	# GenGenesis 是 configtx.yaml 中 Profiles 下的创世块命令
	configtxgen -profile GenGenesis -outputBlock ./genesis.block
```

进入存放 configtx.yaml 的路径中。

```shell
	# GenChannel 是 configtx.yaml 中 Profiles 下的通道命令
	# -channelID 后的参数可以自定义，但必须为小写
	configtxgen -profile GenChannel -channelID mychannel -outputCreateChannelTx ./channel.tx
```

在 configtx.yaml 文件所在目录中新建 channel-artifacts 目录，并将两文件移入其中。

```shell
	mkdir channel-artifacts
	mv channel.tx genesis.block channel-artifacts/
```

## 4.3 更新锚节点

进入存放 configtx.yaml 的路径中。

```shell
	# --channelID 后的参数为所属的通道
	# -asOrg 后的参数为peer节点所设置的组织名字Name
	configtxgen -profile GenChannel --channelID mychannel -asOrg Org1MSP -outputAnchorPeersUpdate Org1AnchorUpdate.tx
	configtxgen -profile GenChannel --channelID mychannel -asOrg Org2MSP -outputAnchorPeersUpdate Org2AnchorUpdate.tx
```

# 5 docker-compose配置

## 5.1 docker-compose-cli文件

```yaml
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#
version: '2'
services:
  orderer.example.com:
    extends:
      file:   base/docker-compose-base.yaml
      service: orderer.example.com
    container_name: orderer.example.com
  peer0.org1.example.com:
    container_name: peer0.org1.example.com
    extends:
      file:  base/docker-compose-base.yaml
      service: peer0.org1.example.com
  peer0.org2.example.com:
    container_name: peer0.org2.example.com
    extends:
      file:  base/docker-compose-base.yaml
      service: peer0.org2.example.com

  cli:
    container_name: cli
    image: hyperledger/fabric-tools:latest
    tty: true
    environment:
      - GODEBUG=netdns=go
      - GOPATH=/opt/gopath
      - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
      # 日志级别：critical、error、warning、notice、info、debug
      - CORE_LOGGING_LEVEL=INFO
      - CORE_PEER_ID=org1peer0
      # 要连接的peer节点
      - CORE_PEER_ADDRESS=peer0.org1.example.com:7051
      # peer组织ID
      - CORE_PEER_LOCALMSPID=Org1MSP
      # 通信时是否使用tls加密
      - CORE_PEER_TLS_ENABLED=true
      # tls证书文件
      - CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.crt
      # tls私钥文件
      - CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.key
      # 根证书文件
      - CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
      # 指定客户端的身份、管理员身份目录
      - CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
    command: /bin/bash #-c './scripts/script.sh ${CHANNEL_NAME}; sleep $TIMEOUT'
    volumes:
        - /var/run/:/host/var/run/
        - ./chaincode/go/:/opt/gopath/src/github.com/hyperledger/fabric/examples/chaincode/go
        - ./crypto-config:/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/
        #- ./scripts:/opt/gopath/src/github.com/hyperledger/fabric/peer/scripts/
        - ./channel-artifacts:/opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts
    depends_on:
      - orderer.example.com
      - peer0.org1.example.com
      - peer0.org2.example.com
```

## 5.2 docker-compose-base文件

```yaml
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#
version: '2'
services:
  orderer.example.com:
    container_name: orderer.example.com
    image: hyperledger/fabric-orderer:latest
    environment:
      - ORDERER_GENERAL_LOGLEVEL=INFO
      - ORDERER_GENERAL_LISTENADDRESS=0.0.0.0
      - ORDERER_GENERAL_GENESISMETHOD=file
      # 创世块对应的文件，不是指本地的，而是docker中的，与下面的volumes进行映射相关联
      - ORDERER_GENERAL_GENESISFILE=/var/hyperledger/orderer/orderer.genesis.block
      - ORDERER_GENERAL_LOCALMSPID=OrdererMSP
      - ORDERER_GENERAL_LOCALMSPDIR=/var/hyperledger/orderer/msp
      # enabled TLS
      - ORDERER_GENERAL_TLS_ENABLED=true
      - ORDERER_GENERAL_TLS_PRIVATEKEY=/var/hyperledger/orderer/tls/server.key
      - ORDERER_GENERAL_TLS_CERTIFICATE=/var/hyperledger/orderer/tls/server.crt
      - ORDERER_GENERAL_TLS_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric
    command: orderer
    volumes:
    - ../channel-artifacts/genesis.block:/var/hyperledger/orderer/orderer.genesis.block
    - ../crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp:/var/hyperledger/orderer/msp
    - ../crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/tls/:/var/hyperledger/orderer/tls
    ports:
      - 7050:7050
  peer0.org1.example.com:
    container_name: peer0.org1.example.com
    extends:
      file: peer-base.yaml
      service: peer-base
    environment:
      - CORE_PEER_ID=peer0.org1.example.com
      - CORE_PEER_ADDRESS=peer0.org1.example.com:7051
      - CORE_PEER_CHAINCODELISTENADDRESS=peer0.org1.example.com:7052
      - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer0.org1.example.com:7051
      - CORE_PEER_LOCALMSPID=Org1MSP
    volumes:
        - /var/run/:/host/var/run/
        - ../crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/msp:/etc/hyperledger/fabric/msp
        - ../crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls:/etc/hyperledger/fabric/tls
    ports:
      - 7051:7051
      - 7052:7052
      - 7053:7053

  peer0.org2.example.com:
    container_name: peer0.org2.example.com
    extends:
      file: peer-base.yaml
      service: peer-base
    environment:
      - CORE_PEER_ID=peer0.org2.example.com
      - CORE_PEER_ADDRESS=peer0.org2.example.com:7051
      - CORE_PEER_CHAINCODELISTENADDRESS=peer0.org2.example.com:7052
      - CORE_PEER_GOSSIP_EXTERNALENDPOINT=peer0.org2.example.com:7051
      - CORE_PEER_LOCALMSPID=Org2MSP
    volumes:
        - /var/run/:/host/var/run/
        - ../crypto-config/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/msp:/etc/hyperledger/fabric/msp
        - ../crypto-config/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls:/etc/hyperledger/fabric/tls
    ports:
      - 8051:7051
      - 8052:7052
      - 8053:7053
```

## 5.3 peer-base文件

```yaml
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#
version: '2'
services:
  peer-base:
    image: hyperledger/fabric-peer:latest
    environment:
      - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock
      # the following setting starts chaincode containers on the same
      # bridge network as the peers
      # https://docs.docker.com/compose/networking/
      # 与 /src/项目名称/ 保持一致，如/src/hyperledger/,则为hyperledger_default
      - CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=hyperledger_default
      #- CORE_LOGGING_LEVEL=ERROR
      - CORE_LOGGING_LEVEL=INFO
      - CORE_PEER_TLS_ENABLED=true
      # 是否自动选举leader节点，建议设置为true了，当一个leader节点挂掉后自动选举一个新的leader节点
      - CORE_PEER_GOSSIP_USELEADERELECTION=true
      # 当前节点是否是leader节点，若设置自动选举，则设置为false
      - CORE_PEER_GOSSIP_ORGLEADER=false
      - CORE_PEER_PROFILE_ENABLED=true
      - CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/fabric/tls/server.crt
      - CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/fabric/tls/server.key
      - CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/fabric/tls/ca.crt
    working_dir: /opt/gopath/src/github.com/hyperledger/fabric/peer
    command: peer node start
```

将 docker-compose-cli.yaml 文件放在 configtx.yaml 文件的同目录下，并在该目录下新建 base 目录，将 docker-compose-base.yaml 和 peer-base.yaml 放入其中。 

## 5.4 docker-compose启动

由于之前验证fabric项目时生成过同名容器，因此需要先将之前生成的容器删除。

```shell
	# 查看当前所有容器
	docker ps -a
	# 删除指定的容器，3b4e5beacu4是容器ID 
	# rmi 删除镜像
	docker rm 3b4e5beacu4
```

启动docker-compose。

```shell
	# 启动
	# -f 指定配置文件
	# -d 以守护进程运行
	docker-compose -f docker-compose-cli.yaml up -d
	# 查看启动状态
	docker-compose -f docker-compose-cli.yaml ps
	# 停止
	# -v 详细信息
	docker-compose down -v
```

# 6 创建通道并添加节点

## 6.1 客户端对peer节点的操作流程

1. 客户端创建通道。
2. 客户端将所有peer节点加入通道。
3. 为每个peer节点安装链码。
4. 初始化链码(init函数)，只需初始化任意一个peer节点的链码，之后会自动同步到其它节点。
5. 数据的读写和查询。

## 6.2 创建通道

进入客户端容器。

```shell
	# 查看容器名
	docker ps -a
	# 打开终端
	# -it 其后参数可以是容器ID或者容器名称
	docker exec -it cli /bin/bash 
	# 退出
	exit
```

创建通道。

```shell
	# 查看帮助
	peer channel create --help
	# -o orderer节点地址，域名/ip：端口号
	# -c 要创建的通道id，必须小写，需要和之前生成通道时的id保持一致
	# -f configtxgen生成的通道文件，./channel-artifacts/channel.tx
	# -t 创建通道的超时时长，默认5s
	# --tls 通信是否使用tls加密，true或false，需要与之前的配置保持一致
	# --cafile tls证书文件，必须是绝对路径，pem后缀的文件
	peer channel create -o orderer.example.com:7050 -c mychannel -f ./channel-artifacts/channel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
	# 验证是否生成后缀为 .block 的文件
	ll
```

## 6.3 将当前节点加入通道

docker-compose-cli 中在最开始配置了一个节点的信息，这个节点就是当前节点。

```shell
	peer channel join -b mychannel.block
```

## 6.4 客户端切换节点

修改环境变量：直接在cli终端中执行即可。

```shell
	# 查看当前环境变量
	echo $CORE_PEER_ID
    echo $CORE_PEER_ADDRESS
    echo $CORE_PEER_LOCALMSPID
    echo $CORE_PEER_TLS_CERT_FILE
    echo $CORE_PEER_TLS_KEY_FILE
    echo $CORE_PEER_TLS_ROOTCERT_FILE
    echo $CORE_PEER_MSPCONFIGPATH
	
	# 修改环境变量
	# peer0.org2.example.com
	export CORE_PEER_ID=org2peer0 CORE_PEER_ADDRESS=peer0.org2.example.com:7051 CORE_PEER_LOCALMSPID=Org2MSP CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/server.crt CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/server.key CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
	
	# peer0.org1.example.com
	export CORE_PEER_ID=org1peer0 CORE_PEER_ADDRESS=peer0.org1.example.com:7051 CORE_PEER_LOCALMSPID=Org1MSP CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.crt CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/server.key CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
```

将新节点加入通道。

```shell
	peer channel join -b mychannel.block
```

## 6.5 问题备忘

执行 docker-compose -f docker-compose-cli.yaml up -d 命令时，提示如下错误：

```shell
fatal error: unexpected signal during runtime execution
[signal SIGSEGV: segmentation violation code=0x1 addr=0x63 pc=0x7f1ed5f24259]
```

解决方法：在 docker-compose-cli.yaml 文件中，cli 的 environment 中添加如下配置：

```yaml
	environment:
	  - GODEBUG=netdns=go
```

# 7 链码

## 7.1 peer节点安装链码

```shell
	peer chaincode install --help
	# -c json格式参数
	# -l 语言，默认golang
	# -n chaincode名字，自定义
	# -p chaincode源码路径，默认是从gopath的src下，因此直接些src下的项目路径即可。由于只是路径，因此不能把多个go文件放在同一路径下，chaincode的挂载需要分项目文件夹
	# -v 链码版本号，可以用来区分更新的版本，不用删除原来安装的链码
	peer chaincode install -n example02 -p github.com/hyperledger/fabric/examples/chaincode/go -v 1.0.0
```

安装完成后会在 /var/hyperledger/production/chaincodes/ 路径下生成文件。

## 7.2 链码实例化

```shell
	peer chaincode instantiate --help
	# -o orderer节点的地址，域名：7050
	# -C channelID
	# -c json格式参数 
	#    '{"Args":["init","a","2","b","3"]}'
	# -l 语言
	# -n chaincode名字，已安装的链码名字
	# -P 背书策略，交易规则，指定都有哪些节点参与
	#    "OR ('Org1.member','Org2.member')"
	# -v 链码版本
	# --tls 是否使用tls加密，true或false
	# --cafile orderer节点下的pem文件路径 
	peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -c '{"Args":["init","a","100","b","200"]}' -n example02 -P "OR ('Org1.member','Org2.member')" -v 1.0.0 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

实例化后会产生新的容器 dev-peer0.org1.example.com-example02-1.0.0 ，可以通过查看进行验证。

注：这里 -P 后为背书策略，即完成一笔交易的规则，比如规定好组织1和组织2都参与，则使用AND；若其中任何一个参与即可，则使用OR。

1.2版本中在使用invoke调用的时候，通过 --peerAddress 参数指定参与的组织成员，若背书策略中为AND，则至少使用两个 --peerAddress 去分别指定组织的任意成员。

## 7.3 背书策略

1. 完成一笔交易，这笔交易前执行的方案就叫背书策略。
2. 要完成一笔交易，必须要有背书策略，使用 -P 指定。
3. 所谓背书策略，即完成一笔交易的规则，比如规定好组织1和组织2都参与，则使用AND；若其中任何一个参与即可，则使用OR。

1.0和1.1版本中关于通过 invoke 进行转账的使用：

```shell
	peer chaincode invoke -o orderer.example.com:7050 -n example02 --tls --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C xypchannel -c '{"Args":["invoke","a","b","20"]}'
```

1.2版本中，需要指定 --peerAddresses 和 --tlsRootCertFiles 两个参数。

## 7.4 命令案例

这里以查询操作为例。

```shell
	peer chaincode query --help
	# -C channelID
	# -c json参数
	# -n 链码名称
	peer chaincode query -C mychannel -n example02 -c '{"Args":["query","a"]}'
	# Query Result: 100	
```

## 7.5 智能合约约束

- package 必须是 main 。

- 必须要有 main 函数，main 函数中执行了 shim.Start 。

  - ```go
    	// SimpleChaincode 是一个结构体
    	err := shim.Start(new(SimpleChaincode))
    ```

- 必须要有一个结构体，除了 main 函数，其它的函数必须都挂载到这个结构体下，使用指针，包括 Init 函数。

- 必须要有 Init 函数，而非小写的 init 。

- 必须要有 Invoke 函数。

## 7.6 shim包

Chaincode interface：

```go
	Init(stub ChaincodeStubInterface) pb.Response
	Invoke(stub ChaincodeStubInterface) pb.Response
```

ChaincodeStubInterface interface：

```go
	GetArgs() [][]byte // 获取所有的参数，包括函数名和参数
	GetStringArgs() []string // 获取所有的参数，包括函数名和参数，转成string类型
	GetFunctionAndParameters() (string, []string) // 获取所有的参数，将函数名与参数分成两部分
	GetArgsSlice() ([]byte, error) // 获取所有的参数，包括函数名和参数，封装成切片类型
	GetTxID() string // 获取交易ID
	InvokeChaincode(chaincodeName string, args [][]byte, channel string) pb.Response // 根据链码名和Channel调用链码
	GetState(key string) ([]byte, error) // 根据指定用户查看状态
	PutState(key string, value []byte) error // 写入用户状态到账本
	DelState(key string) error // 删除指定用户
	GetStateByRange(startKey, endKey string) (StateQueryIteratorInterface, error) // 通过循环方式获取用户状态
	GetStateByPartialCompositeKey(objectType string, keys []string) (StateQueryIteratorInterface, error) // 通过特殊组合键获取用户状态
	CreateCompositeKey(objectType string, attributes []string) (string, error) // 创建组合键
	SplitCompositeKey(compositeKey string) (string, []string, error) // 切割组合键
GetQueryResult(query string) (StateQueryIteratorInterface, error) // CounchDB查询结果，couchdb文档 https://github.com/cloudant/mango
	GetHistoryForKey(key string) (HistoryQueryIteratorInterface, error) // 获取key的历史数据
	GetCreator() ([]byte, error) // 获取当前用户
	GetTransient() (map[string][]byte, error) // 返回可以被链码使用但没有保存在账本中的临时映射表，例如用于加解密的密码学信息
	GetBinding() ([]byte, error) // 获取交易的nonce、creator和epoch拼接结果的SHA256哈希
	GetSignedProposal() (*pb.SignedProposal, error) // 返回完全解码的签名交易对象
	GetTxTimestamp() (*timestamp.Timestamp, error) // 返回交易创建时的时间戳
	SetEvent(name string, payload []byte) error // 当Chaincode提交完毕，会通过Event的方式通知Client
```

## 7.7 链码升级

将代码放到链码挂载的目录中，然后替换原来的旧代码，或者直接从宿主主机将链码拷贝进 cli 容器中。

```shell
 docker cp chaincode01.go cli:/opt/gopath/src/github.com/hyperledger/fabric/examples/chaincode/go/mycc01
```

安装新版本的链码。

```shell
	docker exec -it cli bash
	# -n 后的链码名称必须和之前保持一致，只需要更改版本号
	peer chaincode install -n chaincode01.go -v 1.1 -p github.com/hyperledger/fabric/examples/chaincode/go/mycc01
```

更新链码。

```shell
	peer chaincode upgrade -o orderer.example.com:7050 -C mychannel -c '{"Args":["init","name","zhangsan","age","18"]}' -n example02 -P "OR ('Org1.member','Org2.member')" -v 1.1 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```

## 7.8 交易流程

1. Application SDK，客户端进行交易，使用 peer chaincode invoke 命令。
2. Peer 节点收到客户端的提案（背书策略）。
3. Peer 节点将交易进行预演，而非真正的交易。
4. Peer 节点将预演结果发送至客户端，如果失败则结束整个流程。
5. 客户端将交易提交给排序节点，预演的结果是成功的才会提交到排序节点。
6. 排序节点对交易进行打包，并将打包的数据发送给 Peer 节点。
7. Peer 节点将数据写入区块中，当满足配置条件中的。

## 7.9 链码打包

将链码相关数据（如链码名称、版本、实例化策略等信息）进行封装。

```shell
	# -n 链码名称
	# -p 链码路径
	# -v 链码版本
	# 最后加打包后生成的文件名称，后缀可以为.out或者.pak
	peer chaincode package -n chaincode01 -p github.com/hyperledger/fabric/examples/chaincode/go/mycc01 -v 1.0 chaincode01.1.0.out
```

打包后的链码，可以直接使用 install 进行安装。

```shell
	peer chaincode install chaincode01.1.0.out
```



## 7.10 问题备忘

实例化时出现错误一：

```shell
	# 同一个包中不能出现两个main函数
	xxx main redeclared in this block
    previous declaration at xxx
```

解决方法：

重新安装链码，指定 -p 的路径一定和原来的不在同一个目录， -n 参数不能和原来的一样。

```shell
	# 停止容器
	docker stop containerID
	# 删除容器
	docker rm containerID
	# 删除镜像
	docker rmi imageID
```

实例化时出现错误二：

```shell
	Failed to execute transaction (Timeout expired while executing transaction)
```

解决方法：

查看容器日志，可能存在代码错误。

```shell
	docker logs containerID
```

# 8 fabric账户

## 8.1 什么是账户

- 证书与密钥文件（而非传统的用户名和密码）。

- 证书与密钥文件在 msp（Membership Service Providers）路径中，组织的 msp 和组织下对应节点的 msp 目录不一样。

- 相关命令。

  - ```shell
    	sudo apt-get install tree
    	# 查看树结构
    	cd crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/msp
    	tree
    ```

## 8.2 账户结构

组织 msp：

```shell
msp/
├── admincerts                     # 管理员证书
│   └── Admin@org1.example.com-cert.pem
├── cacerts                        # 根 CA 服务器的证书
│   └── ca.org1.example.com-cert.pem
└── tlscacerts                     # TLS 根 CA 的证书，tls 通信时会使用
    └── tlsca.org1.example.com-cert.pem
```

节点 msp：

```shell
msp/
├── admincerts                     # 管理员证书
│   └── Admin@org1.example.com-cert.pem
├── cacerts                        # 根 CA 服务器的证书
│   └── ca.org1.example.com-cert.pem
├── keystore                       # 节点或者账户的私钥
│   └── 9458c81f63e150e236c13626e2bc2cc69a9ad11dcc52ff73163277977cab4705_sk
├── signcerts                      # 符合 x.509 的节点或者用户证书文件
│   └── peer0.org1.example.com-cert.pem
└── tlscacerts                     # TLS 根 CA 的证书，tls 通信时会使用
    └── tlsca.org1.example.com-cert.pem
```

## 8.3 账户作用

fabric 中每条交易都会加上发起者的标签（签名证书），同时用发起人的私钥进行加密。

如果交易需要其它组织的节点提高背书功能，那么背书节点也会在交易中加入自己的签名。这样每一笔交易的操作过程会非常清晰并且不可篡改。

## 8.4 账户生成

cryptogen 模块是创建 fabric 账户的方式之一。

```shell
	cryptogen generate --config crypto_config.yaml
```

## 8.5 fabric账户的使用场景

1. 生成创世块和通道文件时需要在配置文件中使用。
2. docker-compose 配置文件中需要使用。
3. 创建通道时需要使用。
4. 启动 Orderer 账户。
5. 启动 Peer 账户。

## 8.6 fabric-ca

动态添加账户。

模板路径：examples/e2e_cli/docker-compose-e2e.yaml

在 services 配置项下添加 ca 的配置即可。

```yaml
	ca0:
	  # 镜像
      image: hyperledger/fabric-ca
      environment:
      	# 下面的映射路径
        - FABRIC_CA_HOME=/etc/hyperledger/fabric-ca-server
        # ca服务器名，可以更改
        - FABRIC_CA_SERVER_CA_NAME=ca-org1
        # 是否启用 tls
        - FABRIC_CA_SERVER_TLS_ENABLED=true
        # 证书文件
        - FABRIC_CA_SERVER_TLS_CERTFILE=/etc/hyperledger/fabric-ca-server-config/ca.org1.example.com-cert.pem
        # 密钥文件
        - FABRIC_CA_SERVER_TLS_KEYFILE=/etc/hyperledger/fabric-ca-server-config/2628c774d6def25e2bf6147c30f25fe76469d63d257965ac867544acd090148c_sk
      ports:
        - "7054:7054"
      # -b 指定登录的用户名密码，可以更改
      command: sh -c 'fabric-ca-server start --ca.certfile /etc/hyperledger/fabric-ca-server-config/ca.org1.example.com-cert.pem --ca.keyfile /etc/hyperledger/fabric-ca-server-config/2628c774d6def25e2bf6147c30f25fe76469d63d257965ac867544acd090148c_sk -b admin:adminpw -d'
      # 挂载路径
      volumes:
        - ./crypto-config/peerOrganizations/org1.example.com/ca/:/etc/hyperledger/fabric-ca-server-config
      # 容器名，可以更改
      container_name: ca_peerOrg1
```

注：每个 Peer 组织都需要一个 ca 配置项。

# 9 shell脚本

## 9.1 shell脚本介绍

文档链接：[Shell 教程 | 菜鸟教程 (runoob.com)](https://www.runoob.com/linux/linux-shell.html)

## 9.2 fabric网络搭建执行顺序

1. **手动**拷贝可执行文件到 gopath/bin/ 下。
2. **手动**拷贝 crypto-config.yaml 到 shell 脚本同级目录下。
3. **手动**拷贝 configtx.yaml 到 shell 脚本同级目录下。
4. **手动**拷贝 dockers-compose-cli.yaml 到 shell 脚本同级目录下。
5. **手动**拷贝 docker-compose-base.yaml 和 peer-base.yaml 到 base/ 路径下。
6. **手动**拷贝链码到 chaincode/go/ 目录下。
7. 下载镜像。
8. 生成证书，生成 crypto-config/ 路径。
9. 生成创世块和通道文件，要求生成文件到 channel-artifacts/ 目录。
10. 启动 dockers-compose，启动 peer、orderer、cli 等容器。
11. 进入 cli 容器。
12. 创建通道。
13. 将当前 peer 节点加入通道，切换 peer 节点环境，继续将 peer 节点加入通道。
14. 安装链码。
15. 实例化链码。
16. 调用链码，通过客户端或者 sdk(go-sdk)。

## 9.3 shell脚本

