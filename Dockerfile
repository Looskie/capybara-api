FROM golang:latest AS BUILDER

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go build -o main .


FROM alpine:latest AS RUNNER

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=BUILDER ./app/main /bruh/
COPY --from=BUILDER ./app/capys /capys

CMD ["/bruh/main"]
