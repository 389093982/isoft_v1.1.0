#!/bin/bash

################################## 多实例场景下 service_name 和 package_name 命名不同 ########################################
# 服务名称
service_name=$1
# 不帶后缀的软件包名
package_name=$2
# 运行模式
runmode=$3
###########################################################################################################################

sh_home=`pwd`
deploy_home=`echo $(cd ../.. &&  pwd)`

cd ${deploy_home}/project/goproject/${service_name} && ./"${service_name}_${package_name}" ${runmode} &

cd ${sh_home}

sleep 5

sh ./api_check.sh ${service_name} ${package_name}