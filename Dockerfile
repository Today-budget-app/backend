FROM golang:1.21.6-alpine3.7

COPY . /app/today
WORKDIR /app/today

RUN go build -o today

EXPOSE 8080

CMD [ "today" ]