from alpine:latest
run mkdir /app
copy bin/backendApp /app
cmd [ "/app/backendApp" ]