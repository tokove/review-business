package service

import (
	"context"
	"fmt"

	pb "review-business/api/business/v1"
	"review-business/internal/biz"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer

	uc *biz.BusinessUsecase
}

func NewBusinessService(uc *biz.BusinessUsecase) *BusinessService {
	return &BusinessService{uc: uc}
}

// ReplyReview 回复评论
func (s *BusinessService) ReplyReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	replyID, err := s.uc.CreateReply(ctx, &biz.ReplyParam{
		ReviewID:  req.GetReviewId(),
		StoreID:   req.GetStoreId(),
		Content:   req.GetContent(),
		PicInfo:   req.GetPicInfo(),
		VideoInfo: req.GetVideoInfo(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.ReplyReviewReply{ReplyId: replyID}, nil
}

// AppealReview 申诉评论
func (s *BusinessService) AppealReview(ctx context.Context, req *pb.AppealReviewRequest) (*pb.AppealReviewReply, error) {
	fmt.Printf("[service] AppealReview, req:%v", req)
	appealID, err := s.uc.AppealReview(ctx, &biz.AppealReviewParam{
		ReviewID: req.GetReviewId(),
		StoreID:  req.GetStoreId(),
		Reason:   req.GetReason(),
		Content: req.GetContent(),
		PicInfo: req.GetPicInfo(),
		VideoInfo: req.GetVideoInfo(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.AppealReviewReply{AppealId: appealID}, nil
}