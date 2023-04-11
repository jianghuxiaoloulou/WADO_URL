package v1

import (
	"WowjoyProject/WADO_URL/global"
	"WowjoyProject/WADO_URL/internal/model"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type WadoURLData struct {
	RequestType     string `from:"requestType" binding:"required"`
	StudyUID        string `from:"studyUID" binding:"required"`
	SeriesUID       string `from:"seriesUID" binding:"required"`
	ObjectUID       string `from:"objectUID" binding:"required"`
	RetrieveAETitle string `from:"RetrieveAETitle" binding:"required"`
}

// @Summary WADOURL服务
// @Description
// @Tags file
// @Param requestType query string true "请求的类型"
// @Param studyUID query string true "检查"
// @Param seriesUID query string true "序列"
// @Param objectUID query string true "对象（文件）UID，可以叫做SOP"
// @Param RetrieveAETitle query string true "医院AETitle"
// @Success 200 {object} gin.Context
// @Router /wado [GET] test.dcm
func WadoURL(c *gin.Context) {
	global.Logger.Debug(c.Request.URL)
	objectUID := c.Query("objectUID")
	RetrieveAETitle := c.Query("RetrieveAETitle")
	global.Logger.Debug("请求对象", objectUID)

	if RetrieveAETitle != global.GeneralSetting.RetrieveAETitle {
		global.Logger.Debug("不是本医院的请求", RetrieveAETitle)
		c.JSON(http.StatusOK, gin.H{
			"code": "1",
			"msg":  "不是本医院的请求",
			"data": RetrieveAETitle,
		})
		return
	}
	// if data.requestType != global.GeneralSetting.RequestType {
	// 	global.Logger.Debug("请求类型错误", data.requestType)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": "1",
	// 		"msg":  "请求类型错误",
	// 		"data": data.requestType,
	// 	})
	// 	return
	// }
	// 获取文件路径
	// filePath := "D:\\Go_Work\\src\\WowjoyProject\\WADO_URL\\storage\\file\\CT.0a45aefd69630746402730c57a4eeeb4.dcm"
	filePath := model.GetImagePath(objectUID)
	if filePath == "" {
		global.Logger.Debug("未找到文件", filePath)
		c.JSON(http.StatusOK, gin.H{
			"code": "1",
			"msg":  "未找到文件",
			"data": "",
		})
		return
	}
	global.Logger.Debug("获取的文件：", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		global.Logger.Error(err)
		c.JSON(http.StatusFound, gin.H{
			"code": "1",
			"msg":  err,
			"data": "",
		})
		return
	}

	defer file.Close()

	fileHeader := make([]byte, 1204)
	file.Read(fileHeader)
	fileStat, _ := file.Stat()

	// c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Type", "application/dicom")
	c.Header("Content-Disposition", "attachment; filename="+objectUID+".dcm")
	c.Header("Content-Length", strconv.FormatInt(fileStat.Size(), 10))

	file.Seek(0, 0)
	// c.File(filePath)
	// c.Writer.(file)
	io.Copy(c.Writer, file)
	global.Logger.Info("请求成功")
}

func FileDownload(w http.ResponseWriter, r *http.Request) {
	objectUID := ""
	RetrieveAETitle := ""
	r.ParseForm()
	for k, v := range r.Form {
		if k == "objectUID" {
			objectUID = strings.Join(v, "")
		}
		if k == "RetrieveAETitle" {
			RetrieveAETitle = strings.Join(v, "")
		}
	}
	if RetrieveAETitle != global.GeneralSetting.RetrieveAETitle {
		global.Logger.Debug("不是本医院的请求", RetrieveAETitle)
		fmt.Fprintf(w, "请求的TETitle错误", RetrieveAETitle)
		return
	}
	if objectUID == "" {
		fmt.Fprintf(w, "请求对象objectUID为空")
		return
	}
	global.Logger.Debug("请求对象", objectUID)
	filePath := model.GetImagePath(objectUID)
	if filePath == "" {
		global.Logger.Debug("未找到文件", filePath)
		fmt.Fprintf(w, "没有找到对应的文件")
		return
	}
	global.Logger.Debug("获取的文件：", filePath)

	// filepath := "\\\\10.42.1.14\\image1\\AN_CTAWP75796\\2017\\08\\24\\CT\\1.3.12.2.1107.5.1.4.75796.30000017082406375314500000001\\00031273.dcm"
	file, err := os.Open(filePath)
	if err != nil {
		global.Logger.Error(err)
		fmt.Fprintf(w, "文件打开错误")
		return
	}

	defer file.Close()

	fileHeader := make([]byte, 4096)
	file.Read(fileHeader)
	// fileStat, _ := file.Stat()
	w.Header().Set("Content-Type", "application/dicom")
	w.Header().Set("Content-Disposition", "attachment; filename="+objectUID+".dcm")
	// w.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10))

	file.Seek(0, 0)
	io.Copy(w, file)
	return
}
