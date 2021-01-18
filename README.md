# grpc-mock
grpc-mock is yet another grpc mock server inspired by [GripMock](https://github.com/tokopedia/gripmock).

grpc-mock implements the grpc mock server by reflections of `.proto` files instead of code generation that [GripMock](https://github.com/tokopedia/gripmock) does.

# Installation
## From Source
Use `go` tool to install:
```shell
go get github.com/monlabs/grpc-mock/cmd/gmock
```
# Usage
## Quickstart
**Step 1:** Write your proto file which should contains at least one `service`.
> helloworld.proto
```protobuf
syntax = "proto3";

package helloworld;

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```
**Step 2:** Write the stub files in `JSON` format for the services defined in the proto file.
> stubs/helloworld.json
```JSON
[
  {
    "service": "helloworld.Greeter",
    "method": "SayHello",
    "in": {
      "equals": {
        "name": "hi"
      }
    },
    "out": {
      "data": {
        "message": "lemon"
      }
    }
  },
  {
    "service": "helloworld.Greeter",
    "method": "SayHello",
    "in": {
      "matches": {
        "name": "^[0-9]+$"
      }
    },
    "out": {
      "data": {
        "message": "hi numbers"
      }
    }
  }
]
```
**Step 3:** Start the grpc mock server.

`./stubs` is the directory where all stub files reside.
```Bash
grpc-mock -mock-addr :22222 -import-path . -proto helloworld.proto -stub-dir ./stubs
```
The flag `mock-addr` defines the address mock server listens on.

**Step 4:** Ready to test it.
We use `grpcurl` to invoke rpc methods.
```Bash
grpcurl -plaintext -d '{"name": "01"}' localhost:22222 helloworld.Greeter/SayHello
```
The output is:
```JSON
{
  "message": "hi numbers"
}
```
More [examples](https://github.com/monlabs/grpc-mock/tree/main/examples)
## Stubs
The input and expected output comprise of a `Stub`. When a request is received, the mock server tries to match the stub with the request, and then reponds with the expected output in the matched stub once it's matched. There're two ways to manage your stubs: static and dynamic.
### Static: predefine stubs in files.
Like in the example above, you save the stubs in files. Let the mock server load them on starting.
### Dynamic: call API to CRUD stubs.
You also can start the mock server without specified stubs and call the admin API to dynamically add stubs. The API server runs on `:22220`. 

- `GET /v1/stubs`: get all stubs.
- `POST /v1/stubs`: create stubs.
- `DELETE /v1/stubs`: delete stubs.

### Match Rules
