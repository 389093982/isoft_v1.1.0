#!/bin/bash

# 应用名称
app_name=$1

deploy_home=`echo $(cd ../.. &&  pwd)`

# 先杀进程
sh ./beego_shutdown.sh ${app_name}

# 再删除应用
if [ -d ${deploy_home}/project/goproject/${app_name} ];then
    rm -rf ${deploy_home}/project/goproject/${app_name}
    echo "Remove ${app_name} references file..."
fi