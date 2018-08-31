#!/bin/bash


read -p "Please input a pulltext : " pulltext

docker pull ${pulltext}
