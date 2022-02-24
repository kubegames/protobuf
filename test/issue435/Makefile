regenerate:
	go install github.com/kubegames/protobuf/protoc-gen-gogo
	protoc --gogo_out=. --proto_path=../../../../../:../../protobuf/:. *.proto
