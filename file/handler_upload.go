package file

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
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

	size, err := getFileSize(file)
	if err != nil {
		x.Error("FILE_UPLOAD_SIZE_ERROR", err.Error())

		return
	}

	x.Output(gin.H{
		"slug": name,
		"ext":  ext,
		"size": size,
		"mime": header.Header.Get("Content-Type"),
	})
}

func getFileSize(file multipart.File) (int64, error) {
	type size interface {
		Size() int64
	}

	var fsize string
	var i64 int64

	if sizeInterface, ok := file.(size); ok {
		sizeInterface.Size()

		fsize = fmt.Sprintf("%d", sizeInterface.Size())
	} else {
		return i64, errors.New("cant get file size")
	}

	s, err := strconv.Atoi(fsize)
	if err != nil {
		return i64, err
	}

	return int64(s), nil
}

// DisposeHandler file disposition
func DisposeHandler(c *gin.Context) {
}
