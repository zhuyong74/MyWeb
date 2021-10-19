FROM ubuntu
WORKDIR /app
COPY myweb /app/myweb
EXPOSE 80
ENTRYPOINT ["/app/myweb"]
