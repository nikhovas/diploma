#!/bin/bash

python -m grpc_tools.protoc -I../../proto/servers/Control --python_out=../control_grpc_client \
    --grpc_python_out=../control_grpc_client ../../proto/servers/Control/Control.proto

python -m grpc_tools.protoc -I../../proto/servers/TelegramStaffBot --python_out=../grpc_server \
    --grpc_python_out=../grpc_server ../../proto/servers/TelegramStaffBot/TelegramStaffBot.proto
