from --platform=linux/amd64 alpine:latest
run mkdir /app
copy bin/debtSnowballCalculatorApp /app
cmd [ "/app/debtSnowballCalculatorApp" ]