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

sh ./mysql_uninstall.sh ${remoteDeployHomePath} ${serviceName} ${servicePort} ${rootPwd}

mkdir -p /var/lib/mysql
cp my.cnf /var/lib/mysql
mkdir -p /var/shell/mysql
cp mysql_modify.sh /var/shell/mysql

docker run -p ${servicePort}:3306 --name mysql -v /var/lib/mysql/conf:/etc/mysql/conf.d -v \
/var/lib/mysql/logs:/logs -v /var/lib/mysql/data:/var/lib/mysql -v /var/shell/mysql:/var/shell/mysql \
-e MYSQL_ROOT_PASSWORD=Mysql@123 -d mysql

# 将 root 密码写入临时文件
echo ${rootPwd} > /var/shell/mysql/rootPwd.txt

# 执行优化操作
sh ./mysql_adjust.sh

sleep 2

sh ./mysql_check.sh ${remoteDeployHomePath} ${serviceName} ${servicePort}
