module github.com/ShintaNakama/twitter-clone/bff

go 1.15

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/google/go-cmp v0.5.5
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/vektah/gqlparser/v2 v2.1.0
	go.uber.org/zap v1.16.0
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	google.golang.org/grpc v1.37.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.31.0
	pb v1.0.0
)

replace pb v1.0.0 => ../pb
