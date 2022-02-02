#!/bin/bash

protoc --go_out=../grpc/control --go_opt=paths=source_relative -I=../../proto/servers/Control --go-grpc_out=../grpc/control \
    --go-grpc_opt=paths=source_relative ../../proto/servers/Control/Control.proto

protoc --go_out=../grpc/questionWorker --go_opt=paths=source_relative -I=../../proto/servers/QuestionWorker --go-grpc_out=../grpc/questionWorker \
    --go-grpc_opt=paths=source_relative ../../proto/servers/QuestionWorker/QuestionWorker.proto

protoc --go_out=../grpc/actionEvent --go_opt=paths=source_relative -I=../../proto/data/actionEvent \
    ../../proto/data/actionEvent/ActionEvent.proto

protoc --go_out=../grpc/actionEvent --go_opt=paths=source_relative -I=../../proto/data/actionEvent \
    ../../proto/data/actionEvent/ActionEvent.proto