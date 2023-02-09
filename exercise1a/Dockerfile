FROM ubuntu:18.04

RUN apt-get update -y \
  && apt-get install -y git gcc make golang-1.10  

COPY netskope-cert-bundle.pem /etc/ssl/certs/  
RUN echo netskope-cert-bundle.pem >> /etc/ca-certificates.conf
RUN update-ca-certificates

ENV GOROOT /usr/lib/go-1.10
ENV PATH $GOROOT/bin:$PATH
ENV GOPATH /root/go
ENV APIPATH /root/go/src/api

WORKDIR $APIPATH
COPY . .

RUN \ 
  go get -d -v \
  && go install -v \
  && go build

RUN rm api.go

EXPOSE 3000
CMD ["./api"]
