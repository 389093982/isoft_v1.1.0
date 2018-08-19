#!/bin/bash

# 应用名称
app_name=$1

deploy_home=`echo $(cd ../.. &&  pwd)`

count=`ps -ef | grep "./${app_name}" | grep -v grep |wc -l`

if [ 0 == $count ];then
    if [ -d ${deploy_home}/project/goproject/${app_name} ];then
        echo "beego_check__STOP"
    else
        echo "beego_check__N/A"
    fi
else
    echo "beego_check__RUN"
fi