package reserve

import (
	"context"
	"errors"
	"wallet/internal/core"
	pb "wallet/pkg/grpc"
)

type service struct {
	pb.WalletReservationServiceServer

	walletService core.WalletService
}

func NewService(
	walletService core.WalletService,
) pb.WalletReservationServiceServer {
	return &service{
		walletService: walletService,
	}
}

func (*service) GetBalance(context.Context, *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	return nil, errors.New("Not Implemented")

}

func (*service) CreateReservation(context.Context, *pb.CreateReservationRequest) (*pb.CreateReservationResponse, error) {
	return nil, errors.New("Not Implemented")

}

func (*service) ConfirmReservation(context.Context, *pb.ConfirmReservationRequest) (*pb.ConfirmReservationResponse, error) {
	return nil, errors.New("Not Implemented")

}

func (*service) PartialConfirmReservation(context.Context, *pb.PartialConfirmReservationRequest) (*pb.PartialConfirmReservationResponse, error) {
	return nil, errors.New("Not Implemented")

}

func (*service) CancelReservation(context.Context, *pb.CancelReservationRequest) (*pb.CancelReservationResponse, error) {
	return nil, errors.New("Not Implemented")

}

func (*service) ExtendReservation(context.Context, *pb.ExtendReservationRequest) (*pb.ExtendReservationResponse, error) {
	return nil, errors.New("Not Implemented")

}

func (*service) GetReservation(context.Context, *pb.GetReservationRequest) (*pb.GetReservationResponse, error) {
	return nil, errors.New("Not Implemented")

}

func (*service) ListReservations(context.Context, *pb.ListReservationsRequest) (*pb.ListReservationsResponse, error) {
	return nil, errors.New("Not Implemented")

}
