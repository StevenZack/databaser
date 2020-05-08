package data

type Connection struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Dsn  string `json:"dsn,omitempty"`
	DB   string `json:"db,omitempty"`
}

const (
	TypeMySQL      = "mysql"
	TypeClickhouse = "clickhouse"
	TypeMongoDB    = "mongodb"
)
