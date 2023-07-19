package timer

import (
	"context"
	"go.uber.org/zap"
	"project/repository/cache"
	"project/repository/db/dao"
)

var t = NewTimerTask()

func TimeTask() {
	go func() {
		_, err := t.AddTaskByFunc("updateStarDetailFromRedisToMySQL", "@every 20m", func() {
			err := cache.NewStarCache().UpdateStarDetailFromRedisToMySQL(context.Background(), dao.NewDBClient())
			if err != nil {
				zap.L().Error("updateStarDetailFromRedisToMySQL timerTask error", zap.Error(err))
			}
		})
		if err != nil {
			zap.L().Error("add timerTask error", zap.Error(err))
		}
	}()
}
