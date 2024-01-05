FROM golang:1.21-alpine

COPY discord-bot-2 .

RUN chmod u+x discord-bot-2

CMD ["./discord-bot-2"]