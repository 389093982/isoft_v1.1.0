#!/bin/bash

# --name 为容器指定一个名称
# -e 设置环境变量
# -d 后台运行容器,并返回容器ID

# 先删除容器
result=`docker ps -aq --filter name=docker_mysql`
if [ ! -z "${result}" ];then
	docker rm ${result}
fi

if [ ! -d /root/soft/install/mysql/mysql ];then
	mkdir -p /root/soft/install/mysql/mysql # 用于挂载mysql数据文件
fi
if [ ! -d /root/soft/install/mysql/conf.d ];then
        mkdir -p /root/soft/install/mysql/conf.d # 用于挂载mysql配置文件
fi

# -u 是制定容器内部用户,没必要,默认已经指定mysql

# 再创建容器并运行
docker run --name docker_mysql -p 3306:3306 -v /root/soft/install/mysql/mysql:/var/lib/mysql \
-v /root/soft/install/mysql/conf.d:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=123456 -d docker.io/mysql \
--character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
