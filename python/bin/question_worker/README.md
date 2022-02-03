```
python -m grpc_tools.protoc -I../../proto/servers/QuestionWorker --python_out=. --grpc_python_out=. ../../proto/servers/QuestionWorker/QuestionWorker.proto
```

```
python -m grpc_tools.protoc -I../../proto/servers/QuestionWorker --python_betterproto_out=. --grpc_python_betterproto_out=. ../../proto/servers/QuestionWorker/QuestionWorker.proto
```

```
python -m grpc_tools.protoc -I../../proto/servers/QuestionWorker --python_betterproto_out=. ../../proto/servers/QuestionWorker/QuestionWorker.proto
```