# syntax=docker/dockerfile:1

FROM golang:1.19-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /plate-reader
EXPOSE 11111
CMD [ "/plate-reader" ]