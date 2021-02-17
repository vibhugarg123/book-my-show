FROM golang:1.15.7
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o book-my-show .