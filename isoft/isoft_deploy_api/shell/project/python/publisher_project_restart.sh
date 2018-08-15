#/bin/bash

# 只用一条命令查找并结束Linux进程
# 最后一行会把grep命令也列为结果,使用-v选项进行排除
process=`ps -ef | grep "/root/project/pythonproject/Publisher" | grep -v grep | awk '{print $2}'`
echo "process is ${process}"
kill -9 ${process}

/root/soft/install/python_env/python3.6/bin/python3.6 /root/project/pythonproject/Publisher/main.py --env=yun

exit;
