FROM golang:1.19.4-bullseye

WORKDIR /tmp
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
  apt-get -y upgrade && \
  apt-get install -y wget gcc g++ make sqlite3 && \
  wget -q https://dev.mysql.com/get/mysql-apt-config_0.8.22-1_all.deb && \
  apt-get -y install ./mysql-apt-config_*_all.deb && \
  apt-get -y update && \
  apt-get -y install mysql-client

RUN useradd --uid=1001 --create-home kid
USER kid

RUN mkdir -p /home/kid/webapp/go
WORKDIR /home/kid/webapp/go
COPY --chown=kid:kid ./ /home/kid/webapp/go/

ENV GOPATH=/home/kid/tmp/go
ENV GOCACHE=/home/kid/tmp/go/.cache

CMD ["go", "run", "./cmd/kid/main.go"]