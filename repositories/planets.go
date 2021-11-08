package repositories

import (
	"context"
	"planets-api/config"
	"planets-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type planetsRepository struct {
	PlanetsCollection *mongo.Collection
}

type PlanetsRepository interface {
	GetAll() ([]models.Planet, error)
	GetById(id string) (models.Planet, error)
	GetByName(name string) ([]models.Planet, error)
	Create(planet *models.Planet) error
	Update(id string, planet models.Planet) error
	Delete(id string) error
}

var _ PlanetsRepository = &planetsRepository{}

const (
	COLLECTION = "planets"
)

func NewPlanetsRepository(client *mongo.Client) *planetsRepository {

	planetsCollection := client.Database(config.MONGO_DATABASE).Collection(COLLECTION)

	return &planetsRepository{
		PlanetsCollection: planetsCollection,
	}
}

func (repository *planetsRepository) GetAll() ([]models.Planet, error) {
	var planet models.Planet
	var planets []models.Planet

	filter := bson.D{}

	cursor, err := repository.PlanetsCollection.Find(context.TODO(), filter)
	if err != nil {
		defer cursor.Close(context.TODO())
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&planet)
		if err != nil {
			return nil, err
		}
		planets = append(planets, planet)
	}

	return planets, nil
}

func (repository *planetsRepository) GetById(id string) (models.Planet, error) {
	var planet models.Planet

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return planet, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objID}}

	err = repository.PlanetsCollection.FindOne(context.TODO(), filter).Decode(&planet)
	if err != nil {
		return planet, err
	}
	return planet, nil
}

func (repository *planetsRepository) GetByName(name string) ([]models.Planet, error) {
	var planet models.Planet
	var planets []models.Planet

	filter := bson.M{
		"name": bson.M{
			"$regex": primitive.Regex{
				Pattern: name,
				Options: "i",
			},
		},
	}

	cursor, err := repository.PlanetsCollection.Find(context.TODO(), filter)
	if err != nil {
		defer cursor.Close(context.TODO())
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&planet)
		if err != nil {
			return nil, err
		}
		planets = append(planets, planet)
	}

	return planets, nil
}

func (repository *planetsRepository) Create(planet *models.Planet) error {
	planet.ID = primitive.NewObjectID()

	_, erro := repository.PlanetsCollection.InsertOne(context.TODO(), planet)
	if erro != nil {
		return erro
	}
	return nil
}

func (repository *planetsRepository) Update(id string, planet models.Planet) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objID}}

	_, erro := repository.PlanetsCollection.ReplaceOne(context.TODO(), filter, planet)
	if erro != nil {
		return erro
	}
	return nil
}

func (repository *planetsRepository) Delete(id string) error {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: objID}}

	_, err = repository.PlanetsCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
