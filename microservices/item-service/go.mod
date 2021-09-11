module github.com/dty1er/kubecept/microservices/item-service

go 1.16

require (
	github.com/dty1er/kubecept/microservices/proto/pb v0.0.1
	google.golang.org/grpc v1.40.0
)

replace github.com/dty1er/kubecept/microservices/proto/pb => ../proto/pb
