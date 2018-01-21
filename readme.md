# Automatically generated golang protobuf bindings for TensorFlow.

To regenerate:
```
cd $GOPATH/src/github.com/tensorflow/tensorflow
protoc --go_out ~/go/src/github.com/is8ac/tfpb tensorflow/core/protobuf/!(*service.proto)
protoc --go_out ~/go/src/github.com/is8ac/tfpb tensorflow/core/framework/*.proto
```
