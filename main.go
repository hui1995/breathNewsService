package main

import (
	"breathNewsService/Router"
)

func main() {
	//defer Mysql.DB.Close()
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	Router.InitRouter()
}
