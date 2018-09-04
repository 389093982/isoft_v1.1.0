#!/bin/bash
###################################################################################################################
# 目标机器 deploy_home 路径
remoteDeployHomePath=$1
# 服务名称
serviceName=$2
# 服务占用端口
servicePort=$3
# mysql root 密码
rootPwd=$4
###################################################################################################################
if [ -z ${remoteDeployHomePath} ] || [ -z ${serviceName} ] || [ -z ${servicePort} ] || [ -z ${rootPwd} ];then
    echo "invalid params"
    exit;
fi

if [ "${servicePort}" == "_" ];then
    # 使用默认端口 3306
    servicePort="3306"
fi

sh ./mysql_check.sh ${remoteDeployHomePath} ${serviceName} ${servicePort}

result=`docker ps -aq --filter name="${serviceName}\$"`

if [ "${result}" != "" ];then
    # 先停止再删除运行的容器
    docker stop --time=20 ${serviceName}

    # 强制移除此容器
    docker rm -f $(docker ps -aq --filter name="${serviceName}\$")

    # 清理此容器的网络占用
    docker network disconnect --force bridge ${serviceName}
fi

# 杀掉占用的端口
sh ../common/port_kill.sh ${servicePort}

###################################################################################################
# 数据不删除
#if [ -d /var/lib/mysql ];then
#    rm -rf /var/lib/mysql
#fi
###################################################################################################

sh ./mysql_check.sh ${remoteDeployHomePath} ${serviceName} ${servicePort}
