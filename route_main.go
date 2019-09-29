package main

import (
	"chitchat/data"
	"net/http"
	"strconv"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	var p int
	var err error
	query := request.URL.Query()
	if page,ok := query["page"]; ok{
		if len(page) > 0 {
			if p,err = strconv.Atoi(page[0]);err != nil{
				p = 1
			}
		}
	}
	if p < 1{p=1}
	res := make(map[string]interface{})
	res["pre"] = p-1
	res["next"] = p+1
	threadsList, err := data.Threads(p,5)
	if p == 1{
		res["pre"] = 1
	}
	if len(threadsList) < 5{
		res["next"] = p
	}
	res["list"] = threadsList
	if err != nil {
		error_message(writer, request, "Cannot get threads")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, res, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, res, "layout", "private.navbar", "index")
		}
	}
}
