package vars

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/StevenZack/databaser/data"
)

func GetConnections() []data.Connection {
	vs := []data.Connection{}
	e := json.Unmarshal([]byte(connections.Get()), &vs)
	if e != nil {
		log.Fatal(e)
	}
	return vs
}

func SetConnections(vs []data.Connection) {
	if vs == nil {
		connections.Post("[]")
		return
	}
	b, e := json.Marshal(vs)
	if e != nil {
		log.Fatal(e)
	}
	connections.Post(string(b))
}

func AppendConnection(c data.Connection) {
	vs := GetConnections()
	vs = append(vs, c)
	SetConnections(vs)
}

func GetConnectionByName(name string) (data.Connection, error) {
	for _, v := range GetConnections() {
		if v.Name == name {
			return v, nil
		}
	}
	return data.Connection{}, errors.New("connection '" + name + "' not found")
}
