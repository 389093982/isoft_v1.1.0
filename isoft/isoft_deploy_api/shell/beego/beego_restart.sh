#!/bin/bash

# 应用名称
app_name=$1

sh ./beego_shutdown.sh ${app_name} && sh ./beego_startup.sh ${app_name}
