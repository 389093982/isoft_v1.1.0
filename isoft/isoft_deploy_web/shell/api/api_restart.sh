#!/bin/bash

################################## 多实例场景下 service_name 和 package_name 命名不同 ########################################
# 服务名称
service_name=$1
# 不帶后缀的软件包名
package_name=$2
# 运行模式
runmode=$3
###########################################################################################################################


sh ./api_shutdown.sh ${service_name} ${package_name}
sh ./api_startup.sh ${service_name} ${package_name}
