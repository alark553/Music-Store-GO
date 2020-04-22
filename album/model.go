package album

import "gopkg.in/mgo.v2/bson"

type Album struct {
	ID     bson.ObjectId `bson:"_id"`
	Title  string        `json:"title"`
	Artist string        `json:"artist"`
	Year   int32         `json:"year"`
}

type Albums []Album
