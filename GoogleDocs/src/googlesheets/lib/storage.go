package lib

import (
	"time"

	"gopkg.in/mgo.v2"
)

// StorageItem is for saving objects to Mongo
type StorageItem struct {
	Name       string
	Value      string
	Date       time.Time
	Collection string `bson:"-"`
	DB         string `bson:"-"`
	Host       string `bson:"-"`
}

// NewStorageItem is s constructor func for StorageItem with host and DB default values
func NewStorageItem() StorageItem {
	s := StorageItem{}
	s.Host = "mongodb" // TODO fix this
	s.DB = "glaza"
	return s
}

// Save func saves Storage object in Mongo
func (s *StorageItem) Save() error {
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
