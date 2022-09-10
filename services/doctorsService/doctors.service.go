package doctorsService

import (
	"context"
	"log"
	"os"

	databaseConfig "fita-test-02/config/databaseConfig"
	models "fita-test-02/models"

	"go.mongodb.org/mongo-driver/bson"
)

func ReturnAllDoctors() ([]*models.Doctors, error) {
	var doctors []*models.Doctors
	ctx := context.TODO()

	client := databaseConfig.GetClient()

	fita01DB := client.Database(os.Getenv("DB_DATABASE"))
	doctorsCollection := fita01DB.Collection("doctors")

	filter := bson.M{}
	cur, err := doctorsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(context.TODO()) {
		var doctor models.Doctors
		err = cur.Decode(&doctor)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		doctors = append(doctors, &doctor)
	}
	defer client.Disconnect(ctx)
	return doctors, nil

}
