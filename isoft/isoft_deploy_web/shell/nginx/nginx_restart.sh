#!/bin/bash

# 目标机器 deploy_home 路径
remoteDeployHomePath=$1
# 服务名称
serviceName=$2
# 服务占用端口
servicePort=$3

if [ -z ${remoteDeployHomePath} ] || [ -z ${serviceName} ] || [ -z ${servicePort} ];then
    echo "invalid params"
    exit;
fi

if [ "${servicePort}" == "_" ];then
    # 使用默认端口 80
    servicePort="80"
fi

sh ./nginx_check.sh ${remoteDeployHomePath} ${serviceName} ${servicePort}

# 杀掉占用的端口
sh ../common/port_kill.sh ${servicePort}

# 先停止再删除运行的容器
docker stop --time=20 ${serviceName} && docker rm -f $(docker ps -aq --filter name="${serviceName}\$")

sh ./nginx_check.sh ${remoteDeployHomePath} ${serviceName} ${servicePort}

###################################################################################################
# nginx 安装目录
nginx_install_home="${remoteDeployHomePath}/soft/install/${serviceName}"

if [ ! -d ${nginx_install_home} ];then
    mkdir -p ${nginx_install_home}
else
    rm -rf ${nginx_install_home} && mkdir -p ${nginx_install_home}
fi
# 拷贝主配置文件 nginx.conf 和子配置文件 default.conf 到对应位置
if [ ! -d "${nginx_install_home}/conf.d" ];then
    cp -r ./conf.d ${nginx_install_home}
fi
if [ ! -f "${nginx_install_home}/nginx.conf" ];then
    cp -r ./nginx.conf ${nginx_install_home}
fi
if [ ! -d "${nginx_install_home}/html" ];then
    cp -r ./html ${nginx_install_home}
fi
###################################################################################################

# 运行 docker nginx 容器
result=`docker run \
          --name ${serviceName} \
          -d -p ${servicePort}:80 \
          -v ${nginx_install_home}/html:/usr/share/nginx/html \
          -v ${nginx_install_home}/nginx.conf:/etc/nginx/nginx.conf:ro \
          -v ${nginx_install_home}/conf.d:/etc/nginx/conf.d \
          nginx`

echo ${result}

sh ./nginx_check.sh ${remoteDeployHomePath} ${serviceName} ${servicePort}

echo "#########################################################################"
echo "Please use the following command to enter Nginx"
echo "docker exec -it \$(docker ps -aq --filter name="${serviceName}\$") /bin/bash"
echo "#########################################################################"
