package database

import (
	"go.mongodb.org/mongo-driver/bson"
)

func RetrieveRule(RuleName []Rule, checkvar bool) (map[string]string, bool) {
	filter := bson.M{"rulename": RuleName}
	projection := bson.D{{Key: "_id", Value: -1}}
	dataMap, checkvar := FindOneQuery(filter, projection, RuleDB, checkvar)
	return dataMap, checkvar
}
