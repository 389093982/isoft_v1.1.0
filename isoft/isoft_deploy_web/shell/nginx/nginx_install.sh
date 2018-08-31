#!/bin/bash

# 目标机器 deploy_home 路径
remoteDeployHomePath=$1
# 服务名称
serviceName=$2
# 服务占用端口
servicePort=$3

if [ -z ${remoteDeployHomePath} ] || [ -z ${serviceName} ] || [ -z ${servicePort} ];then
    echo "invalid params"
    exit;
fi

# 是否需要下载 docker nginx 镜像
nginxCheck=`docker images | grep nginx | grep -v grep`
if [ ! -n "${nginxCheck}" ];then
    docker pull nginx
fi

sh ./nginx_restart.sh ${remoteDeployHomePath} ${serviceName} ${servicePort}