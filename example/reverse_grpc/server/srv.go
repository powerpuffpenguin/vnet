package server

import (
	"context"
	"io"
	"log"
	"math/rand"
	grpc_math "reverse_grpc/math"
)

// implemented grpc
type Server struct {
	grpc_math.UnimplementedMathServer
}

func (s Server) Add(ctx context.Context, req *grpc_math.AddRequest) (resp *grpc_math.AddResponse, e error) {
	var sum int64
	for _, val := range req.Vals {
		sum += val
	}
	log.Println(`add`, req.Vals, `=`, sum)
	resp = &grpc_math.AddResponse{
		Sum: sum,
	}
	return
}

func (s Server) Random(stream grpc_math.Math_RandomServer) (e error) {
	var (
		req *grpc_math.RandomRequest
	)
	for {
		req, e = stream.Recv()
		if e != nil {
			if e == io.EOF {
				e = nil
				return
			}
			break
		}
		val := rand.Int63n(req.Max + 1)
		log.Println(`rand`, req.Max, `=`, val)
		e = stream.Send(&grpc_math.RandomResponse{
			Val: val,
		})
		if e != nil {
			break
		}
	}
	log.Println(`Random`, e)
	return
}
