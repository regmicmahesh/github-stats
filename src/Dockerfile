FROM golang:1.16-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o github-stats

FROM alpine
COPY --from=build /app/github-stats /github-stats 
ENTRYPOINT ["/github-stats"]
