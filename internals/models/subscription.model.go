package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Subscription struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID         primitive.ObjectID `bson:"user_id" json:"user_id"`
	PlanName       string             `bson:"plan_name" json:"plan_name"`
	StartDate      primitive.DateTime `bson:"start_date" json:"start_date"`
	EndDate        primitive.DateTime `bson:"end_date" json:"end_date"`
	IsActive       bool               `bson:"is_active" json:"is_active"`
	AutoRenew      bool               `bson:"auto_renew" json:"auto_renew"`
	PaymentMethod  string             `bson:"payment_method" json:"payment_method"`
}
