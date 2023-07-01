FROM golang:1.20-mini

ENV APPLICATION_NAME=soat1_challenge1

ENV APPLICATION_PACKAGE=./cmd
ENV APPLICATION_WORKING_DIR=app

# Setup and install MySQL database for testing
RUN echo "mysql-community-server mysql-community-server/root-pass password root" | debconf-set-selections; \
    echo "mysql-community-server mysql-community-server/re-root-pass password root" | debconf-set-selections

RUN apt-get update && apt-get install -y gnupg2 && echo 'deb http://repo.mysql.com/apt/ubuntu/ bionic mysql-8.0' >> /etc/apt/sources.list.d/mysql.list

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 467B942D3A79BD29 && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get -y install mysql-server

# Install netcat for database status check
RUN apt-get install -y netcat

ENV DATABASE_HOST localhost
ENV DATABASE_PASSWORD admin
ENV DATABASE_NAME restaurante

COPY sql/*.sql /docker-entrypoint-initdb.d/

COPY migrations/sql/*.sql /docker-entrypoint-initdb.d/
