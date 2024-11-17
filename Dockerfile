FROM golang:1.23.3-bookworm as buildx

WORKDIR /app

COPY . .

RUN go build -o /server

FROM gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=buildx /server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]



