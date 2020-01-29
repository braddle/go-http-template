# BUILD
FROM golang:1 as build

WORKDIR /service
ADD . /service

RUN go build -o /http-service .

CMD /http-service

# TEST
FROM build as test
