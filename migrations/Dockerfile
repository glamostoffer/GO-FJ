FROM postgres:alpine

EXPOSE 5432

ADD *.up.sql /docker-entrypoint-initdb.d

RUN chmod +x /docker-entrypoint-initdb.d/*.up.sql