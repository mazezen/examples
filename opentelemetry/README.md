### Go + OpenTelemetry + Jaeger

#### 安装Jaeger web UI
```shell
docker run --rm --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 4317:4317 \
  -p 4318:4318 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.55

```

#### OR
```shell
docker-compose -f opentelemetry.yaml up -d
```

#### 使用
 * echo
 * gin
 * zap
 * gorm
 * http
 * redis
 * zap
 * grpc

#### 效果
![img.png](./images/img.png)

#### span 嵌套效果
![img.png](./images/img2.png)