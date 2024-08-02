from --platform=linux/amd64 alpine:latest
run mkdir /app
copy bin/linkServiceApp /app
cmd [ "/app/linkServiceApp" ]