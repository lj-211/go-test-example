package mock

//go:generate mockgen -source=../method-and-interface/interface.go  -destination=./mock_interface.go -package=mock
//go:generate mockgen -source=../nettest/proto/example.pb.go  -destination=./mock_example.pb.go -package=mock
