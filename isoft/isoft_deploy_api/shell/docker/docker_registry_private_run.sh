#!/bin/bash

#########################################################################################################
# docker pull registry		# 下载镜像
# 默认情况下,会将仓库存放于容器内的/tmp/registry目录下,这样如果容器被删除,则存放于容器中的镜像也会丢失,
# 所以我们一般情况下会指定本地一个目录挂载到容器内的/tmp/registry下
# docker run -d -p 5000:5000 registry
#########################################################################################################

docker run -d -p 5000:5000 -v /root/soft/install/docker/registry:/tmp/registry registry