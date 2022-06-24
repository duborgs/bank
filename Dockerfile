FROM mysql
COPY ../src/pkg/db/script.sql /docker-entrypoint-initdb.d

