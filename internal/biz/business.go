package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// ReplyParam is a Reply model.
type ReplyParam struct {
	ReviewID  int64
	StoreID   int64
	Content   string
	PicInfo   string
	VideoInfo string
}

// AppealReviewParam is a AppealReview model.
type AppealReviewParam struct {
	ReviewID  int64
	StoreID   int64
	Reason    string
	Content   string
	PicInfo   string
	VideoInfo string
}

// BusinessRepo is a Business repo.
type BusinessRepo interface {
	Reply(context.Context, *ReplyParam) (int64, error)
	Appeal(context.Context, *AppealReviewParam) (int64, error)
}

// BusinessUsecase is a Business usecase.
type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewBusinessUsecase(repo BusinessRepo) *BusinessUsecase {
	return &BusinessUsecase{repo: repo}
}

// CreateReply 创建回复
func (uc *BusinessUsecase) CreateReply(ctx context.Context, param *ReplyParam) (int64, error) {
	uc.log.WithContext(ctx).Infof("[biz] CreateReply param:%v", param)
	return uc.repo.Reply(ctx, param)
}

// AppealReview 申诉评论
func (uc *BusinessUsecase) AppealReview(ctx context.Context, param *AppealReviewParam) (int64, error) {
	uc.log.WithContext(ctx).Infof("[biz] AppealReview param:%v", param)
	return uc.repo.Appeal(ctx, param)
}
