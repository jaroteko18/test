package variable

import (
	"os"
	"path/filepath"

	"github.com/eaciit/dbox"
)

// Setting For Web
const (
	WebHost      string = "localhost"
	WebPort      string = "9090"
	WebGoPath    string = "C:/Users/ASUS/go/src/"
	WebDir       string = WebGoPath + "exercise/eaciit/"
	WebUploadDir string = WebDir + "views/public/file/upload/"
	WebImageDir  string = WebDir + "views/public/file/static/image/"
	WebHTMLDir   string = WebDir + "views/"
	WebStatic    string = WebDir + "views/public/"
)

// User Default
const (
	UserSystem string = "0"
)

// Contain Message
const (
	MsgSuccess        string = "00"
	MsgFailedUnkown   string = "-01"
	MsgFailedValidate string = "-02"
)

// db Setting
const (
	DBHost string = "127.0.0.1"
	DBPort string = "27017"
	DBUser string = ""
	DBPass string = ""
	DBName string = "exercise"
)

// DBConn a
var DBConn dbox.IConnection

var (
	Dir = func() string {
		d, _ := os.Getwd()
		d = filepath.Join(d, "views")
		return d
	}()
)
