package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	// Cria uma nova sessão com as credenciais e configurações da AWS
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("sa-east-1"),
		Credentials: credentials.NewStaticCredentials(
			"KEY AQUI",
			"SECRET AQUI",
			"",
		),
	})

	if err != nil {
		panic(err)
	}

	// Cria um cliente S3 usando a sessão criada acima
	s3Client = s3.New(sess)
	s3Bucket = "go-expert-pos"
}

func main() {
	// Abre o diretório "./tmp"
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	uploadControl := make(chan struct{}, 100)

	errorFileUpload := make(chan string, 10)

	// Go routine para monitorar os erros de upload e tentar novamente
	go func() {
		for {
			select {
			case filename := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(filename, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		// Lê os arquivos do diretório, um por vez
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}
		// Chama a função uploadFile para fazer o upload do arquivo
		wg.Add(1)

		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()

	completeFileName := fmt.Sprintf("./tmp/%s", filename)

	fmt.Printf("Uploading file %s to Bucket %s\n", completeFileName, s3Bucket)

	// Abre o arquivo para leitura
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", completeFileName, err)

		errorFileUpload <- completeFileName // Adiciona o arquivo ao canal de erro

		<-uploadControl // Esvazia o canal

		return
	}

	defer f.Close()

	// Faz o upload do arquivo para o bucket S3
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})

	if err != nil {
		fmt.Printf("Error uploading file %s: %s\n", completeFileName, err)

		errorFileUpload <- completeFileName // Adiciona o arquivo ao canal de erro

		<-uploadControl // Esvazia o canal

		return
	}

	fmt.Printf("File %s uploaded successfully\n", completeFileName)
	<-uploadControl // Esvazia o canal
}
