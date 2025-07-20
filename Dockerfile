FROM debian:stable-slim

WORKDIR /root

RUN apt-get update -qq && apt-get install -y apt-transport-https \
                                             ca-certificates \
                                             vim \
                                             curl \
                                             wget \
                                             jq \
                                             unzip \
                                             dnsutils \
                                             traceroute \
                                             telnet \
                                             netcat \
                                             net-tools && \
                                             apt-get clean

CMD [ "/bin/bash" ]