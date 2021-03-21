FROM golang:1.16-alpine AS build

WORKDIR /go
COPY . link/
RUN cd link \
    && go build -o go-link \
    && mv ./go-link /usr/bin/ \
    && rm -rf /go/link

CMD ["go-link"]
