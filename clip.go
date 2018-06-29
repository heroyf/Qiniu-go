package main

import (
	"fmt"
	"strings"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

func clip() {
	bucket := "your-bucket"
	key := "your-key"
	accessKey := "Your-accessKey"
	secretKey := "Your-secrectKey"
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		UseHTTPS: false,
	}
	saveBucket := "purpose-bucket"

	// 处理指令集合
	fopAvthumb := fmt.Sprintf("avthumb/mp4/ss/60/t/60|saveas/%s",
		storage.EncodedEntry(saveBucket, "pfop_test_qiniu.mp4"))

	fopBatch := []string{fopAvthumb}
	fops := strings.Join(fopBatch, ";")

	// 强制重新执行数据处理任务
	force := true
	// 数据处理指令全部完成之后，通知该地址
	notifyURL := "http://api.example.com/pfop/callback"
	// 数据处理的私有队列，必须指定以保障处理速度
	pipeline := "yangfan"
	operationManager := storage.NewOperationManager(mac, &cfg)
	persistentID, err := operationManager.Pfop(bucket, key, fops, pipeline, notifyURL, force)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(persistentID)

}
