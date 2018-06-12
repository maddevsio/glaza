FROM golang:latest
RUN mkdir -p /app
ADD . /app
ENV GOPATH /app
WORKDIR /app
CMD ["./test-run.sh"]