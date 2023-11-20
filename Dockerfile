FROM golang
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o xpense-tracker cmd/web/*.go
RUN mkdir /sqlite
CMD ["sh", "-c", "/app/xpense-tracker"]