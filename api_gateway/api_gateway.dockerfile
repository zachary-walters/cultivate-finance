from alpine:latest
run mkdir /app
copy bin/apiGateway /app
cmd [ "/app/apiGateway" ]