package controllers

import (
	"context"

	Config "github.com/mirzafaizan/gom-api/config"
	Models "github.com/mirzafaizan/gom-api/models"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

// DB connection
var dbLoan = Config.DB().Database("admin-dashboard").Collection("Loan")

func handleErr1(ctx iris.Context, err error) {
	ctx.JSON(iris.Map{"response": err.Error()})
}

func GetAllLoans(ctx iris.Context) {
	var results = []*Models.Loans{}
	c := context.TODO()
	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := dbLoan.Find(c, bson.D{{}})
	if err != nil {
		handleErr1(ctx, err)
		return
	}
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(c) {
		// create a value into which t he single document can be decoded
		var elem Models.Loans
		err := cur.Decode(&elem)
		if err != nil {
			handleErr1(ctx, err)
			return
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		handleErr1(ctx, err)
		return
	}
	// Close the cursor once finished
	cur.Close(c)
	ctx.JSON(iris.Map{"response": results})
}
