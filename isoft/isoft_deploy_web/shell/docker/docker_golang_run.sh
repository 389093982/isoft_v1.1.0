#!/bin/bash

# 在创建容器的时候加了 --rm 参数,所以当我们执行 exit 命令退出容器的时候,这个临时容器会被删除
docker run -it --rm golang bash
