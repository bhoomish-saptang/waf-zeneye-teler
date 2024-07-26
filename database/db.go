package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Rule struct {
	Rulename  string `json:"rulename,omitempty"`
	Field     string `json:"field" bson:"field"`
	FieldVal  string `json:"value" bson:"value"`
	Condition string `json:"condition" bson:"condition"`
	Action    string `json:"action" bson:"action"`
}

type CustomRules struct {
	IP        string `json:"ip" bson:"ip"`
	Rules     []Rule `json:"rules" bson:"rules"`
	CreatedAt string `json:"createdat" bson:"createdat"`
	URL       string `json:"url" bson:"url"`
}

var client, dberr = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
var RuleDB = client.Database("WAF").Collection("Testing")
var LogDB = client.Database("WAF").Collection("RequestLogs")

func InsertOneQuery(todo Rule, collection *mongo.Collection) error {
	if dberr != nil {
		fmt.Println("DB Error: ", dberr)
	}
	res, insErr := collection.InsertOne(context.TODO(), todo)
	if insErr != nil {
		fmt.Println(insErr)
		return insErr
	} else {
		fmt.Println("Inserted into: ", collection.Name(), res.InsertedID)
		fmt.Println()
		return insErr
	}
}

func FindOneQuery(filter, projection interface{}, collection *mongo.Collection, checkvar bool) (map[string]string, bool) {

	if dberr != nil {
		fmt.Println("DB Error: ", dberr)
	}
	var interMap map[string]string

	doc := collection.FindOne(context.Background(), filter, options.FindOne().SetSort(projection)).Decode(&interMap)

	if doc != nil {
		if doc == mongo.ErrNoDocuments {
			return map[string]string{}, false
		}
	}
	if len(interMap) != 0 {
		checkvar = true
	}
	return interMap, checkvar
}
