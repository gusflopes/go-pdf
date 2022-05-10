package main

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/ledongthuc/pdf"
)

type receipt struct {
	Account string `json:"string"`
	Content string `json:"string"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/files", func(ctx *gin.Context) {
		ctx.Request.ParseMultipartForm(10 * 1024 * 1024)
		account := ctx.PostForm("account")

		form, err := ctx.MultipartForm()
		if err != nil {
			// panic("Error reading request")
			ctx.JSON(400, gin.H{
				"message": "Bad request.",
				"error":   err.Error(),
			})
			return
		}

		files := form.File["files"]

		var Response []receipt

		for _, file := range files {
			fmt.Println("--------------------------")
			fmt.Println("Account: ", account)
			fmt.Println("File Name: ", file.Filename)
			fmt.Println("File Size: ", file.Size)
			fmt.Println("File Type: ", file.Header.Get("Content-Type"))
			fmt.Println("--------------------------")
			filename := filepath.Base(file.Filename)
			if err := ctx.SaveUploadedFile(file, filename); err != nil {
				ctx.String(http.StatusBadRequest, "upload file err: %s", err.Error())
				return
			}
			pdfData, err := readPdf(filename)
			if err != nil {
				panic("Failed.")
			}
			item := receipt{account, pdfData}
			fmt.Print(item)
			Response = append(Response, item)
			// fmt.Print(item)
		}

		fmt.Print(Response)

		ctx.JSON(200, &Response)
		// ctx.String(http.StatusOK, "Uploaded successfully %d files for account=%s.", len(files), account)
	})

	fmt.Println("Starting up the server...")
	r.Run()
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}
