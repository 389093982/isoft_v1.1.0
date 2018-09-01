#!/bin/bash

# 服务名称
service_name=$1
# 不帶后缀的软件包名
package_name=$2

sh_home=`pwd`
deploy_home=`echo $(cd ../.. &&  pwd)`

cd ${deploy_home}/project/goproject && ./${service_name}/${package_name}/${package_name} &

cd ${sh_home}

sleep 5

sh ./beego_status.sh ${service_name} ${package_name}