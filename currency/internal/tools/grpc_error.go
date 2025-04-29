package tools

import (
	"currencyService/currency/internal/repository"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GrpcError(err error) error {
	switch {
	case errors.Is(err, repository.ErrNotFound):
		return status.Errorf(codes.NotFound, "not found: %v", err)
	default:
		return status.Errorf(codes.Internal, "internal server error: %v", err)
	}
}
