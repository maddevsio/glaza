package lib

import (
	"gopkg.in/mgo.v2"
)

// Storage is for saving objects to Mongo
type Storage struct {
	Name  string
	Value string
}

// Save func saves Storage object in Mongo
func (s *Storage) Save() error {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		return err
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("docs")
	err = c.Insert(s)
	if err != nil {
		return err
	}
	return nil
}
