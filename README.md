# Kaf-Protobuf-Struct-Analyser

$ cd Kaf-Protobuf-Struct-Analyser/protos
$ protoc --go_out=../pkg/pbprotos/ *.proto

// Start Kafka

// Run the Kafka Consumer

$ cd Kaf-Protobuf-Struct-Analyser/pkg/test-read
$ go run test_read.go


// Run the Kafka Producer

$ cd Kaf-Protobuf-Struct-Analyser/pkg/test-write
$ go run test_write.go
