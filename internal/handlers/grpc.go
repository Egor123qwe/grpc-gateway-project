package handlers

import (
	"context"

	"Cortlex_test/internal/models"
	quotepb "Cortlex_test/proto/api"
)

type Server struct {
	QS QuoteService
	quotepb.UnimplementedQuoteServiceServer
}

type QuoteService interface {
	RandomQuote(ctx context.Context) (models.Quote, error)
	LikeQuote(ctx context.Context, id string) (int64, error)
	DislikeQuote(ctx context.Context, id string) (int64, error)
}

func (s Server) GetRandomQuote(ctx context.Context, in *quotepb.Empty) (*quotepb.Quote, error) {
	quote, err := s.QS.RandomQuote(ctx)
	if err != nil {
		return nil, err
	}
	return &quotepb.Quote{
		Id:      quote.ID.Hex(),
		Content: quote.Content,
		Rating:  quote.Rating,
	}, nil
}

func (s Server) LikeQuote(ctx context.Context, in *quotepb.RatingRequest) (*quotepb.RatingResponse, error) {
	rating, err := s.QS.LikeQuote(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return &quotepb.RatingResponse{
		Rating:  rating,
		Message: "successfully liked",
	}, nil
}

func (s Server) DislikeQuote(ctx context.Context, in *quotepb.RatingRequest) (*quotepb.RatingResponse, error) {
	rating, err := s.QS.DislikeQuote(ctx, in.GetId())
	if err != nil {
		return nil, err
	}
	return &quotepb.RatingResponse{
		Rating:  rating,
		Message: "successfully disliked",
	}, nil
}
