# 使用官方的MySQL镜像作为基础镜像
FROM mysql:8.0

# 设置环境变量，例如MySQL的root用户的密码
ENV MYSQL_ROOT_PASSWORD root

# 将数据库初始化脚本添加到容器的/docker-entrypoint-initdb.d目录
# 这个脚本会在MySQL服务启动时执行
ADD schema.sql /docker-entrypoint-initdb.d

# 暴露MySQL服务的端口
EXPOSE 3306

# 设置容器启动时执行的命令
CMD ["mysqld"]
