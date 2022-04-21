FROM golang:latest

RUN apt update && apt install -y netcat

RUN go install github.com/canthefason/go-watcher/cmd/watcher@latest

WORKDIR /go/src/github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge

COPY . .

RUN go mod download && go mod verify

ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for && chmod -R +x ./scripts


ENTRYPOINT ["sh", "./scripts/entrypoint.sh"]