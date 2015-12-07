package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
	Jwt    string        `json:"jwt" bson:"jwt"`
}

type Message struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Time string        `json:"time" bson:"time"`
	Body string        `json:"body" bson:"body"`
	User string        `json:"user" bson:"user"`
	Room string        `json:"room" bson:"room"`
	To   string        `json:"to" bson:"to"`
}
