package util

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StringToInteger(i string) int {
	r, _ := strconv.Atoi(i)
	return r
}

func IntegerToString(i int) string {
	r := strconv.Itoa(i)
	return r
}

func StringToObjecID(i string) primitive.ObjectID {
	oid, _ := primitive.ObjectIDFromHex(i)
	return oid
}
