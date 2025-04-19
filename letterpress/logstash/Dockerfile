# 使用最新的 Logstash 镜像
FROM docker.elastic.co/logstash/logstash:8.6.0

# 安装最新的 JDBC 插件
RUN /opt/logstash/bin/logstash-plugin install logstash-integration-jdbc

# 安装 PostgreSQL JDBC 驱动
RUN mkdir -p /opt/logstash/vendor/jdbc && \
    curl -o /opt/logstash/vendor/jdbc/postgresql-42.3.1.jar https://jdbc.postgresql.org/download/postgresql-42.3.1.jar


# 设置 Logstash JVM 配置（根据需要调整）
ENV LOGSTASH_JAVA_OPTS="-Xmx1g -Xms512m -XX:+UseG1GC"

# 设置启动命令
ENTRYPOINT ["/usr/local/bin/docker-entrypoint"]
