package mongostorage

import (
	"encoding/json"
	"time"

	"gopkg.in/mgo.v2"
)

// Item is for saving objects to Mongo
type Item struct {
	Name       string
	Value      string
	Date       time.Time
	JSON       interface{}
	Collection string `bson:"-"`
	DB         string `bson:"-"`
	Host       string `bson:"-"`
}

// NewItem is s constructor func for StorageItem with host and DB default values
func NewItem() Item {
	s := Item{}
	s.Host = "mongodb" // TODO fix this
	s.DB = "glaza"
	return s
}

// Save func saves Storage object in Mongo
func (s *Item) Save() error {
	if s.JSON != nil {
		var m interface{}
		err := json.Unmarshal([]byte(s.JSON.(string)), &m)
		if err != nil {
			return err
		}
		s.JSON = m
	}

	session, err := mgo.Dial(s.Host)
	if err != nil {
		return err
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(s.DB).C(s.Collection)
	s.Date = time.Now()
	err = c.Insert(s)
	if err != nil {
		return err
	}
	return nil
}
