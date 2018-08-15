#!/bin/bash

# 应用名称
app_name=$1

# 先杀进程
PROCESS=`ps -ef | grep "./${app_name}" | grep -v grep | grep -v PPID | awk '{ print $2}'`
for i in $PROCESS
do
    kill -9 $i
    echo "Kill the ${app_name} process [ $i ]"
done

sh ./beego_status.sh ${app_name}