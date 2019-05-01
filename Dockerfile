FROM golang:latest

WORKDIR /app
COPY ./src src/

RUN go get -d ./src/...
RUN go build -o ./dist/crocs-bot.exe ./src/main.go

RUN mv ./src/config.json ./dist/config.json

WORKDIR /app/dist

CMD [ "./crocs-bot.exe" ]