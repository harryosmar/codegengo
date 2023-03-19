## install protoc
```
brew install protobuf

protoc --version
```

## don't forget add $(go env GOPATH) into ENV PATH variable 
```
export PATH="$PATH:$(go env GOPATH)/bin"
```

## install golang pb code generator. Generate *.pb.go : service definition and message types. command `protoc --go_out=`
```
# will be installed into $(go env GOPATH)/bin
go install google.golang.org/protobuf/cmd/protoc-gen-go

# make sure verison of protoc go plugin matched with protoc
protoc-gen-go --version  && protoc --version
```

## # install go grpc service code generator. Generate *_grpc.pb.go : gRPC server and client interfaces. command `protoc --go-grpc_out=`
```
# will be installed into $(go env GOPATH)/bin
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## install go http service code generator. Generate *.pb.gw.go : GRPC-Gateway reverse-proxy definitions Http server. command `protoc --grpc-gateway_out=`
```
# # will be installed into $(go env GOPATH)/bin
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.6
```


## setting up google annotation proto
```
cd ~ && \
git clone https://github.com/googleapis/googleapis/tree/master/google/api && \
sudo cp -a ~/googleapis/google/api /usr/local/include/google/api
```

## build pb, run code generator
```
protoc -I. -I/usr/local/include \
--go_out=./ --go_opt=paths=source_relative \
--go-grpc_out=./ --go-grpc_opt=paths=source_relative \
--grpc-gateway_out=logtostderr=true:./ --grpc-gateway_opt=paths=source_relative \
./proto/example/example.proto
```

