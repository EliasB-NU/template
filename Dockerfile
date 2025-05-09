FROM ubuntu:latest
LABEL authors="EliasBraun"

ENTRYPOINT ["top", "-b"]