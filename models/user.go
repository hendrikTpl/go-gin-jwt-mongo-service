package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username      *string            `json:"username" bson:"username" validate:"required,min=4,max=50"`
	FirstName     *string            `json:"first_name" bson:"first_name" validate:"required,min=2,max=100"`
	LastName      *string            `json:"last_name" bson:"last_name" validate:"max=100"`
	Password      *string            `json:"password" bson:"password" validate:"required,min=6"`
	Email         *string            `json:"email" bson:"email"`
	Phone         *string            `json:"phone" bson:"phone"`
	Token         *string            `json:"token" bson:"token"`
	Level         *string            `json:"level" bson:"level" validate:"required,eq=ADMIN|eq=USER"`
	RefreshToken  *string            `json:"refresh_token" bson:"refresh_token"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
}
