FROM alpine:latest

WORKDIR /app

COPY bin/car_server /app/car_server

EXPOSE 3000

ENTRYPOINT [ "/app/car_server" ]