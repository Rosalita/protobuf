# protobuf

A simple example of encoding and decoding protobuf messages.

Why would someone want to do this?
-  Protobuf messages are much smaller than xml and json messages.

1 ) Install `protoc` compiler, which is used to compile `.proto` files.
- Download from [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)
- Add `/bin` to path and verify with `protoc --version`

2 ) Install the Go protobuf plugin with `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

3 ) Write a `foo.proto` file which will contain.
- A message consisting of a collection of typed fields. 
- Each typed field must have a unique tag.
- Include the option `option go_package = "/bar";` 
- This will become the package name which contains auto generated code.

4 ) Compile the `foo.proto` file with `protoc --go_out=. *.proto`
- This generates a `bar.pb.go` file which contains a data type struct.
- The generated struct has helpful methods for `String`, `Reset`, `GetField` etc.

5 ) Write Go code that uses the generated package.

Note when editing `bar.proto` 
- Any changes to field tag numbers will be a breaking change.
- Fields can be deleted.
- New fields can be added but they must use a new unique tag number. 
