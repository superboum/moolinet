FROM debian:latest

RUN adduser --system moolinet && \
    apt-get update && \
    apt-get install -y xvfb libxrender1 libxtst6 libxi6 default-jdk wget && \
    wget http://download.processing.org/processing-3.3.6-linux64.tgz -O processing.tgz && \
    tar xzf processing.tgz && \
    rm processing.tgz && \
    mv processing-3.3.6 processing

COPY ./moolinet-write /usr/bin

WORKDIR /home/moolinet
USER moolinet
