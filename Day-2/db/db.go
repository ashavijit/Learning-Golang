package db

import(
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//item represents in a db
type item struct {
	ID string `bson:"_id"`
	Value int `bson:"value"`
}
var db *mgo.Database

func init(){
	session, err := mgo.Dial("db_url")
	if err!=nil{
		log.fatal("Error connecting to MongoDB: ", err)
	}
	db = session.DB("api_db")
}

func collection() *mgo.Collection{
	return db.C("items")
}

/* func GetAllItems() ([]item, error){
	var items []item
	err := collection().Find(bson.M{}).All(&items)
	return items, err
} */

func GetAllItems() ([]item, error){
	var items []item
	err := collection().Find(bson.M{}).All(&items)
	return items, err
}

func GetOneItem(id string) (item, error){
	rs := item{}
	if err := collection().FindId(id).One(&rs); err != nil {
		return rs, err
	}
	return rs, nil
}

func Save (item Item) error {
	return collection().Insert(item)
}
	
func Remove(id string) error {
	return collection().RemoveId(id)
}

func Upadate (	item Item) error {
	return collection().UpdateId(item.ID, item)
}
