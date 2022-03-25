FROM golang:1.17-alpine

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o wb-l0 .

CMD [ "/app/wb-l0" ]
