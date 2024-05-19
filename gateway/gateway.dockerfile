from alpine:latest
run mkdir /app
copy bin/gateway /app
cmd [ "/app/gateway" ]