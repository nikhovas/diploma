#!/bin/bash

join_by_char() {
    local IFS=" "
    echo "$*"
}

proto_files=("common" "consumer_actions" "consumer_bot" "controller" "question_worker" "staff_bot")
joined=$(join_by_char "${proto_files[@]}")

script_dir=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
project_base_dir=$script_dir/..
proto_path=$script_dir/../proto

cd $proto_path
for file_name in ${proto_files[@]}
do
	protoc --proto_path=./ \
           --go_out=../go/lib/proto/$file_name \
           --go_opt=paths=source_relative \
           --go-grpc_out=../go/lib/proto/$file_name \
           --go-grpc_opt=paths=source_relative \
           $file_name.proto
done

cd $project_base_dir/python
for file_name in ${proto_files[@]}
do
       python3 -m grpc_tools.protoc \
               --proto_path=../proto \
               --python_out=./lib/proto/proto/$file_name \
               --grpc_python_out=./lib/proto/proto/$file_name \
               ../proto/$file_name.proto
       
       file_to_change="./lib/proto/proto/$file_name/${file_name}_pb2.py"
       python $proto_path/fix_files.py $file_to_change $joined
       file_to_change="./lib/proto/proto/$file_name/${file_name}_pb2_grpc.py"
       python $proto_path/fix_files.py $file_to_change $joined
done
