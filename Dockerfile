FROM alpine:latest
RUN apk update
RUN apk add chromium-chromedriver
RUN apk add --no-cache chromium
RUN apk add openssh
RUN apk add ca-certificates
RUN apk add openssl
RUN apk add musl
RUN apk add terraform
#RUN cp chromedriver /usr/bin/chromedriver
RUN apk add libc-dev
RUN apk add gcc
RUN apk add go
RUN apk add git
RUN go get github.com/sclevine/agouti

RUN mkdir /app
ADD . /app/
WORKDIR /app/
RUN mv terraform /usr/bin/terraform
RUN go build -o mkcaasp caasp/*.go
