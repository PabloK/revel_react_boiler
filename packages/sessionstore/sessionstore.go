package sessionstore

import (
	"gopkg.in/mgo.v2"
)

var instantiated *mgo.Session = nil

func GetSession(connectionString string) *mgo.Session {
	if instantiated == nil {
		var session *mgo.Session
		session, err := mgo.Dial(connectionString)
		if err != nil {
			panic(err)
		}
		session.SetMode(mgo.Monotonic, true)

		instantiated = session

	}

	return instantiated.Copy()
}
