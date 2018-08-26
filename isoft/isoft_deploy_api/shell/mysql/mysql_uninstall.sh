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

# 先停止再删除运行的容器
docker stop --time=20 ${serviceName}

# 强制移除此容器
docker rm -f $(docker ps -aq --filter name="${serviceName}\$")

# 删除与容器相关联的卷
docker rm -v ${serviceName}

# 清理此容器的网络占用
docker network disconnect --force bridge ${serviceName}

# 杀掉占用的端口
sh ../common/port_kill.sh ${servicePort}

###################################################################################################
# mysql 安装目录
mysql_install_home="${remoteDeployHomePath}/soft/install/${serviceName}"

if [ -d ${mysql_install_home} ];then
    rm -rf ${mysql_install_home}
fi
###################################################################################################

sh ./mysql_check.sh ${remoteDeployHomePath} ${serviceName} ${servicePort}