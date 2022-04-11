package company_repository

import (
	"fmt"
	"github.com/techagentng/boiler-go/domain/helpers"
	domain "github.com/techagentng/boiler-go/domain/company-repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"log"
)

// RepositoryDB struct
type RepositoryDB struct {
	db *mongo.Client
}

// NewCompanyGatewayRepositoryDB function to initialize RepositoryDB struct
func NewCompanyGatewayRepositoryDB(client *mongo.Client) *RepositoryDB {
	return &RepositoryDB{
		db: client,
	}
}

func (companyRepo *RepositoryDB) CheckIfEmailExists(email string) (bson.M, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Checking if email exists: %s ...", email))
	collection := companyRepo.db.Database("company-repo").Collection("user")

	var result bson.M
	err := collection.FindOne(
		context.TODO(),
		bson.D{{"email", email}},
	).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal(err)
	}
	fmt.Printf("found document %v", result)
	return result, nil
}

func (companyRepo *RepositoryDB) CheckIfUserExists(userReference string) (bson.M, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Checking if user exists : %s ...", userReference))
	collection := companyRepo.db.Database("company-repo").Collection("user")

	var result bson.M
	err := collection.FindOne(
		context.TODO(),
		bson.D{{"reference", userReference}},
	).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal(err)
	}
	fmt.Printf("found document %v", result)
	return result, nil
}

func (companyRepo *RepositoryDB) CreateUser(user *domain.User) (*domain.User, error) {
	helpers.LogEvent("INFO", fmt.Sprintf("Authorising Merchant with reference: %s ...", user))

	collection := companyRepo.db.Database("company-repo").Collection("user")
	result, err := collection.InsertOne(context.TODO(), user)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return user, err
}
