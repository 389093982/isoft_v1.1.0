#!/bin/bash

# 服务名称
service_name=$1
# 不帶后缀的软件包名
package_name=$2

sh ./beego_shutdown.sh ${service_name} ${package_name} && sh ./beego_startup.sh ${service_name} ${package_name}
