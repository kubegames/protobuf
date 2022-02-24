regenerate:
	go install github.com/kubegames/protobuf/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=\
	Mtest_proto/test.proto=github.com/kubegames/protobuf/proto/test_proto,\
	Mgoogle/protobuf/any.proto=github.com/kubegames/protobuf/types:. \
	--proto_path=../../protobuf:../:. proto3.proto
	
