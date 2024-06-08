FROM golang:alpine

WORKDIR /data
COPY . .

RUN go build -o benchit .

ENTRYPOINT [ "/data/benchit" ]