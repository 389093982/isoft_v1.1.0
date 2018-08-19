#!/bin/bash

# 服务名称
service_name=$1
# 不帶后缀的软件包名
package_name=$2

deploy_home=`echo $(cd ../.. &&  pwd)`

count=`ps -ef | grep "./${service_name}/${package_name}" | grep -v grep |wc -l`

if [ 0 == $count ];then
    if [ -d ${deploy_home}/project/goproject/${service_name}/${package_name} ];then
        echo "beego_check__STOP"
    else
        echo "beego_check__N/A"
    fi
else
    echo "beego_check__RUN"
fi