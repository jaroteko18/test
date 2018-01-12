package controller

import (
	"exercise/eaciit/model"
	"exercise/eaciit/variable"
	"fmt"
	"time"

	"github.com/eaciit/dbox"
	"github.com/eaciit/knot/knot.v1"
	"gopkg.in/mgo.v2/bson"
)

type DBoxController struct {
}

func (db DBoxController) Index(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson
	fmt.Println(string("$eq"))
	// > db.profile.find({"_id":1})
	q, err := variable.DBConn.NewQuery().
		From("profile").
		// Where(dbox.Eq("_id", 1)).
		Cursor(nil)

	if err != nil {
		panic("Query Failed")
	}

	defer q.Close()

	users := []map[string]interface{}{}
	q.Fetch(&users, 0, false)

	return users
}

// Store a
func (db DBoxController) Store(r *knot.WebContext) interface{} {
	q := variable.DBConn.NewQuery().From("profile").SetConfig("multiexec", true).Save()
	// Make sure q is closed when exiting function
	defer q.Close()

	r.Config.OutputType = knot.OutputJson

	model := model.Profile

	e := r.GetPayload(&model)
	if e != nil {
		model.Name = e.Error()
	}

	model.ID = generateID()
	model.Age = generateAge(model.Birthday)

	newdata := map[string]interface{}{"data": model}

	err := q.Exec(newdata)
	if err != nil {
		panic("Query Failed")
	}

	return model
}

// Update a
func (db DBoxController) Update(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson

	// Baca Masukkan START
	model := model.Profile
	e := r.GetPayload(&model)
	if e != nil {
		model.Name = e.Error()
	}
	// Baca Masukkan END

	// Get Object START
	q, err := variable.DBConn.NewQuery().
		From("profile").
		Where(dbox.Eq("_id", model.ID)).
		Cursor(nil)

	if err != nil {
		panic("Query Failed xxxxx " + err.Error())
	}

	defer q.Close()

	users := []map[string]interface{}{}
	q.Fetch(&users, 0, false)
	// Get Object END

	if len(users) == 0 {
		mapD := map[string]string{"status": "-01", "message": "Profile Not Found"}

		return mapD
	}

	// Create new cursor
	qInsert := variable.DBConn.NewQuery().From("profile").SetConfig("multiexec", true).Save()
	// Make sure q is closed when exiting function
	defer qInsert.Close()
	model.Age = generateAge(model.Birthday)

	update := map[string]interface{}{"data": model}

	err = qInsert.Exec(update)
	// fmt.Println(err.Error())
	if err != nil {
		mapD := map[string]string{"status": "-01", "message": "Failed When Update"}

		return mapD
	}

	return update
	// {
	// 	"ID":"5a558ca5acf74d0e485fd7a9",
	// 	"name":"Budi",
	// 	"age":26,
	// 	"birthday":"1989-03-10T01:50:34+00:00",
	// 	"parent":["Agussss","Jeko"]
	// }
}

// Destroy a
func (db DBoxController) Destroy(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson

	// Baca Masukkan START
	var model struct {
		ID bson.ObjectId
	}

	e := r.GetPayload(&model)
	if e != nil {

	}
	// Baca Masukkan END

	// Delete query
	q := variable.DBConn.NewQuery().
		From("profile").
		Where(dbox.Eq("_id", model.ID)).
		Delete()
	// Make sure q is closed when exiting function
	defer q.Close()

	// Execute delete
	err := q.Exec(nil)
	if err != nil {
		mapD := map[string]string{"status": "-01", "message": "Delete Failed "}

		return mapD
	}

	mapD := map[string]string{"status": "00", "message": "Delete Success"}

	return mapD
}

// Search a
func (db DBoxController) Search(r *knot.WebContext) interface{} {
	r.Config.OutputType = knot.OutputJson

	// Baca Masukkan START
	var model struct {
		ID bson.ObjectId
	}

	e := r.GetPayload(&model)
	if e != nil {

	}

	fmt.Printf("%v \n", model.ID)
	// Baca Masukkan END

	// > db.Person.find({"_id":1})
	q, err := variable.DBConn.NewQuery().
		From("profile").
		Where(dbox.Eq("_id", model.ID)).
		Cursor(nil)

	if err != nil {
		panic("Query Failed")
	}

	defer q.Close()

	users := []map[string]interface{}{}
	q.Fetch(&users, 0, false)

	return users
}

func generateID() bson.ObjectId {
	fmt.Println(bson.NewObjectId())
	return bson.NewObjectId()
}

func generateAge(sBDate string) int {
	// "2006-01-02T15:04:05-0700"
	bDate, err := time.Parse(time.RFC3339, sBDate)

	if err != nil {
		panic(err)
	}

	return int(time.Since(bDate) / time.Hour / 24 / 365)
}
