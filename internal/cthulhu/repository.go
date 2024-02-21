package cthulhu

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/ahonesa/new-character-repository/internal/db"
	"github.com/ahonesa/new-character-repository/internal/model/cthulhu" // Corrected import path
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CthulhuCharacterRepository struct {
	collection *mongo.Collection
}

func NewCthulhuCharacterRepository() *CthulhuCharacterRepository {
	client := db.GetMongoDBClient() // This line should use the imported db package
	return &CthulhuCharacterRepository{
		collection: client.Database("rqg-depo-prod").Collection("cthulhu-characters"),
	}
}

func (r *CthulhuCharacterRepository) GetByID(characterId string) (*cthulhu.CthulhuCharacter, error) {
	var character cthulhu.CthulhuCharacter
	filter := bson.M{"character.characterId": characterId}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, filter).Decode(&character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

func (r *CthulhuCharacterRepository) GetAll() ([]cthulhu.CthulhuCharacter, error) {
	var characters []cthulhu.CthulhuCharacter

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var character cthulhu.CthulhuCharacter
		if err = cursor.Decode(&character); err != nil {
			log.Fatal(err)
		}
		characters = append(characters, character)
	}

	return characters, nil
}

// Create inserts a new Cthulhu character into the database.
func (r *CthulhuCharacterRepository) Create(ctx context.Context, ownerId string, cthulhuChar *cthulhu.CthulhuCharacter) error {
	// Assign the OwnerId to the character.
	cthulhuChar.OwnerId = ownerId

	// Insert the new CthulhuCharacter document into the collection.
	_, err := r.collection.InsertOne(ctx, cthulhuChar)
	return err
}

// Update updates an existing character based on CharacterId.
func (r *CthulhuCharacterRepository) Update(ctx context.Context, characterId string, cthulhuChar *cthulhu.CthulhuCharacter) error {
	filter := bson.M{"character.characterId": characterId}

	// Ensure the update document matches the structure of CthulhuCharacter.
	update := bson.M{"$set": bson.M{"character": cthulhuChar.Character}}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("character not found")
	}

	return nil
}

func (r *CthulhuCharacterRepository) Delete(ctx context.Context, characterId string) error {
	filter := bson.M{"character.characterId": characterId}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no character found with given ID")
	}

	return nil
}
