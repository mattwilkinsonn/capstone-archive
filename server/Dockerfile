FROM golang AS builder

ENV GO111MODULE=on

WORKDIR /builder

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN mkdir -p out
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./out ./cmd/server.go


FROM scratch as runtime
COPY --from=builder /builder/out /app
ENV GIN_MODE=release
ENV PORT=8080
ENV ENVIRONMENT=release

EXPOSE 8080
ENTRYPOINT ["/app/server"]