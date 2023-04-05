package global

import (
	"app_ws/config"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-redis/redis"
	"github.com/minio/minio-go/v6"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG    *zap.Logger
	TRANS      ut.Translator
	MINIO      *minio.Client
)
