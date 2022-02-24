package controllers

import (
	"context"
	Config "github.com/mirzafaizan/gom-api/config"
	Models "github.com/mirzafaizan/gom-api/models"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

// DB connection
var dbYear = Config.DB().Database("bankfinal").Collection("yearlyUpdate")

func handleErr2(ctx iris.Context, err error) {
	ctx.JSON(iris.Map{"response": err.Error()})
}

// GetAllUsers ...
// Method:   GET
// Resource: this to get all all users
func GetYearlyUpdates(ctx iris.Context) {
	results := []*Models.Year{}
	c := context.TODO()
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := dbYear.Find(c, bson.D{{}})
	if err != nil {
		handleErr(ctx, err)
		return
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(c) {
		// create a value into which the single document can be decoded
		var elem Models.Year
		err := cur.Decode(&elem)
		if err != nil {
			handleErr(ctx, err)
			return
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		handleErr(ctx, err)
		return
	}
	// Close the cursor once finished
	cur.Close(c)
	ctx.JSON(iris.Map{"response": results})
}
