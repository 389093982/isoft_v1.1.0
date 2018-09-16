#!/bin/bash

# docker pull elasticsearch:2.4.4

# 9200端口:ES节点和外部通讯使用
# 9300端口:ES节点之间通讯使用
# http://193.112.162.61:9200/
docker run -d -p 9200:9200 -p 9300:9300 --name elasticsearch elasticsearch:2.4.4


ES配置文件elasticsearch.yml加入2个参数，并重启


http.cors.enabled: true
http.cors.allow-origin: "*"

# 创建索引
# curl 193.112.162.61:9200/metadata -XPUT -d'{"mappings":{"objects":{"properties":{"name":{"type":"string","index":"not_analyzed"},"version":{"type":"integer"},"size":{"type":"integer"},"hash":{"type":"string"}}}}}'

# 通过如下语句,列出所有索引
# curl '193.112.162.61:9200/_cat/indices?v'


###################################################################################################
docker pull mobz/elasticsearch-head:5

docker run -d -p 9100:9100 docker.io/mobz/elasticsearch-head:5