package vars

import (
	"encoding/json"
	"log"
	"os"

	"github.com/StevenZack/databaser/data"

	"github.com/StevenZack/db"
	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/strToolkit"
)

var m = db.MustNewDB(strToolkit.Getrpath(fileToolkit.GetHomeDir())+".config"+string(os.PathSeparator)+"databaser", "cypher")

var connections = m.String("connections", "[]")

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
