FROM golang

RUN go get github.com/BabetteB/DISYS_MiniProject02

RUN mkdir -p /chittychatapp
WORKDIR /chittychat
COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 3000

CMD [ "./server/server.go" ]