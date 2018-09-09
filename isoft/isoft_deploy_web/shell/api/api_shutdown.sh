#!/bin/bash

################################## 多实例场景下 service_name 和 package_name 命名不同 ########################################
# 服务名称
service_name=$1
# 不帶后缀的软件包名
package_name=$2
# 运行模式
runmode=$3
###########################################################################################################################


# 先杀进程
PROCESS=`ps -ef | grep "./${service_name}_${package_name}" | grep -v grep | grep -v PPID | awk '{ print $2}'`
for i in $PROCESS
do
    kill -9 $i
    echo "Kill the ${package_name} process [ $i ]"
done

sh ./api_check.sh ${service_name} ${package_name}