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

type ConnectList []Connection

func (a ConnectList) Len() int      { return len(a) }
func (a ConnectList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ConnectList) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}
