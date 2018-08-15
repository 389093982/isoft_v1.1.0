#!/bin/bash

# 目标机器 deploy_home 路径
remoteDeployHomePath=$1
# 服务名称
serviceName=$2
# 服务占用端口
servicePort=$3
# mysql root 密码
rootPwd=$4

if [ -z ${remoteDeployHomePath} ] || [ -z ${serviceName} ] || [ -z ${servicePort} ] || [ -z ${rootPwd} ];then
    echo "invalid params"
    exit;
fi

if [ "${servicePort}" == "_" ];then
    # 使用默认端口 3306
    servicePort="3306"
fi

# 是否需要下载 docker mysql 镜像
mysqlCheck=`docker images | grep mysql | grep -v grep`
if [ ! -n "${mysqlCheck}" ];then
    docker pull mysql
fi

sh ./mysql_restart.sh ${remoteDeployHomePath} ${serviceName} ${servicePort} ${rootPwd}