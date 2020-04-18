FROM ubuntu:18.04

RUN apt-get update && apt-get install -y curl

COPY file_server /
COPY client /
COPY fileserver/audio /audio
EXPOSE 8088
ENTRYPOINT ["sh", "-c", "./file_server"]