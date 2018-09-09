#!/bin/bash

################################## 多实例场景下 service_name 和 package_name 命名不同 ########################################
# 服务名称
service_name=$1
# 不帶后缀的软件包名
package_name=$2
# 运行模式
runmode=$3
###########################################################################################################################


deploy_home=`echo $(cd ../.. &&  pwd)`

count=`ps -ef | grep "./${service_name}_${package_name}" | grep -v grep |wc -l`

if [ 0 == $count ];then
    if [ -d ${deploy_home}/project/goproject/${service_name}/"${service_name}_${package_name}" ];then
        echo "api_check__STOP"
    else
        echo "api_check__N/A"
    fi
else
    echo "api_check__RUN"
fi