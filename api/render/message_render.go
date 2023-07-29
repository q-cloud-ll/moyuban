package render

import (
	"project/consts"
	"project/repository/db/model"
	"project/types"
	"project/utils/myburls"

	"github.com/tidwall/gjson"
)

func BuildMessage(from *model.User, msg *model.Message) *types.MessageResp {
	if msg == nil {
		return nil
	}

	detailUrl := getMessageDetailUrl(msg)

	resp := &types.MessageResp{
		MessageId:    msg.MessageId,
		From:         from,
		UserId:       msg.UserId,
		Title:        msg.Title,
		Content:      msg.Content,
		QuoteContent: msg.QuoteContent,
		Type:         msg.Type,
		DetailUrl:    detailUrl,
		ExtraData:    msg.ExtraData,
		Status:       msg.Status,
		CreateTime:   msg.CreatedAt,
	}

	return resp
}

// BuildMessages 渲染消息列表
func BuildMessages(user *model.User, messages []*model.Message) []types.MessageResp {
	if len(messages) == 0 {
		return nil
	}
	var responses []types.MessageResp
	for _, message := range messages {
		responses = append(responses, *BuildMessage(user, message))
	}
	return responses
}

func getMessageDetailUrl(t *model.Message) string {
	msgType := consts.Type(t.Type)
	if msgType == consts.TypePostComment {
		entityType := gjson.Get(t.ExtraData, "entityType")
		entityId := gjson.Get(t.ExtraData, "entityId")
		if entityType.String() == consts.EntityPost {
			return myburls.PostUrl(entityId.Int())
		} else if entityType.String() == consts.EntityTopic {
			return myburls.TopicUrl(entityId.Int())
		}
	} else if msgType == consts.TypePostCommentReply {
		entityType := gjson.Get(t.ExtraData, "rootEntityType")
		entityId := gjson.Get(t.ExtraData, "rootEntityId")
		if entityType.String() == consts.EntityPost {
			return myburls.PostUrl(entityId.Int())
		} else if entityType.String() == consts.EntityTopic {
			return myburls.TopicUrl(entityId.Int())
		}
	} else if msgType == consts.TypePostLike || msgType == consts.TypePostRecommend {
		postId := gjson.Get(t.ExtraData, "postId")
		if postId.Exists() && postId.Int() > 0 {
			return myburls.PostUrl(postId.Int())
		}
	}

	return myburls.AbsUrl("/user/messages")
}
