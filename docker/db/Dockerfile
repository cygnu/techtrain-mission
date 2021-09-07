FROM mysql:8.0

COPY ./init/init.sql /docker-entrypoint-initdb.d/init.sql
COPY ./my.cnf /etc/mysql/my.cnf
