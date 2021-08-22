docker run -p 9200:9200 -p 9300:9300 --name es --net myes -e "discovery.type=single-node"  -e ES_JAVA_OPTS="-Xms64m -Xmx512m" elasticsearch:7.13.4
