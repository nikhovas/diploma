#!/bin/bash


join_by_char() {
    local IFS=" "
    echo "$*"
}


proto_files=("common" "consumer_actions" "consumer_bot" "controller" "question_worker" "staff_bot" "vk_products_updater")
joined=$(join_by_char "${proto_files[@]}")


script_dir=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
project_base_dir=$script_dir/..
proto_path=$project_base_dir/proto
go_proto_path=$project_base_dir/go/lib/proto
python_proto_path=$project_base_dir/python/lib/proto/proto


cd $script_dir
for file_name in ${proto_files[@]}
do
	protoc --proto_path=$proto_path \
        --go_out=$go_proto_path/$file_name \
        --go_opt=paths=source_relative \
        --go-grpc_out=$go_proto_path/$file_name \
        --go-grpc_opt=paths=source_relative \
        $proto_path/$file_name.proto
    
    python3 -m grpc_tools.protoc \
        --proto_path=$proto_path \
        --python_out=$python_proto_path/$file_name \
        --grpc_python_out=$python_proto_path/$file_name \
        $proto_path/$file_name.proto
       
    file_to_change="$python_proto_path/$file_name/${file_name}_pb2.py"
    python fix_files.py $file_to_change $joined
    file_to_change="$python_proto_path/$file_name/${file_name}_pb2_grpc.py"
    python fix_files.py $file_to_change $joined
done
