package file

import (
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	x "github.com/javinc/mango/module"
)

const (
	uploadPath  = "./upload/"
	uploadField = "file"
)

// UploadHandler file
func UploadHandler(c *gin.Context) {
	os.Mkdir(uploadPath, 0777)

	file, header, err := c.Request.FormFile(uploadField)
	if err != nil {
		x.Error("FILE_UPLOAD_ERROR", err.Error())

		return
	}

	raw := strings.Split(header.Filename, ".")
	ext := raw[len(raw)-1]
	name := x.GenerateHash() + "." + ext
	filePath := uploadPath + name
	out, err := os.Create(filePath)
	if err != nil {
		x.Error("FILE_UPLOAD_CREATE_ERROR", err.Error())

		return
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		x.Error("FILE_UPLOAD_COPY_ERROR", err.Error())

		return
	}

	x.Output(gin.H{
		"slug": name,
		"ext":  ext,
		"mime": header.Header.Get("Content-Type"),
	})
}

// DisposeHandler file disposition
func DisposeHandler(c *gin.Context) {

}
