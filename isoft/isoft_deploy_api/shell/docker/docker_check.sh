#!/bin/bash

# 获取 docker 版本
docker_version=`docker version | grep Version | grep -v grep | head -n 1 | awk '{ print $2 }'`
echo "docker_version__${docker_version}"

# 从第二行开始读取每一行
docker images | awk 'NR>1' | while read line
do
    echo "docker_images__${line}"
done
