package main

import (
	"bytes"
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ledongthuc/pdf"
)

type receipt struct {
	Account     string      `json:"account"`
	Transaction transaction `json:"transaction"`
	Content     string      `json:"raw_content"`
}

type transaction struct {
	Date   string `json:"date"`
	Payee  string `json:"payee"`
	TaxId  string `json:"tax_id"`
	Amount string `json:"amount"`
	PixKey string `json:"pix_key"`
	Memo   string `json:"memo"`
}

func main() {
	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Server up and running.",
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

		Response := make([]receipt, 0)

		for _, file := range files {
			// println("file.Filename: ", file.Filename)                        // Original name if needed
			newFileName := uuid.New().String() + filepath.Ext(file.Filename) // New Name
			dst := path.Join("./tmp/", newFileName)                          // Path after saving

			if err := ctx.SaveUploadedFile(file, dst); err != nil {
				ctx.String(http.StatusBadRequest, "upload file err: %s", err.Error())
				return
			}

			pdfData, err := readPdf(dst)
			if err != nil {
				panic("Failed.")
			}

			payee := stringParser(pdfData, "Conta de crédito:", " | CNPJ:")
			date := regexParser(pdfData, `\d{2}\/\d{2}\/\d{4}(?!.*\d{2}\/\d{2}\/\d{4})`)
			cnpj := regexParser(pdfData, `\d{2}.\d{3}.\d{3}\/\d{4}-\d{2}`)
			amount := regexParser(pdfData, `(?<=Tarifa:R\$ )(.*)(?=Valor)`)

			transaction := transaction{Date: date, Amount: amount, Payee: payee, TaxId: cnpj}
			fmt.Println("Transaction: ", transaction)

			// bytes, err := json.Marshal(transaction)
			// if err != nil {
			// 	panic(err)
			// }
			// fmt.Println(string(bytes))

			item := receipt{Account: account, Transaction: transaction, Content: pdfData}

			Response = append(Response, item)
		}

		ctx.JSON(200, Response)
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

func regexParser(payload string, regex string) string {
	rxp := regexp.MustCompile(regex, regexp.None)

	output, err := rxp.FindStringMatch(payload)
	if err != nil {
		panic("Erro no regexp")
	}
	// print(output.String())
	return output.String()
}

// This function is called to extract specific information from a parsed PDF file
func stringParser(payload string, start string, end string) string {
	firstIndex := strings.Index(payload, start)
	secondIndex := strings.Index(payload, end)
	// println("Start: ", firstIndex, " - End: ", secondIndex)

	// Problem!!
	// secondIndex must be bigger then first Index - Maybe I'll need to split the string
	// to make sure i'll get the next string
	// For now I'm solving with Regex - but it'd be good to fix.

	extracted := payload[firstIndex+len(start) : secondIndex]
	// println("Extracted Data: ", extracted)

	// split := strings.Split(payload, start)
	// print(split)

	return extracted
}
