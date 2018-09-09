#!/bin/bash

################################################################################################################
# 环境变量名
env_name=$1
# 环境变量值
env_value=$2
################################################################################################################

count=`cat /etc/profile |grep export | grep ${env_name}=| wc -l`

if [ "${count}" == "1" ];then
	echo "${env_name} is exist"
	# 查找行号并进行删除
    search_row_num=`cat -n /etc/profile |grep export | grep ${env_name}= | awk '{print $1}'`
    # sed 中变量转义使用 '$(echo $search_row_num)' 格式
    # 根据行号进行删除
    sed -i ''$(echo $search_row_num)'d' /etc/profile
    # 在文件末尾追加一行
    echo "export ${env_name}=${env_value}" >> /etc/profile && source /etc/profile
    echo "modify success..."
else
    if [ "${count}" == "0" ];then
        # 在文件末尾追加一行
        echo "export ${env_name}=${env_value}" >> /etc/profile && source /etc/profile
        echo "write success..."
    else
        echo "multi row was found..."
    fi
fi


