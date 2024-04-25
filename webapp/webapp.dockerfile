from alpine:latest
run mkdir /app
run mkdir /assets
copy bin/webapp /app
copy cmd/assets/ /assets
cmd [ "/app/webapp" ]