from --platform=linux/amd64 alpine:latest
run mkdir /app
copy bin/apiGateway /app
cmd [ "/app/apiGateway" ]