#!/bin/bash

########################################################
virtual_env_parent_path=$1
virtual_env_name=$2
########################################################
cd ${virtual_env_parent_path} && /usr/local/python3/bin/virtualenv ${virtual_env_name}
