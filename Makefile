VERSION=t0.0.1
OPTION=--platform linux/amd64

.PHONY: all test clean

grpc-client:
	docker build grpc-client -t myoon/grpc-client:${VERSION} ${OPTION}
	docker push myoon/grpc-client:${VERSION}

grpc-server:
	docker build grpc-server -t myoon/grpc-server:${VERSION} ${OPTION}
	docker push myoon/grpc-server:${VERSION}

http-client:
	docker build http-client -t myoon/http-client:${VERSION} ${OPTION}
	docker push myoon/http-client:${VERSION}

http-server:
	docker build http-server -t myoon/http-server:${VERSION} ${OPTION}
	docker push myoon/http-server:${VERSION}

all: grpc-client grpc-server http-client http-server

