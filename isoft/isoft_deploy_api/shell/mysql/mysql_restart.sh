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

docker restart ${serviceName}

sh ./mysql_check.sh ${remoteDeployHomePath} ${serviceName} ${servicePort}

echo "#########################################################################"
echo "Please use the following command to enter MySQL"
echo "docker exec -it \$(docker ps -aq --filter name="${serviceName}\$") /bin/bash"
echo "#########################################################################"
