# microservices-in-golang


## 如何运行

0. 构建容器
```shell script
docker-compose build --no-cache
```
注意，目前仍无法使用 `docker-compose run or up`来运行服务，需要手动启用每个服务。

1. 启动 mongoDB(如果不看日志可以略, vessel-service依赖于这个服务)
2. 启动 postgres(如果不看日志可以略, user-service依赖于这个服务)
3. 启动 vessel-service
4. 启动 user-service && 启动 user-cli(如果未获取token的话)
5. 启动 consignment-service
6. 修改 `consignment-cli/Dockefile`的token && 启动 consignment-cli

## TODO

1. 如何用 `docker-compose` 来编排容器
2. 容器会重复启动吗
3. 如何设计错误处理，特别是调用链非常长的情况下
4. ctx 是如何传递的(在token为空情况下这个bug竟然一路传递到了user-service上)
