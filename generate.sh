protoc sayname/saynamepb/sayname.proto \
    --go_out=. --go-grpc_out=.
protoc calculator/calculatorpb/calculator.proto \
    --go_out=. --go-grpc_out=.
protoc greet/greetpb/greet.proto \
    --go_out=. --go-grpc_out=.
