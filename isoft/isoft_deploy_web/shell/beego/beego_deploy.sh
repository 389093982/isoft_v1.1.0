#!/bin/bash

################################## 多实例场景下 service_name 和 package_name 命名不同 ########################################
# 服务名称
service_name=$1
# 不帶后缀的软件包名
package_name=$2
# 运行模式
runmode=$3
# 端口号
service_port=$4
###########################################################################################################################

sh_home=`pwd`
deploy_home=`echo $(cd ../.. &&  pwd)`

# 进行卸载操作
sh ./beego_undeploy.sh ${service_name} ${package_name}

# 重新创建对应目录
if [ ! -d ${deploy_home}/project/goproject/${service_name} ];then
    mkdir -p ${deploy_home}/project/goproject/${service_name}
    echo "mkdir -p ${deploy_home}/project/goproject/${service_name} successful..."
fi

# 进入应用所在目录并解压 tar.gz 包,并设置可执行权限
cd ${deploy_home}/project/goproject/${service_name}
# 解压操作
tar -xzf ${deploy_home}/upload/beego/packages/${service_name}/${package_name}.tar.gz
# 重命名并授权
mv ${package_name} "${service_name}_${package_name}" && chmod +x ./"${service_name}_${package_name}"

# 修改配置文件
old_runmode=`cat ./conf/app.conf | grep runmode | grep -v grep`
new_runmode="runmode = ${runmode}"
sed -i s/"${old_runmode}"/"${new_runmode}"/g ./conf/app.conf
# 修改端口号
old_httpport=`cat ./conf/app.conf | grep httpport | grep -v grep`
new_httpport="httpport = ${service_port}"
sed -i s/"${old_httpport}"/"${new_httpport}"/g ./conf/app.conf

# 启动应用
cd ${sh_home} && sh ./beego_startup.sh ${service_name} ${package_name}
