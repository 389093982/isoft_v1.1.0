#!/bin/bash

ID=`ps -ef | grep sshd | grep root | grep notty | grep -v "grep" | awk '{print $2}'`
echo $ID
echo "---------------"
for id in $ID
do
kill -9 $id
echo "killed $id"
done