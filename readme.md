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

# make sure verison of protoc-gen-go plugin matched with protoc
protoc-gen-go --version  && protoc --version
```

##  install go grpc service code generator. Generate *_grpc.pb.go : gRPC server and client interfaces. command `protoc --go-grpc_out=`
```
# will be installed into $(go env GOPATH)/bin
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## install go http service code generator. Generate *.pb.gw.go : GRPC-Gateway reverse-proxy definitions Http server. command `protoc --grpc-gateway_out=`
```
# will be installed into $(go env GOPATH)/bin
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.6

# also install open api doc swagger
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
```


## setting up google annotation proto
```
cd ~ && \
git clone https://github.com/googleapis/googleapis/tree/master/google/api && \
sudo cp -a ~/googleapis/google/api /usr/local/include/google/api
```

Configure IDE import path to include google annotation proto

![image](https://user-images.githubusercontent.com/6024289/226167761-67571647-14f8-4083-a915-19fc99541a27.png)

## build pb, run code generator
```
protoc -I. -I/usr/local/include -I=$GOPATH/src \
--go_out=$GOPATH/src \
--go-grpc_out=$GOPATH/src \
--grpc-gateway_out=logtostderr=true:$GOPATH/src \
./grpc/proto/example/example.proto
```

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

##  install go grpc service code generator. Generate *_grpc.pb.go : gRPC server and client interfaces. command `protoc --go-grpc_out=`
```
# will be installed into $(go env GOPATH)/bin
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## install go http service code generator. Generate *.pb.gw.go : GRPC-Gateway reverse-proxy definitions Http server. command `protoc --grpc-gateway_out=`
```
# will be installed into $(go env GOPATH)/bin
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.6

# also install open api doc swagger
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
```


## setting up google annotation proto
```
cd ~ && \
git clone https://github.com/googleapis/googleapis/tree/master/google/api && \
sudo cp -a ~/googleapis/google/api /usr/local/include/google/api
```

## build pb, run code generator
```
protoc -I. -I/usr/local/include -I=$GOPATH/src \
--go_out=$GOPATH/src \
--go-grpc_out=$GOPATH/src \
--grpc-gateway_out=logtostderr=true:$GOPATH/src \
./grpc/proto/example/example.proto
```

## Generated Files

![image](https://user-images.githubusercontent.com/6024289/226167633-32f7eac1-ff12-4535-aed5-ba603ecacde2.png)

