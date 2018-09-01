#!/bin/bash

# 服务名称
service_name=$1
# 不帶后缀的软件包名
package_name=$2

deploy_home=`echo $(cd ../.. &&  pwd)`

# 先杀进程
sh ./beego_shutdown.sh ${service_name} ${package_name}

# 再删除应用
if [ -d ${deploy_home}/project/goproject/${service_name} ];then
    rm -rf ${deploy_home}/project/goproject/${service_name}
    echo "Remove ${package_name} references file..."
fi