FROM golang:1.17

RUN apt update \
  && apt install -y vim

ENV GO111MODULE on
RUN mkdir /app
WORKDIR /app

# install go tools
RUN go get github.com/gin-gonic/gin \
  && go get github.com/go-sql-driver/mysql \ 
  && go get github.com/volatiletech/sqlboiler/v4@latest \
  && go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest