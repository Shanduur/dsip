all: install clean proto gtest

proto:
	protoc \
	--proto_path=pkg/proto \
	--go_out=plugins=grpc:. \
	pkg/proto/*.proto

clean:
	rm pkg/pb/*.pb.go
	rm */**/*.tmp.*

gtest:
	go test -cover -race ./...

ptest: proto gtest

install:
	cd $(GOPATH)/src/gocv.io/x/gocv && $(MAKE) install

PATH=$(go env GOPATH")
githubActionsSetup:
	go env GOPATH
	cd $(PATH)/src/gocv.io/x/gocv && $(MAKE) install