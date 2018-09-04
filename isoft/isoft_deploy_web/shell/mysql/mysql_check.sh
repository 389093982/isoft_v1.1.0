#!/bin/bash
###################################################################################################
# 目标机器 deploy_home 路径
remoteDeployHomePath=$1
# 服务名称
serviceName=$2
# 服务占用端口
servicePort=$3
###################################################################################################
if [ -z ${remoteDeployHomePath} ] || [ -z ${serviceName} ] || [ -z ${servicePort} ];then
    echo "invalid params"
    exit;
fi

if [ "${servicePort}" == "_" ];then
    # 使用默认端口 3306
    servicePort="3306"
fi

check_result=`docker ps -aq --filter name="${serviceName}\$"`
if [ "${result}" != "" ];then
    echo "mysql_check__N/A"
else
    # 根据端口号查询对应的pid
    pid=$(netstat -nlp | grep :${servicePort} | awk '{print $7}' | awk -F"/" '{ print $1 }');
    if [  -n  "$pid"  ];  then
        echo "servicePort ${servicePort} is running..."

        result=`docker ps -aq --filter name="${serviceName}\$"`
        echo "docker ps -aq --filter name=${serviceName}\$ result is ${result}"

        if [ "${result}" != "" ];then
            echo "mysql_check__RUN"
        else
            echo "mysql_check__STOP"
        fi
    else
        echo "mysql_check__STOP"
    fi
fi
