package render

import (
	"project/types"

	"github.com/mlogclub/simple/common/jsons"
	"github.com/mlogclub/simple/common/strs"
	"go.uber.org/zap"
)

func buildImage(imageStr string) *types.ImageInfo {
	if strs.IsBlank(imageStr) {
		return nil
	}
	var img *types.ImageDTO
	if err := jsons.Parse(imageStr, &img); err != nil {
		zap.L().Error("jsons.Parse(imageStr, &img); err", zap.Error(err))
		return nil
	}

	return &types.ImageInfo{
		Url:     HandleOssImageStyleDetail(img.Url),
		Preview: HandleOssImageStylePreview(img.Url),
	}
}
