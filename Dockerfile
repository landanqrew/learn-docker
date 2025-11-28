FROM debian:stable-slim

ENV PORT=8991

COPY goserver /bin/goserver

CMD ["/bin/goserver"]

# docker run -p 8010:8010 goserver