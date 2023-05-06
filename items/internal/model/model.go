package model


//type Item itemservice.Item

type Item struct {
	ID          string `bson:"_id,omitempty"`
	Name        string `bson:"name,omitempty"`
	Description string `bson:"description,omitempty"`
	Damage      int    `bson:"damage,omitempty"`
	Healing     int    `bson:"healing,omitempty"`
	Protection  int    `bson:"protection,omitempty"`
}