#!/bin/bash


read -p "Please input a container_name : " container_name

docker kill --signal=SIGINT ${container_name}
