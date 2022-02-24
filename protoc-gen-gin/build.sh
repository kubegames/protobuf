#!/bin/bash
#protoc --go_out=./ --go_opt=plugins=grpc --go_opt=paths=source_relative --gin_out=./ --gin_opt=paths=source_relative ./proto/*.proto
# protoc --go_out=plugins=grpc:./ --go_opt=paths=source_relative ./google/api/http.proto
# protoc --go_out=plugins=grpc:./ --go_opt=paths=source_relative ./google/api/annotations.proto
# protoc --go_out=plugins=grpc:./ --go_opt=paths=source_relative ./google/validator/validator.proto

# #生成simple.validator.pb.go和simple.pb.go
# protoc --govalidators_out=. --go_out=plugins=grpc:./ --go_opt=paths=source_relative ./example/grpc/proto/simple.proto
# #生成simple.pb.gw.go
# protoc --grpc-gateway_out=logtostderr=true:./ ./example/grpc/proto/simple.proto
#--plugin=protoc-gen-gin=./protoc-gen-gin
#--go_out=plugins=grpc:service
#--plugin=protoc-gen-foo=/path/to/protoc-gen-foo
#protoc --plugin=protoc-gen-gin=./protoc-gen-gin --go_out=./ --go_opt=paths=source_relative --gin_out=./ --gin_opt=paths=source_relative ./proto/v1.proto

#protoc --plugin=protoc-gen-gin=./protoc-gen-gin --go_out=./ --go_opt=plugins=grpc --go_opt=paths=source_relative --gin_out=./ --gin_opt=paths=source_relative ./proto/*.proto

protoc --go_out=./ --go_opt=plugins=grpc --go_opt=paths=source_relative ./proto/*.proto
protoc --go_out=./ --go_opt=plugins=grpc --go_opt=paths=source_relative ./proto/v2/*.proto

protoc --plugin=protoc-gen-gin=./protoc-gen-gin --go_out=./ --gin_out=./ --gin_opt=paths=source_relative ./proto/*.proto