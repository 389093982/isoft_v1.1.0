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

# 进行卸载操作
sh ./api_undeploy.sh ${service_name} ${package_name}

# 重新创建对应目录
if [ ! -d ${deploy_home}/project/goproject/${service_name} ];then
    mkdir -p ${deploy_home}/project/goproject/${service_name}
    echo "mkdir -p ${deploy_home}/project/goproject/${service_name} successful..."
fi

# 进入应用所在目录拷贝对应文件,并设置可执行权限
cd ${deploy_home}/project/goproject/${service_name}
cp ${deploy_home}/upload/api/packages/${service_name}/${package_name} .
# 重命名并授权
mv ${package_name} "${service_name}_${package_name}" && chmod +x ./"${service_name}_${package_name}"

# 启动应用
cd ${sh_home} && sh ./api_startup.sh ${service_name} ${package_name} ${runmode}

