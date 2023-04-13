package initialize

import (
	"fmt"
	"github.com/minio/minio-go/v6"
	"im_app/global"
)

func InitMinIO() (minioClient *minio.Client) {
	fmt.Println(global.GVA_CONFIG.Minio)
	// 初使化 minio client对象。false是关闭https证书校验
	minioClient, err := minio.New(global.GVA_CONFIG.Minio.Endpoint, global.GVA_CONFIG.Minio.AccessKeyID, global.GVA_CONFIG.Minio.SecretAccessKey, global.GVA_CONFIG.Minio.Secure)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return minioClient
}
