package model

type Signal struct {
	Key   string `json:"key" bson:"key"`
	Value string `json:"value" bson:"value"`
}
