package router

import (
	"exercise/eaciit/variable"
	"fmt"

	"github.com/eaciit/dbox"
	_ "github.com/eaciit/dbox/dbc/mongo"
)

// DBConn a
type DBConn struct {
}

// Init a
func (db DBConn) Init() dbox.IConnection {

	fmt.Println("Connect To " + variable.DBHost)

	// Create Connection
	ci := dbox.ConnectionInfo{
		variable.DBHost,
		variable.DBName,
		variable.DBUser,
		variable.DBPass,
		nil,
	}

	conn, err := dbox.NewConnection("mongo", &ci)
	if err != nil {
		panic("Try To connect to " + variable.DBHost + " But Failed")
	}

	fmt.Println("Connect To " + variable.DBHost)

	err = conn.Connect()
	if err != nil {
		panic("Try To connect to " + variable.DBHost + " But Failed")
	}

	return conn
}
