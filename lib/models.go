package lib

type Review struct {
	TransactionID     int    `bson:"transactionid,omitempty" json:"transactionid"`
	Rating            int    `bson:"rating,omitempty" json:"rating"`
	Review            string `bson:"review,omitempty" json:"review"`
	Language          string `bson:"language,omitempty" json:"language"`
	ImageURLFullxfull string `bson:"imageurl,omitempty" json:"imageurl"`
	CreateTimestamp   int    `bson:"created,omitempty" json:"created"`
	UpdateTimestamp   int    `bson:"updated,omitempty" json:"updated"`
}
