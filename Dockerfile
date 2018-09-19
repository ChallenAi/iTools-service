#final stage
FROM alpine:latest
COPY ./iTools-service /app/
WORKDIR /app
ENTRYPOINT ./iTools-service
LABEL Name=itools-service Version=0.0.1
EXPOSE 9215