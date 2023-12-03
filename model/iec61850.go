package model

type Signal struct {
	Key   string `json:"key" bson:"key"`
	Value int    `json:"value" bson:"value"`
}
