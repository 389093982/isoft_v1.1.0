#!/bin/bash

# 容器名称
serviceName=$1

# 进入容器
if [ -z ${serviceName} ];then
    echo "invalid params"
    exit;
fi

docker exec -it $(docker ps -aq --filter name="${serviceName}\$") /bin/bash