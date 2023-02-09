FROM ubuntu:18.04 

WORKDIR /root
COPY . .

RUN adduser --shell /bin/bash --disabled-password --gecos "" --uid 1001 customer
USER 'customer'
RUN whoami

