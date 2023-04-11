package object

import (
	"WowjoyProject/WADO_URL/global"
)

//var token string

// 封装对象相关操作
type Object struct {
	InstanceKey int64
	ResId       string
	Key         string
	Tags        map[string]string
	Path        string
	Count       int
}

func NewObject(data global.ObjectData) *Object {
	return &Object{
		InstanceKey: data.InstanceKey,
		// ResId:       global.ObjectSetting.OBJECT_ResId,
		Key:   data.Key,
		Path:  data.Path,
		Count: data.Count,
	}
}

// // 上传对象[POST]
// func (obj *Object) UploadObject() {
// 	global.Logger.Info("开始上传对象：", *obj)
// 	// 判断文件大小，来区别是否开始分段上传
// 	// var code string
// 	// fileSize := general.GetFileSize(obj.Path)
// 	// if fileSize >= (int64(global.ObjectSetting.File_Fragment_Size << 20)) {
// 	// 	code = UploadLargeFile(obj, fileSize)
// 	// } else {
// 	code := UploadSmallFile(obj)
// 	if code == "00000" {
// 		//上传成功更新数据库
// 		global.Logger.Info("数据上传成功", obj.InstanceKey)
// 		model.UpdateUplaode(obj.InstanceKey, obj.Key, true)
// 	} else {
// 		global.Logger.Info("数据上传失败", obj.InstanceKey)
// 		// 上传失败时先补偿操作，补偿操作失败后才更新数据库
// 		if !ReDo(obj, global.UPLOAD) {
// 			global.Logger.Info("数据补偿失败", obj.InstanceKey)
// 			// 上传失败更新数据库
// 			model.UpdateUplaode(obj.InstanceKey, obj.Key, false)
// 		}
// 	}
// }

// // 下载对象[GET]
// func (obj *Object) DownObject() {
// 	// if token == "" {
// 	// 	// 获取token
// 	// 	token = "Bearer " + GetToken()
// 	// }
// 	// 请求处理太快，http资源没来得及关闭
// 	// time.Sleep(50 * time.Millisecond)
// 	global.Logger.Info("开始下载对象：", *obj)
// 	flag := DownFile(obj)
// 	if flag {
// 		global.Logger.Info("下载成功：" + obj.Path)
// 		// model.UpdateDown(obj.InstanceKey, obj.Key, true)
// 	} else {
// 		// 下载失败时先补偿操作，补偿操作失败后才更新数据库
// 		if !ReDo(obj, global.DOWNLOAD) {
// 			global.Logger.Info("数据补偿失败", obj.InstanceKey)
// 			// 下载失败更新数据库
// 			// model.UpdateDown(obj.InstanceKey, obj.Key, false)
// 		}
// 	}
// }

// // // UploadLargeFile 上传大文件
// // func UploadLargeFile(obj *Object, size int64) string {
// // 	global.Logger.Debug("开始执行大文件上传", obj.Key)
// // 	// num := math.Ceil(float64(size) / float64(global.ObjectSetting.Each_Section_Size))
// // 	// 1.初始化
// // 	UploadId := Multipart_Upload_Init(obj)
// // 	if UploadId == "" {
// // 		global.Logger.Error("分段上传初始化获取UploadId是空,结束任务")
// // 		return ""
// // 	}
// // 	global.Logger.Info("UploadId: ", UploadId)
// // 	// 2.开始上传小段对象
// // 	if Multipart_Upload(obj, UploadId) {
// // 		// 文件上传成功完结操作
// // 		Multipart_Completion(obj, UploadId)
// // 	} else {
// // 		// 文件上传失败取消操作
// // 		Multipart_Abortion(obj, UploadId)
// // 	}
// // 	return ""
// // }
