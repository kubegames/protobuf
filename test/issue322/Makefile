regenerate:
	go install github.com/kubegames/protobuf/protoc-gen-gogofast
	protoc-min-version --version="3.0.0" --gogofast_out=Mgoogle/protobuf/descriptor.proto=github.com/kubegames/protobuf/protoc-gen-gogo/descriptor:. --proto_path=../../../../../:../../protobuf/:. *.proto
