package controller

import (
	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/app/services/products"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

func IndexProducts(c *gin.Context) {
	all, err := products.ListAll()

	if err != nil {
		switch err.(type) {
		case *helpers.GenericError:
			c.JSON(err.(*helpers.GenericError).Code, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": all})
}

func ShowProducts(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := products.Details(models.Product{ID: id})

	if err != nil {
		switch err.(type) {
		case *helpers.GenericError:
			c.JSON(err.(*helpers.GenericError).Code, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func CreateProducts(c *gin.Context) {
	var params models.Product

	if err := c.BindJSON(&params); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	product, err := products.Create(params, true)

	if err != nil {
		switch err.(type) {
		case *helpers.GenericError:
			c.JSON(err.(*helpers.GenericError).Code, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	var (
		kafkaBrokers = []string{"kafka:29092"}
		topic        = "go-to-rails"
	)

	message, err := json.Marshal(product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Configuração do produtor Kafka
	config := kafka.WriterConfig{
		Brokers: kafkaBrokers,
		Topic:   topic,
	}

	writer := kafka.NewWriter(config)

	// Escreve a mensagem no Kafka
	err = writer.WriteMessages(c, kafka.Message{
		Value: message,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao enviar para Kafka",
			"error":   err.Error(),
		})
	}

	writer.Close()
	c.JSON(http.StatusCreated, gin.H{"data": product})
}

func UpdateProducts(c *gin.Context) {
	var params models.Product

	if err := c.BindJSON(&params); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	params.ID = id

	product, err := products.Update(params, true)

	if err != nil {
		switch err.(type) {
		case *helpers.GenericError:
			c.JSON(err.(*helpers.GenericError).Code, gin.H{"message": err.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}
