package products

import (
	"context"
	"ms-go/app/helpers"
	"ms-go/app/models"
	"ms-go/db"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(data models.Product, isAPI bool) (*models.Product, error) {
	collection := db.Connection()

	if data.ID == 0 {
		var max models.Product

		opts := options.FindOne().SetSort(bson.D{{Key: "created_at", Value: -1}})

		err := collection.FindOne(context.TODO(), bson.D{}, opts).Decode(&max)

		if err != nil && err.Error() != "mongo: no documents in result" {
			return nil, &helpers.GenericError{Msg: "Error fetching max ID", Code: http.StatusInternalServerError}
		}

		data.ID = max.ID + 1
	}

	if err := data.Validate(); err != nil {
		return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusUnprocessableEntity}
	}

	data.CreatedAt = time.Now()
	data.UpdatedAt = data.CreatedAt

	var product models.Product

	err := collection.FindOne(context.TODO(), bson.M{"id": data.ID}).Decode(&product)

	if err == nil {
		if isAPI {
			return nil, &helpers.GenericError{Msg: "Produto já existe", Code: http.StatusConflict}
		} else {
			setUpdate(&data, &product)

			updateResult := collection.FindOneAndUpdate(context.TODO(), bson.M{"id": data.ID}, bson.M{"$set": data})
			
			if updateResult.Err() != nil {
				return nil, &helpers.GenericError{Msg: updateResult.Err().Error(), Code: http.StatusInternalServerError}
			}
		}
	} else {
		_, err := collection.InsertOne(context.TODO(), data)
		if err != nil {
			return nil, &helpers.GenericError{Msg: err.Error(), Code: http.StatusInternalServerError}
		}
	}

	defer db.Disconnect()

	return &data, nil
}