package main

import (
	"blog/util"
)

func main() {
	util.DbInit()

	defer util.Db.Close()

	err := util.Db.Ping()
	util.CheckError(err)

}
