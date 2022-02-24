regenerate:
	go install github.com/kubegames/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/kubegames/protobuf/protobuf/ \
	--gogo_out=. events.proto 
