package controllers

import (
	"github.com/kataras/iris/v12"
	Config "github.com/mirzafaizan/gom-api/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DB connection
var dbLoan = Config.DB().Database("bankfinal").Collection("Loan")

func handleErr1(ctx iris.Context, err error) {
	ctx.JSON(iris.Map{"response": err.Error()})
}

func GetAllLoans(ctx iris.Context) {
	groupStage := bson.D{{"$group", bson.D{{"_id", "$product"}, {"total", bson.D{{"$sum", "$amount"}}}}}}

	showInfoCursor, err := dbLoan.Aggregate(ctx, mongo.Pipeline{groupStage})
	if err != nil {
		panic(err)
	}
	var showsWithInfo []bson.M
	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}

	//j, _ := json.MarshalIndent(showsWithInfo, "", " ")
	//log.Println(string(j))

	if err != nil {
		handleErr1(ctx, err)
		return
	}

	ctx.JSON(iris.Map{"response": showsWithInfo})

}

func CountLoans(ctx iris.Context) {
	//db.Loan.find({loan_status : 1}).count()
	query1 := bson.A{bson.D{{"$match", bson.D{{"loan_status", 1}}}}, bson.D{{"$count", "count"}}}
	facetStage := bson.D{{"$facet", bson.D{{"query1", query1}}}}

	showInfoCursor, err := dbLoan.Aggregate(ctx, mongo.Pipeline{facetStage})
	if err != nil {
		panic(err)
	}
	var showsWithInfo []bson.M
	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}

	ctx.JSON(iris.Map{"response": showsWithInfo})

}

func PortfolioOutstanding(ctx iris.Context) {
	groupStage := bson.D{{"$group", bson.D{{"_id", "$loan_status"}, {"total", bson.D{{"$sum", "$amount"}}}}}}
	showInfoCursor, err := dbLoan.Aggregate(ctx, mongo.Pipeline{groupStage})
	if err != nil {
		panic(err)
	}
	var showsWithInfo []bson.M
	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}

	//j, _ := json.MarshalIndent(showsWithInfo, "", " ")
	//log.Println(string(j))

	if err != nil {
		handleErr1(ctx, err)
		return
	}

	ctx.JSON(iris.Map{"response": showsWithInfo})

}
