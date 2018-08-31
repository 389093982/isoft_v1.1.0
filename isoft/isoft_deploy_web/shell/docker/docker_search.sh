#!/bin/bash


read -p "Please input a searchtext : " searchtext

docker search ${searchtext}
