FROM ubuntu:20.04

ENV DEFAULT_CODE all
ENV BIN_DIR /opt/bin
ENV PKG_DIR /tmp/pkg
ENV GOBIN=/root/go/bin
ENV PATH=$PATH:$GOBIN
ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8

COPY pkg/* ${PKG_DIR}/

RUN apt-get update && apt-get install -y software-properties-common
RUN add-apt-repository ppa:longsleep/golang-backports

RUN apt-get -y upgrade && apt-get install -y \
    $(cat ${PKG_DIR}/apt_packages.txt)

RUN python3 -m pip install --upgrade pip && \
    python3 -m pip install --upgrade -r ${PKG_DIR}/pip_requirements.txt

RUN unzip ${PKG_DIR}/protoc-3.6.1-linux-x86_64.zip -d /usr/local
RUN go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

WORKDIR ${BIN_DIR}
CMD ["python3", "build.py", "-h"]
