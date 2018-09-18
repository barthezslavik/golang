package main

import (
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Car struct {
	Name string
	Type string
}

func main() {
	app := iris.Default()

	// Database connection
	session, err := mgo.Dial("localhost")
	if nil != err {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Database name and collection name
	// car-db is database name car is collation name
	c := session.DB("car-db").C("car")
	c.Insert(&Car{"Audi", "Luxury car"})

	// Get /car
	app.Get("/car", func(ctx iris.Context) {
		result := Car{}
		err = c.Find(bson.M{"name": "Audi"}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		ctx.JSON(result)
	})

	app.Run(iris.Addr(":8080"))
}
