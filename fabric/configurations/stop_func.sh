#!/bin/bash

# 删除链码容器
function RemoveChaincodeContainers() {
    container_ids=$(docker ps -a |awk '($2 ~ /dev-*/){print $1}')
    if [ -z "$container_ids" -o "$container_ids" == "" ]; then
      echo "---- 没有可删除的容器 ----"
    else
      docker rm -f $container_ids
    fi
}

# 删除链码镜像
function RemoveChaincodeImages() {
  image_ids=$(docker images |awk '($1 ~ /dev-*/){print $3}')
  if [ -z "$image_ids" -o "$image_ids" == " " ]; then
    echo "---- 没有可删除的镜像 ----"
  else
    docker rmi -f $image_ids
  fi
}

echo "---- 关闭fabric网络 ----"

# 停止docker-compose
docker-compose -f ./docker-compose-cli.yaml down -v

# 调用函数删除链码容器
RemoveChaincodeContainers

# 调用函数删除链码镜像
RemoveChaincodeImages