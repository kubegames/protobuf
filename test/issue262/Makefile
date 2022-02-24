regenerate:
	go install github.com/kubegames/protobuf/protoc-min-version
	go install github.com/kubegames/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/kubegames/protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/kubegames/protobuf/types:. timefail.proto 
