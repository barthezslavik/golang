package dao

import "../models"

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MoviesDAO : describes a data access object for the Movie type
type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// COLLECTION
const (
	COLLECTION = "movies"
)

// Connect : Connect to the database
func (m *MoviesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// FindAll : Returns all stored movies and their details
func (m *MoviesDAO) FindAll() ([]models.Movie, error) {
	var movies []models.Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

// FindByID : Searches database for a single movie by id and returns its details
func (m *MoviesDAO) FindByID(id string) (models.Movie, error) {
	var movie models.Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

// Insert : Adds a new movie item and its details to the database
func (m *MoviesDAO) Insert(movie models.Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

// Delete : Deletes stored movie item and its details from the database
func (m *MoviesDAO) Delete(movie models.Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

// Update : Edits a stored movie item and its details
func (m *MoviesDAO) Update(movie models.Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
