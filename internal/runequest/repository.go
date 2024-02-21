package runequest

import (
	"context"
	"errors"
	"github.com/ahonesa/new-character-repository/internal/db"
	"github.com/ahonesa/new-character-repository/internal/model/runequest"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type RuneQuestCharacterRepository struct {
	collection *mongo.Collection
}

func NewRuneQuestCharacterRepository() *RuneQuestCharacterRepository {
	client := db.GetMongoDBClient() // Assumes db.GetMongoDBClient() initializes and returns a MongoDB client
	return &RuneQuestCharacterRepository{
		collection: client.Database("rqg-depo-prod").Collection("characters"), // Adjust the database and collection names as needed
	}
}

func (r *RuneQuestCharacterRepository) GetByID(characterId string) (*runequest.RuneQuestCharacter, error) {
	var character runequest.RuneQuestCharacter
	filter := bson.M{"characterId": characterId}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, filter).Decode(&character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

func (r *RuneQuestCharacterRepository) GetAll() ([]runequest.RuneQuestCharacter, error) {
	var characters []runequest.RuneQuestCharacter

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var character runequest.RuneQuestCharacter
		if err = cursor.Decode(&character); err != nil {
			log.Fatal(err)
		}
		characters = append(characters, character)
	}

	return characters, nil
}

func (r *RuneQuestCharacterRepository) Create(ctx context.Context, character *runequest.RuneQuestCharacter) error {
	_, err := r.collection.InsertOne(ctx, character)
	return err
}

func (r *RuneQuestCharacterRepository) Update(ctx context.Context, characterId string, character *runequest.RuneQuestCharacter) error {
	filter := bson.M{"character.characterId": characterId}

	update := bson.M{"$set": bson.M{"character": character}}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("character not found")
	}

	return nil
}

func (r *RuneQuestCharacterRepository) Delete(ctx context.Context, characterId string) error {
	filter := bson.M{"characterId": characterId}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no character found with given ID")
	}

	return nil
}
