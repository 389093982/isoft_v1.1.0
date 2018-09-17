#!/bin/bash

# 目标机器 deploy_home 路径
remoteDeployHomePath=$1
# elasticsearch 安装目录
installPath=$2

if [ -z ${remoteDeployHomePath} ] || [ -z ${installPath} ];then
    echo "invalid param error..."
    exit;
fi

sh ./elasticsearch_uninstall.sh ${installPath}
# elasticsearch 安装包路径
packageName=`ls ${remoteDeployHomePath}/install/elasticsearch | grep elasticsearch | grep tar.gz`

if [ -z ${packageName} ];then
    echo "packageName was not found..."
    exit;
fi

elasticsearch_targz=${remoteDeployHomePath}/install/elasticsearch/${packageName}

# 解压
sh_home=`pwd`
cd ${installPath} && tar -xzf ${elasticsearch_targz}
echo "tar -xzf ${elasticsearch_targz} success..."
cd ${sh_home}
installName=`ls ${installPath} | grep elasticsearch | grep -v logs | grep -v data | grep -v tar.gz | head -1`
mv ${installPath}/${installName} ${installPath}/elasticsearch

# 修改 elasticsearch 数据和日志存储目录
#设置索引数据的存储路径
#设置日志的存储路径
elasticsearch_yml=${installPath}/elasticsearch/config/elasticsearch.yml

old_elasticsearch_data=`cat ${elasticsearch_yml} | grep path.data`
old_elasticsearch_logs=`cat ${elasticsearch_yml} | grep path.logs`

# 需要对变量中的 / 转义成 \/
old_elasticsearch_data=`echo $old_elasticsearch_data | sed 's#\/#\\\/#g'`
old_elasticsearch_logs=`echo $old_elasticsearch_logs | sed 's#\/#\\\/#g'`

if [ ! -d ${installPath}/elasticsearch_data ];then
    mkdir -p ${installPath}/elasticsearch_data
    echo "mkdir -p ${installPath}/elasticsearch_data"
fi
if [ ! -d ${installPath}/elasticsearch_logs ];then
    mkdir -p ${installPath}/elasticsearch_logs
    echo "mkdir -p ${installPath}/elasticsearch_logs"
fi

elasticsearch_data="path.data: ${installPath}/elasticsearch_data"
elasticsearch_logs="path.logs: ${installPath}/elasticsearch_logs"

# 需要对变量中的 / 转义成 \/
elasticsearch_data=`echo $elasticsearch_data | sed 's#\/#\\\/#g'`
elasticsearch_logs=`echo $elasticsearch_logs | sed 's#\/#\\\/#g'`

# 单引号: shell处理命令时,对其中的内容不做任何处理.即此时是引号内的内容是sed命令所定义的格式.
# 双引号: shell处理命令时,要对其中的内容进行算术扩展.如果想让shell扩展后得到sed命令所要的格式,使用双引号即可.
sed -i "s/$old_elasticsearch_data/$elasticsearch_data/g" ${elasticsearch_yml}
sed -i "s/$old_elasticsearch_logs/$elasticsearch_logs/g" ${elasticsearch_yml}
# 允许所有 ip 访问
old_network_host=`cat ${elasticsearch_yml} | grep network.host`
network_host="network.host: 0.0.0.0"
sed -i "s/$old_network_host/${network_host}/g" ${elasticsearch_yml}


# 设置允许跨域
enabled="http.cors.enabled: true"
origin="http.cors.allow-origin: \"*\""
echo ${enabled} >> ${elasticsearch_yml}
echo ${origin} >> ${elasticsearch_yml}
echo "modify ${elasticsearch_yml} file success..."

userdel elasticsearch
userdel elasticsearchgrp
groupdel elasticsearch
groupdel elasticsearchgrp
# 创建用户和属组
groupadd elasticsearchgrp
useradd elasticsearch
usermod -g elasticsearchgrp elasticsearch
usermod -d ${installPath}/elasticsearch elasticsearch
chown -R elasticsearch:elasticsearchgrp ${installPath}/elasticsearch
chown -R elasticsearch:elasticsearchgrp ${installPath}/elasticsearch_data
chown -R elasticsearch:elasticsearchgrp ${installPath}/elasticsearch_logs
echo "chown -R elasticsearch success..."

# 修改 JVM 内存设置
jvm_file=${installPath}/elasticsearch/config/jvm.options
old_Xms=`cat ${jvm_file} | grep -v grep | grep Xms | grep -v '#'`
old_Xmx=`cat ${jvm_file} | grep -v grep | grep Xmx | grep -v '#'`
Xms="-Xms100m"
Xmx="-Xmx100m"
sed -i "s/$old_Xms/$Xms/g" ${jvm_file}
sed -i "s/$old_Xmx/$Xmx/g" ${jvm_file}

# 普通用户无法使用java,首先切换普通用户判断 $PATH 是否有 JAVA_HOME 环境变量,其次判断该路径是否有访问权限
su - elasticsearch -c "cd ${installPath}/elasticsearch/bin && ./elasticsearch &"
echo "start elasticsearch success..."


######################################################################################################################################
# 1.[1]: max file descriptors [4096] for elasticsearch process is too low, increase to at least [65536] 意思是说你的进程不够用了
# 解决方案： 切到root 用户：进入到security目录下的limits.conf；执行命令 vim /etc/security/limits.conf 在文件的末尾添加下面的参数值：
# * soft nofile 65536
# * hard nofile 131072
# * soft nproc 2048
# * hard nproc 4096
# 前面的*符号必须带上，然后重新启动就可以了。执行完成后可以使用命令 ulimit -n 查看进程数
# 2.[2]: max virtual memory areas vm.max_map_count [65530] is too low, increase to at least [262144]  需要修改系统变量的最大值了
# 解决方案：切换到root用户修改配置sysctl.conf  增加配置值： vm.max_map_count=655360
# 执行命令 sysctl -p   这样就可以了，然后重新启动ES服务 就可以了
######################################################################################################################################