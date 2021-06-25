package main;

import (
	"fmt"
	"context"
	"mygrpc/src/service"
)

type myServer struct {

}

func (s *myServer) sayName(ctx context.Context, point *pb.point) {
	fmt.Println("In sayName")
}
func main()  {
	fmt.Println("Hii")
}