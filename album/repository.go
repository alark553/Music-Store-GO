package album

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

const SERVER = "localhost:27017"

const DBNAME = "musicstore"

const DOCNAME = "albums"

// GetAlbums returns the list of Albums

func (r Repository) GetAlbums() Albums {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish a connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Albums{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to read results:", err)
	}
	return results
}

func (r Repository) AddAlbum(album Album) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	album.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(album)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (r Repository) DeleteAlbum(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}

	// Grab ID
	oid := bson.ObjectIdHex(id)

	// Remove user
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERR"
	}

	return "OK"
}

func (r Repository) UpdateAlbum(album Album) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	session.DB(DBNAME).C(DOCNAME).UpdateId(album.ID, album)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
