FROM golang:1.15.2

RUN apt-get update && \
    apt-get install -y apt-utils \
    zip \
    unzip \
    vim \
    curl \
    netcat

WORKDIR  /go/src/github.com/BharathKumarRavichandran/k8s-playground/server
COPY . .

EXPOSE 8000

CMD ["./scripts/app_entry.sh"]
