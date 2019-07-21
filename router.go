package main


import (
	"net/http"
)

 
// Route is for routing 
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}
 
// Routes is for routing
type Routes []Route


var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "TodoIndex",
        "GET",
        "/todos",
        TodoIndex,
    },
    Route{
        "TodoShow",
        "GET",
        "/todos/{todoId}",
        TodoShow,
    },
}