#!/bin/bash


read -p "Please input a container_name : " container_name

docker logs ${container_name}
