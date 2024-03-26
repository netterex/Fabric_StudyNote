#!/bin/bash

# 停止docker-compose
docker-compose -f ./docker-compose-cli.yaml down -v
# 获取链码容器id
contain_ids=$(docker ps -a |awk '($2 ~ /dev-*/){print $1}')
echo $contain_ids
# 停止链码容器
docker stop $contain_ids
# 删除链码容器
docker rm $contain_ids
# 获取镜像id
image_ids=$(docker images |awk '($1 ~ /dev-*/){print $3}')
echo $image_ids
# 删除镜像
docker rmi $image_ids
# 删除通道文件
rm -rf channel-artifacts/ policechannel.block housemngchannel.block creditchannel.block
