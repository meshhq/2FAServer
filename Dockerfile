FROM golang:1.8

WORKDIR /go/src/server

COPY . .
 
RUN go install -v

ENV PG_USERNAME="postgres"
ENV PG_PASSWORD="swordfsih987"
ENV PG_DATABASE="2FAuth"

EXPOSE 8080

CMD ["run"]
