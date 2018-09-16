#!/bin/bash

# elasticsearch 安装目录
installPath=$1
elasticsearch_install_path=${installPath}/elasticsearch

ID=`ps -ef | grep elasticsearch | grep -v "elasticsearch_install.sh" |grep -v "elasticsearch_" | grep -v "$0" | grep -v "grep" | grep ".sh" | awk '{print $2}'`
echo $ID
echo "---------------"
for id in $ID
do
kill -9 $id
echo "killed $id"
done

if [ -d ${elasticsearch_install_path} ];then
    rm -rf ${elasticsearch_install_path}
fi
