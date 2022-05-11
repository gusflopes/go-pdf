package files

import (
	"fmt"
	"net/http"
)

func uploadFiles(request http.Request, response http.ResponseWriter) {
	request.ParseMultipartForm(10 * 1024 * 1024)

	files := request.MultipartForm.File["myFiles"]

	for _, file := range files {
		fmt.Println("File Name: ", file.Filename)
		fmt.Println("File Size: ", file.Size)
		fmt.Println("File Type: ", file.Header.Get("Content-Type"))
		fmt.Println("---------------------------------------------")
	}

}
