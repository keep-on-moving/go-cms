package models

type Page struct {
	Website string
	Email  string
}

func GetPage() Page {
	var rtr = Page{Website:"hellowbeego.com", Email:"models@beego.com"}

	return  rtr
}