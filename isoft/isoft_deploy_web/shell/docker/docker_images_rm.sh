#!/bin/bash


read -p "Please input a image_name : " image_name
# -f 强制删除
# --no-prune 不移除该镜像的过程镜像,默认移除
docker rmi ${image_name}

