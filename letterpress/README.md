# 将 Elasticsearch、Logstash 和 Kibana 与 Go 应用程序结合

> 结合 Elasticsearch 和 Golang 和 Logstash 和 Kibana 的强大功能。我们将构建一个基本的内容管理系统，该系统能够创建、读取、更新和删除帖子，以及通过 Elasticsearch 搜索帖子的能力

## 依赖

- https://github.com/lib/pq
- github.com/elastic/go-elasticsearch
- https://gin-gonic.com/
- https://github.com/rs/zerolog

## 安装依赖

```shell
go get github.com/lib/pq \
go get github.com/elastic/go-elasticsearch/v9@latest \
go get -u github.com/gin-gonic/gin \
go get -u github.com/rs/zerolog/log
```

## run

```shell
docker-compose up
```
