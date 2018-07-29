#!/bin/bash

# 应用名称
app_name=$1

sh_home=`pwd`
deploy_home=`echo $(cd ../.. &&  pwd)`

cd ${deploy_home}/project/goproject/${app_name} && ./${app_name} &

cd ${sh_home}

sleep 5

sh ./beego_status.sh ${app_name}