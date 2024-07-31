package grpcproduct

import (
	sandjmav1 "github.com/Sanchir01/protos_files_job/pkg/gen/golang/auth"
	"google.golang.org/grpc"
)

type Products interface {
}

type serverProductsApi struct {
	sandjmav1.UnimplementedProductServer
	products Products
}

func NewProductsApi(gRPC *grpc.Server, products Products) *Products {
	return &Products{&serverProductsApi{gRPC, products: products}}
}
