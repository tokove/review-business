package data

import (
	"context"
	v1 "review-business/api/review/v1"
	"review-business/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type businessRepo struct {
	data *Data
	log  *log.Helper
}

// NewBusinessRepo .
func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &businessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Reply 创建评论
func (r *businessRepo) Reply(ctx context.Context, param *biz.ReplyParam) (int64, error) {
	r.log.WithContext(ctx).Infof("[data] Reply, param:%v", param)
	ret, err := r.data.rc.ReplyReview(ctx, &v1.ReplyReviewRequest{
		ReviewId:  param.ReviewID,
		StoreId:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	r.log.WithContext(ctx).Debugf("ReplyReview return, ret:%v, err:%v", ret, err)
	if err != nil {
		return 0, err
	}
	return ret.GetReplyId(), nil
}

// Appeal 申诉评论
func (r *businessRepo) Appeal(ctx context.Context, param *biz.AppealReviewParam) (int64, error) {
	r.log.WithContext(ctx).Infof("[data] Appeal, param:%v", param)
	ret, err := r.data.rc.AppealReview(ctx, &v1.AppealReviewRequest{
		ReviewId:  param.ReviewID,
		StoreId:   param.StoreID,
		Reason:    param.Reason,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	r.log.WithContext(ctx).Debugf("AppealReview return, ret:%v, err:%v", ret, err)
	if err != nil {
		return 0, err
	}
	return ret.GetAppealId(), nil
}