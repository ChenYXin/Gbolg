package models

import (
	"GBolg/conf/errmsg"
	"GBolg/utils/viper_config"
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

func UploadToQiniu(file multipart.File, fileSize int64) (code int, url string) {
	accessKey := viper_config.VC.GetString("qiniu.accessKey")
	secretKey := viper_config.VC.GetString("qiniu.secretKey")
	bucket := viper_config.VC.GetString("qiniu.bucket")
	imgUrl := viper_config.VC.GetString("qiniu.qiniuServer")

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		// 空间对应的机房
		Zone: &storage.ZoneHuanan,
		// 是否使用https域名
		UseHTTPS: false,
		// 上传是否使用CDN上传加速
		UseCdnDomains: false,
	}

	// 可选配置
	putExtra := storage.PutExtra{
		//Params: map[string]string{
		//	"x:name": "qiniuyun logo",
		//},
	}

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return errmsg.EerrorQiniuUploadFail, ""
	}
	url = imgUrl + ret.Key
	fmt.Println(ret.Key, ret.Hash)
	return errmsg.SUCCESS, url
}
