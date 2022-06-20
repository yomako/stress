FROM golang:onbuild

WORKDIR /app

COPY stress.go ./

RUN go build stress.go

CMD [ "/app/stress" ]

