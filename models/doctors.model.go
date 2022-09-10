package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Doctors struct {
	ID                    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name                  string             `bson:"name" json:"name"`
	TimezoneGmt           string             `json:"timezone_gmt"`
	TimezoneArea          string             `json:"timezone_area"`
	DayOfWeek             string             `bson:"day_of_week" json:"day_of_week"`
	AvailableAt           string             `bson:"available_at" json:"available_at"`
	AvailableAtMeridiem   string             `bson:"available_at_meridiem" json:"available_at_meridiem"`
	AvailableUtil         string             `bson:"available_util" json:"available_util"`
	AvailableUtilMeridiem string             `bson:"available_util_meridiem" json:"available_util_meridiem"`
}
