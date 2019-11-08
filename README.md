# microservices-in-golang


## 如何运行

0. 构建容器
```shell script
docker-compose build --no-cache
```
注意，目前仍无法使用 `docker-compose run or up`来运行服务，需要手动启用每个服务。

1. 启动 mongoDB
2. 启动 postgres
3. 启动 vessel-service
4. 启动 user-service && 启动 user-cli(如果未获取token的话)
5. 启动 consignment-service
6. 修改 `consignment-cli/Dockefile`的token && 启动 consignment-cli

## TODO

1. 如何用 `docker-compose` 来编排容器
2. 容器会重复启动吗
