#!/bin/bash

###################################################################################################################
# mysql root 密码
rootPwd=$1
###################################################################################################################

if [ -z ${rootPwd} ];then
    echo "invalid params"
    exit;
fi

# 修改加密规则为 mysql_native_password
command1="alter user 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'Mysql@123';flush privileges;"
# 修改 root 密码
command2="alter user 'root'@'%' IDENTIFIED BY '${rootPwd}';flush privileges;"
command3="alter user 'root'@'localhost' IDENTIFIED BY '${rootPwd}';flush privileges;"

echo ${command1}${command2}${command3}

mysql -uroot -pMysql@123 -e "${command1}${command2}${command3}"
