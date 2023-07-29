package msg

import (
	"context"
	"encoding/json"
	"project/consts"
	"project/repository/db/dao"
	"project/repository/db/model"
	"project/repository/kafka"
	"project/types"
	"project/utils/snowflake"
	"strconv"

	"github.com/mlogclub/simple/common/jsons"
	"go.uber.org/zap"
)

func SendPostLikeMsgToKafka(ctx context.Context, postId, likeUserId int64) {
	post, err := dao.NewPostModel().GetPostDetailById(ctx, strconv.FormatInt(postId, 10))
	if err != nil {
		zap.L().Error("SendPostLikeMsgToKafka get post user id failed", zap.Error(err))
		return
	}

	ed := &types.PostLikeExtraData{
		PostId:     postId,
		LikeUserId: likeUserId,
	}

	m := &model.Message{
		MessageId:    snowflake.GenID(),
		FromId:       likeUserId,
		UserId:       post.AuthorId,
		Type:         int(consts.TypePostLike),
		Status:       consts.StatusUnread,
		Title:        consts.MessageLikePost,
		Content:      "",
		QuoteContent: "《" + post.Title + "》",
		ExtraData:    jsons.ToJsonStr(ed),
	}

	msgStr, _ := json.Marshal(m)

	err = kafka.SendMessage(consts.KafkaKey, consts.TopicLike, string(msgStr))
	if err != nil {
		zap.L().Error("SendMessage PostLikeMsgToKafka failed", zap.Error(err))
	}
}
