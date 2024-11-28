FROM golang:alpine AS bulider

WORKDIR /app

COPY . .

RUN apk add --no-cache  ca-certificates tzdata build-base musl-dev libc6-compat libgcc libstdc++

# 构建Go项目
RUN go mod tidy && go build

FROM alpine:latest AS runner

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY --from=bulider /app/blog-server ./blog-server
COPY --from=bulider /app/static ./static

EXPOSE 8000 8001

CMD ["./blog-server"]