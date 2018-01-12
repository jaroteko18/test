package main

import (
	"exercise/eaciit/variable"
	"exercise/router"
)

func main() {

	variable.DBConn = router.DBConn{}.Init()
	router.Router{}.Init()
	// router.Router{}.Init()
}
