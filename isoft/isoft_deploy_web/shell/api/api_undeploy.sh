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

# 先杀进程
sh ./api_shutdown.sh ${service_name} ${package_name} ${runmode}

# 再删除应用
if [ -d ${deploy_home}/project/goproject/${service_name} ];then
    rm -rf ${deploy_home}/project/goproject/${service_name}
    echo "Remove ${package_name} references file..."
fi