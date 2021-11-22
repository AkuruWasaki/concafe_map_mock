FROM golang:1.17

RUN apt update \
  && apt install -y vim

RUN mkdir /app
WORKDIR /app

COPY . ./
RUN go mod download
RUN go build .

CMD [ "./concafe_map_mock" ]