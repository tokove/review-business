package biz

import "context"

// ReplyParam is a Reply model.
type ReplyParam struct {
	ReviewID  int64
	StoreID   int64
	Content   string
	PicInfo   string
	VideoInfo string
}

// BusinessRepo is a Business repo.
type BusinessRepo interface {
	Reply(context.Context, *ReplyParam) (int64, error)
}

// BusinessUsecase is a Business usecase.
type BusinessUsecase struct {
	repo BusinessRepo
}

// NewGreeterUsecase new a Greeter usecase.
func NewBusinessUsecase(repo BusinessRepo) *BusinessUsecase {
	return &BusinessUsecase{repo: repo}
}

// Reply 创建回复
func (uc *BusinessUsecase) CreateReply(ctx context.Context, param *ReplyParam) (int64, error) {

	return uc.repo.Reply(ctx, param)
}
