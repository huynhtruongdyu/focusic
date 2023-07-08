package domains

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	BaseEntity
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
}
