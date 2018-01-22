FROM golang:latest 

ENV PG_USERNAME="postgres"
ENV PG_PASSWORD=""
ENV PG_DATABASE="2FactorAuth"
ENV PG_HOSTNAME="127.0.0.1"
ENV PG_PORT="32768"
ENV GOPATH="/go"
ENV DUBUG="true"

RUN mkdir $GOPATH/src/2FAServer 
WORKDIR $GOPATH/src/2FAServer
ADD . .

RUN go get -v ./... 
RUN go get -u github.com/jteeuwen/go-bindata/...

RUN make install 

EXPOSE 1323

CMD ["2FAServer"]
