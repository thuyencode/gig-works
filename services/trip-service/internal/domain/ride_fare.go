package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PackageSlug string

type RideFareModel struct {
	ID                primitive.ObjectID
	UserID            string
	PackageSlug       PackageSlug
	TotalPriceInCents float64
}
