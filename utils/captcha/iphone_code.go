package captcha

import (
	"context"
	"encoding/json"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"go.uber.org/zap"
	"project/consts"
	"project/repository/cache"
	"project/setting"
)

// SendMessage 短信发送
func SendMessage(ctx context.Context, phone string, code string) error {
	//参数一：连接的节点地址（有很多节点选择，这里我选择杭州）
	//参数二：AccessKey ID
	//参数三：AccessKey Secret
	pConfig := setting.Conf.IPhoneMsgConfig
	client, err := dysmsapi.NewClientWithAccessKey(pConfig.RegionId, pConfig.AccessKeyId, pConfig.AccessKeySecret)

	request := dysmsapi.CreateSendSmsRequest()       //创建请求
	request.Scheme = pConfig.Scheme                  //请求协议
	request.PhoneNumbers = phone                     //接收短信的手机号码
	request.SignName = pConfig.SignName              //短信签名名称
	request.TemplateCode = pConfig.TemplateCode      //短信模板ID
	par, err := json.Marshal(map[string]interface{}{ //定义短信模板参数（具体需要几个参数根据自己短信模板格式）
		"code": code,
	})
	request.TemplateParam = string(par) //将短信模板参数传入短信模板

	response, err := client.SendSms(request) //调用阿里云API发送信息

	if err != nil { //处理错误
		zap.L().Error("sendsms failed", zap.Error(err))
		return consts.UserSendMsgErr
	}
	if err := cache.NewUserCache().SavePhoneMsg(ctx, phone, code); err != nil {
		zap.L().Error("cache.SavePhoneMsg(ctx, phone, code)", zap.Error(err))
		return err
	}

	zap.L().Info("sendsms response", zap.Any("response ", response)) //控制台输出响应

	return nil
}
