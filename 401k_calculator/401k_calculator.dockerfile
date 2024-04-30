from alpine:latest
run mkdir /app
copy bin/401kCalculatorApp /app
cmd [ "/app/401kCalculatorApp" ]