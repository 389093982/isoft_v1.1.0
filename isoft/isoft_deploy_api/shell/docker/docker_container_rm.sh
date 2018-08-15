#!/bin/bash

docker rm $(docker ps -aq --filter name=${container_name})
