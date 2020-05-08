package vars

import (
	"os"

	"github.com/StevenZack/db"
	"github.com/StevenZack/tools/fileToolkit"
	"github.com/StevenZack/tools/strToolkit"
)

var m = db.MustNewDB(strToolkit.Getrpath(fileToolkit.GetHomeDir())+".config"+string(os.PathSeparator)+"databaser", "cypher")

var connections = m.String("connections", "[]")
