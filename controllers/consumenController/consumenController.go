package consumenController

import (
	"golang-dev-kreditplus/models"
	"net/http"
	"os"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"golang-dev-kreditplus/helper"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func Index(c *gin.Context) {
	var consumen []models.Consumen

	models.DB.Find(&consumen)
	c.JSON(http.StatusOK, gin.H{"consumen": consumen})
}

func InsertConsumen(c *gin.Context){

	//Load env file
	if _, err := os.Stat(".env"); err == nil {
		//l.Info("Loading .env file")
		err := godotenv.Load()
		if err != nil {
			response := helper.APIResponse("Load env failed", http.StatusBadGateway, "error", err.Error())
			c.AbortWithStatusJSON(http.StatusBadGateway, response)
			return
		}
	}

	// Setup s3 uploader
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		response := helper.APIResponse("Load config failed", http.StatusBadRequest, "error", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)


	var consumen models.Consumen
	
	consumen.Nik = c.PostForm("nik")
	consumen.Fullname = c.PostForm("fullname")
	consumen.LegalName = c.PostForm("legal_name")
	consumen.BirthPlace = c.PostForm("birth_place")
	consumen.BirthDate = c.PostForm("birth_date")
	consumen.Salary = c.PostForm("salary")
	fileIdCardPhoto, errIdCardPhoto := c.FormFile("id_card_photo")

	if consumen.Nik == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Nik cannot be empty."})
		return
	}

	
	if fileIdCardPhoto != nil {
		//upload to s3
		if errIdCardPhoto == nil {

			T, openErr := fileIdCardPhoto.Open()
			if openErr != nil {
				response := helper.APIResponse("Can't open the file ID Card photo", http.StatusBadRequest, "Error", nil)
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
				return
			}

			// upload file ke aws s3
			resIdCard, uploadErr := uploader.Upload(context.TODO(), &s3.PutObjectInput{
				Bucket:      aws.String("test-dev-kreditplus"),
				Key:         aws.String("public/images/" + consumen.Nik + "-IDCard.jpg"),
				Body:        T,
				ACL:         "private",
				ContentType: aws.String(fileIdCardPhoto.Header.Get("Content-Type")),
			})

			if uploadErr != nil {
				response := helper.APIResponse("Failed Upload File ID Card Photo", http.StatusBadRequest, "Error", nil)
				c.AbortWithStatusJSON(http.StatusBadRequest, response)
				return
			}

			consumen.IdCardPhoto = resIdCard.Location

		}
	}
	

	

	err = models.DB.Create(&consumen).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Error add new consumen"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"consumen": consumen})
}