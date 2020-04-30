package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

var db *mgo.Session

func dialdb() error{
	var err error
	log.Println("Dialing mongodb: localhost")
	db, err = mgo.Dial("localhost")
	return err
}
func main()
{
	
}

