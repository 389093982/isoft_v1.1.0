#!/bin/bash

# 5672是项目中连接rabbitmq的端口(我这里映射的是5673),15672是rabbitmq的web管理界面端口(我映射为15673)
# 访问地址 193.112.162.61:15673
docker run -d --name myrabbitmq -p 5673:5672 -p 15673:15672 docker.io/rabbitmq:latest

# docker pull docker.io/rabbitmq:3-management
# tag为3-management时带有web管理界面的,latest不带管理界面,两种都可以正常运行
# 默认账号密码都是guest
docker run -d --name myrabbitmq -p 5673:5672 -p 15673:15672 docker.io/rabbitmq:3-management
