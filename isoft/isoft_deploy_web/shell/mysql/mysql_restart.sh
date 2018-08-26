#!/bin/bash

# 目标机器 deploy_home 路径
remoteDeployHomePath=$1
# 服务名称
serviceName=$2
# 服务占用端口
servicePort=$3
# mysql root 密码
rootPwd=$4

if [ -z ${remoteDeployHomePath} ] || [ -z ${serviceName} ] || [ -z ${servicePort} ] || [ -z ${rootPwd} ];then
    echo "invalid params"
    exit;
fi

if [ "${servicePort}" == "_" ];then
    # 使用默认端口 3306
    servicePort="3306"
fi

sh ./mysql_check.sh ${remoteDeployHomePath} ${serviceName} ${servicePort}

# 杀掉占用的端口
sh ../common/port_kill.sh ${servicePort}

# 先停止再删除运行的容器
docker stop --time=20 ${serviceName} && docker rm -f $(docker ps -aq --filter name="${serviceName}\$")

sh ./mysql_check.sh ${remoteDeployHomePath} ${serviceName} ${servicePort}

###################################################################################################
# mysql 安装目录
mysql_install_home="${remoteDeployHomePath}/soft/install/${serviceName}"

if [ -d ${mysql_install_home} ];then
    rm -rf ${mysql_install_home}
fi
mkdir -p ${mysql_install_home}/logs && mkdir -p ${mysql_install_home}/conf && mkdir -p ${mysql_install_home}/data

# 拷贝配置文件 my.conf 到对应位置
if [ ! -f "${mysql_install_home}/conf/my.cnf" ];then
    cp -r ./my.cnf ${mysql_install_home}/conf
fi
###################################################################################################

# 运行 docker mysql 容器
result=`docker run -p ${servicePort}:3306 --name ${serviceName} \
    -v ${mysql_install_home}/conf:/etc/mysql/conf.d -v ${mysql_install_home}/logs:/logs \
    -v ${mysql_install_home}/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=${rootPwd} -d mysql`

echo ${result}

sh ./mysql_check.sh ${remoteDeployHomePath} ${serviceName} ${servicePort}

echo "#########################################################################"
echo "Please use the following command to enter MySQL"
echo "docker exec -it \$(docker ps -aq --filter name="${serviceName}\$") /bin/bash"
echo "#########################################################################"
