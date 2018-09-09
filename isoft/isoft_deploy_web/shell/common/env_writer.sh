#!/bin/bash

# 远程ssh无法获取环境变量问题分析
# ssh登录时要发送本地环境变量到远程服务器上,更改远程服务器上的/etc/ssh/sshd_config,里面有AcceptEnv xxx的配置,注释掉就不会接收ssh发送过来的变量了

# 本机使用 su 命令切换到普通用户(属于 Login 方式)
# Login 之前,系统 PATH 为：/usr/local/bin:/bin:/usr/bin
# Login 方式,文件调用顺序为： /etc/profile -> /etc/bashrc -> ~/.bashrc -> ~/.bash_profile
# Login 之后,系统 PATH 为：/usr/local/bin:/bin:/usr/bin:/usr/local/sbin:/usr/sbin:/sbin:/home/user01/bin

# 远程机使用 ssh 命令以普通用户身份连接到主机执行获取 PATH 的命令(属于 NoLogin 方式)
# NoLogin 方式,命令获取的 PATH 为该远程机的,并未拿到目标主机的 PATH
# NoLogin 方式,文件调用顺序为:/etc/bashrc -> ~/.bashrc
# NoLogin 方式,目标主机 User 用户 PATH 为:/usr/local/bin:/bin:/usr/bin

# 综上,如需修改 PATH,建议修改 bashrc 文件,从而保证任何方式访问时 PATH 的正确性

################################################################################################################
# 环境变量名
env_name=$1
# 环境变量值
env_value=$2
################################################################################################################


function write_or_modify_env(){
    env_name=$1
    env_value=$2
    env_file_path=$3
    count=`cat ${env_file_path} |grep export | grep ${env_name}=| wc -l`

    if [ "${count}" == "1" ];then
        echo "${env_name} is exist"
        # 查找行号并进行删除
        search_row_num=`cat -n ${env_file_path} |grep export | grep ${env_name}= | awk '{print $1}'`
        # sed 中变量转义使用 '$(echo $search_row_num)' 格式
        # 根据行号进行删除
        sed -i ''$(echo $search_row_num)'d' ${env_file_path}
        # 在文件末尾追加一行
        echo "export ${env_name}=${env_value}" >> ${env_file_path} && source ${env_file_path}
        echo "modify success..."
    else
        if [ "${count}" == "0" ];then
            # 在文件末尾追加一行
            echo "export ${env_name}=${env_value}" >> ${env_file_path} && source ${env_file_path}
            echo "write success..."
        else
            echo "multi row was found..."
        fi
    fi
}

write_or_modify_env ${env_name} ${env_value} "/etc/profile"
write_or_modify_env ${env_name} ${env_value} "/etc/bashrc"



