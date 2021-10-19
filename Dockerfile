FROM ubuntu
WORKDIR /app
COPY myweb /app/myweb
EXPOSE 8000
ENTRYPOINT ["/app/myweb"]
