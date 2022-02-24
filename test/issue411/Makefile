regenerate:
	go install github.com/kubegames/protobuf/protoc-gen-gogo
	protoc-min-version --version="3.0.0" --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/kubegames/protobuf/types:. --proto_path=../../../../../:../../protobuf/:. issue411.proto
