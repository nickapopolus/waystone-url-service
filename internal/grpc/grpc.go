package grpcCon

import (
	"context"
	urlgrpcv1 "github.com/nickapopolus/waystone-url-service/proto/urlservice/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type URLServiceGRPC struct {
	urlgrpcv1.UnimplementedURLServiceServer
}

func NewURLServiceGRPC() *URLServiceGRPC {
	return &URLServiceGRPC{}
}

func (grpc *URLServiceGRPC) CreateURL(ctx context.Context, newURL *urlgrpcv1.CreateURLRequest) (*urlgrpcv1.CompleteURLResponse, error) {
	return &urlgrpcv1.CompleteURLResponse{
		Id:             "This is the id of the URL",
		UserId:         "This is the id of the user",
		OriginalUrl:    newURL.OriginalUrl,
		Slug:           "skjdfhs",
		CustomSlug:     newURL.CustomSlug,
		ShortUrl:       "https://blo.wme/skjdhs",
		ClickCount:     1,
		MaxClicks:      newURL.MaxClicks,
		IsActive:       newURL.IsActive,
		ExpirationDate: timestamppb.New(time.Now().UTC().AddDate(1, 1, 0)),
		CreationAt:     timestamppb.New(time.Now().UTC()),
	}, nil
}
