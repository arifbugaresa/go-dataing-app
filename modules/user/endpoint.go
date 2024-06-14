package user

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	LoginEndpoint endpoint.Endpoint
}

func MakeLoginEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UserReq)
		output, err := s.Login(req)
		if err != nil {
			return output, nil
		}
		return output, nil
	}
}
