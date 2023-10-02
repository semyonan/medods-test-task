FROM golang

RUN mkdir /main

ADD . /main

WORKDIR /main

RUN go build -o app ./cmd/app/main.go

EXPOSE 8080
CMD [ "./app" ]