package topicModel

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Topic is a type for topic collection in mongo
type Topic struct {
	ID   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `bson:"name" json:"name"`
}

// Topics is a array of Topic
type Topics []Topic

var (
	mongoDB *mgo.Database
)

const (
	// CollectionTopic holds the name of the topics collection
	collection = "topics"
)

// SetMongoDB is a function to set a mgo.Database
func SetMongoDB(m *mgo.Database) {
	mongoDB = m
}

func checkMongoDBNotNull() error {
	if mongoDB == nil {
		return errors.New("mongoDB client is null please set it before use")
	}
	return nil
}

// Save  function for topic
func (t *Topic) Save() error {
	err := checkMongoDBNotNull()
	err = mongoDB.C(collection).Insert(t)
	return err
}

// GetAll function for get all topics
func (t *Topics) GetAll() error {
	err := checkMongoDBNotNull()
	err = mongoDB.C(collection).Find(nil).All(t)
	return err
}

// GetOneByID function for get topic by id
func (t *Topic) GetOneByID(id string) error {
	err := mongoDB.C(collection).FindId(bson.ObjectIdHex(id)).One(t)
	return err
}

// GetOneByName function for get topic by name
func (t *Topic) GetOneByName(name string) error {
	err := mongoDB.C(collection).Find(bson.M{"name": name}).One(t)
	return err
}

// Update function is a function that will query and update value of topic
func (t *Topic) Update(doc bson.M) error {
	query := bson.M{"_id": t.ID}
	err := mongoDB.C(collection).Update(query, doc)
	return err
}

// Delete is a function that can delete a topic
func (t *Topic) Delete() error {
	err := mongoDB.C(collection).RemoveId(t.ID)
	return err
}
