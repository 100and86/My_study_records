package main

import (
	"encoding/json"
	"testing"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func TestFun(t *testing.T) {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	data, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		t.Logf("JSON marshaling failed: %s", err)
	}
	t.Logf("%s\n", data)
	titles := []struct{ Title string }{}
	if err:=json.Unmarshal(data,&titles);err !=nil{
		t.Logf("JSON unmarshaling failed: %s", err)
	}
	t.Log(titles)
}
