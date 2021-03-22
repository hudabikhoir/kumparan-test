FROM golang:alpine as builder

RUN mkdir -p /kumparan/backend 
ADD . /kumparan/backend/
WORKDIR /kumparan/backend 

RUN mkdir -p /kumparan/backend/public/export/vital_sign_log/
RUN mkdir -p /kumparan/backend/public/export/statistic_report/
RUN mkdir -p /kumparan/backend/public/export/inpatient/
RUN mkdir -p /kumparan/backend/public/upload/

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN export GOPROXY=https://proxy.golang.org \
    && go mod download

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main app/server/main.go
RUN GOOS=linux GOARCH=amd64 \
    go build -ldflags="-w -s" \
    -o main \
    app/main.go

FROM alpine

RUN addgroup -S appgroup && adduser -S -D -H -h /app appuser -G appgroup

COPY --from=builder /kumparan/backend/main /app/
COPY --from=builder /kumparan/backend/config/ /app/config/

USER appuser

WORKDIR /app

STOPSIGNAL SIGINT

EXPOSE 1323

ENTRYPOINT ["./main"]