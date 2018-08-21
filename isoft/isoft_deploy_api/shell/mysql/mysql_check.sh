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

if [ "${servicePort}" == "_" ];then
    # 使用默认端口 3306
    servicePort="3306"
fi

# mysql 安装目录
mysql_install_home="${remoteDeployHomePath}/soft/install/${serviceName}"

if [ -d ${mysql_install_home} ];then
    # 根据端口号查询对应的pid
    pid=$(netstat -nlp | grep :${servicePort} | awk '{print $7}' | awk -F"/" '{ print $1 }');
    if [  -n  "$pid"  ];  then
        echo "mysql_check__RUN"
    else
        echo "mysql_check__STOP"
    fi
else
    echo "mysql_check__N/A"
fi

# docker logs $(docker ps -aq --filter name=${serviceName})

###################################################################################################################
# MySQL测试环境遇到 mmap(xxx bytes) failed; errno 12解决方法
# 明显的swap问题,适当增加swap
# sudo dd if=/dev/zero of=/swapfile bs=1M count=1024 #增加1G的SWAP进去
# sudo mkswap /swapfile
# sudo swapon /swapfile
# free
###################################################################################################################
# ERROR 2059 (HY000): Authentication plugin 'caching_sha2_password' cannot be loaded; 的解决办法
# 关于这个问题,看起来很难,实则很简单,例如我需要在IP地址为192.168.78.138的主机上,
# 远程登录到安装好的MySQL数据库服务,则需要在MySQL服务上添加一个IP为192.168.78.138的用户即可
# 注意:如果你的也是最新版本,则需要在my.ini的[mysqld]下添加一行:
# default_authentication_plugin = mysql_native_password
# 在重新初始化MySQL服务即可,用户添加完成后,现在就可以远程进行登录了
###################################################################################################################