#!/bin/bash

########################################################
# 项目所在路径
isoft_iuip_project_path="/root/project/pythonproject/IUIP"
isoft_iuip_project_port="7000"
########################################################

# 杀掉相关端口号对应的进程
sh ../../common/port_kill.sh ${isoft_iuip_project_port}

sleep 5

# 启动项目
# linux shell怎么开启多个进程,在所在的命令或者脚本后面加上&
cd ${isoft_iuip_project_path} && /root/soft/install/python_env/python3.6/bin/uwsgi --http :${isoft_iuip_project_port} --module IUIP.wsgi --static-map=/static=static &
exit;
