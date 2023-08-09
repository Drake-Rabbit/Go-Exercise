package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const UPLOAD_DIR = "uplaod"

func Upload(c *gin.Context) {
	file, err := c.FormFile("img")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"err":  err.Error(),
		})
	}
	arr := strings.Split(file.Filename, ".")

	filePath, ok := checkExtAndgetFilePath(arr[1])
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"err":  "invaild ext",
		})
		return
	}
	fmt.Println(filePath)
	if err = c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"err":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"msg":     "upload success",
		"img_url": filePath,
	})
}
func checkExtAndgetFilePath(fileExt string) (string, bool) {
	allowExt := map[string]bool{
		"jpg":  true,
		"png":  true,
		"jpeg": true,
	}
	if _, ok := allowExt[fileExt]; !ok {
		return " ", false
	}

	if _, ok := os.Stat(UPLOAD_DIR); ok != nil {
		os.Mkdir(UPLOAD_DIR, os.ModePerm)
	}

	filepath := UPLOAD_DIR + "/" + strconv.FormatInt(time.Now().Unix(), 10) + "." + fileExt
	return filepath, true
}
