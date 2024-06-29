FROM golang:1.22.4

WORKDIR /src/

COPY ./src /src/

RUN go mod tidy

RUN go install github.com/cosmtrek/air@v1.52.1 

CMD [ "air" ]