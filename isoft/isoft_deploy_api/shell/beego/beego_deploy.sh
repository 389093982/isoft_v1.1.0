#!/bin/bash

# 应用名称
app_name=$1
# 运行模式
runmode=$2

sh_home=`pwd`
deploy_home=`echo $(cd ../.. &&  pwd)`

# 进行卸载操作
sh ./beego_undeploy.sh ${app_name}

if [ ! -d ${deploy_home}/project/goproject/${app_name} ];then
    mkdir -p ${deploy_home}/project/goproject/${app_name}
    echo "mkdir -p ${deploy_home}/project/goproject/${app_name}"
fi

# 进入应用所在目录并解压 tar.gz 包,并设置可执行权限
cd ${deploy_home}/project/goproject/${app_name} && tar -xzf ../../../beego/packages/${app_name}.tar.gz && chmod +x ./${app_name}

# 修改配置文件
old_runmode=`cat ./conf/app.conf | grep runmode | grep -v grep`
new_runmode="runmode = ${runmode}"
sed -i s/"${old_runmode}"/"${new_runmode}"/g ./conf/app.conf

# 启动应用
cd ${sh_home} && sh ./beego_startup.sh ${app_name}

