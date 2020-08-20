FROM mysql:5.7

RUN apt-get update && apt-get install -y curl \
    && curl -sL https://github.com/k0kubun/sqldef/releases/download/v0.5.14/mysqldef_linux_amd64.tar.gz > mysqldef.tar.gz \
    && tar -xvzf mysqldef.tar.gz -C /bin \
    && rm mysqldef.tar.gz

RUN { \
    echo '[mysqld]'; \
    echo 'character-set-server=utf8mb4'; \
    echo 'collation-server=utf8mb4_general_ci'; \
    echo '[client]'; \
    echo 'default-character-set=utf8mb4'; \
    } > /etc/mysql/conf.d/charset.cnf
