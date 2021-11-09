package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Planet struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Name           string             `bson:"name" json:"name"`
	Climate        string             `bson:"climate" json:"climate"`
	Terrain        string             `bson:"terrain" json:"terrain"`
	ViewedQuantity int                `bson:"viewed_quantity" json:"viewed_quantity"`
}

func (planet *Planet) IsValid() error {
	if planet.Name == "" {
		return errors.New("Name not informed")
	}
	if planet.Climate == "" {
		return errors.New("Climate not informed")
	}

	if planet.Terrain == "" {
		return errors.New("Terrain not informed")
	}

	return nil
}
