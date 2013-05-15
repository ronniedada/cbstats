package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func serveBarChart(w http.ResponseWriter, r *http.Request) {
	ddoc, view := mux.Vars(r)["ddoc"], mux.Vars(r)["view"]

	var args map[string]interface{}

	groupLevel := r.FormValue("group_level")
	i, err := strconv.Atoi(groupLevel)

	if err != nil {
		args = map[string]interface{}{"stale": "update_after"}
	} else {
		args = map[string]interface{}{"stale": "update_after",
			"group_level": i}
	}

	vr, err := fetchView(ddoc, view, args)

	if err != nil {
		log.Printf("err: %v", err)
		return
	}

	if err != nil {
		showError(w, r, err.Error(), 404)
		return
	}

	histo := vr.histo()
	if len(histo) == 0 {
		showError(w, r, "oops, empty results", 404)
		return
	}
	axes := []string{"x", "y"}

	yIndex := r.FormValue("y_index")
	i, _ = strconv.Atoi(yIndex)
	vr.bar(histo[i], axes, w)
}

func serveLineChart(w http.ResponseWriter, r *http.Request) {
	ddoc, view := mux.Vars(r)["ddoc"], mux.Vars(r)["view"]

	var args map[string]interface{}

	groupLevel := r.FormValue("group_level")
	i, err := strconv.Atoi(groupLevel)

	if err != nil {
		args = map[string]interface{}{"stale": "update_after"}
	} else {
		args = map[string]interface{}{"stale": "update_after",
			"group_level": i}
	}

	vr, err := fetchView(ddoc, view, args)

	if err != nil {
		showError(w, r, err.Error(), 404)
		return
	}

	histo := vr.histo()
	if len(histo) == 0 {
		showError(w, r, "oops, empty results", 404)
		return
	}
	axes := []string{"x", "y"}

	yIndex := r.FormValue("y_index")
	i, _ = strconv.Atoi(yIndex)
	vr.line(histo[i], axes, w)
}

func serveStackedBarChart(w http.ResponseWriter, r *http.Request) {
	ddoc, view := mux.Vars(r)["ddoc"], mux.Vars(r)["view"]

	var args map[string]interface{}

	groupLevel := r.FormValue("group_level")
	i, err := strconv.Atoi(groupLevel)

	if err != nil {
		args = map[string]interface{}{"stale": "update_after"}
	} else {
		args = map[string]interface{}{"stale": "update_after",
			"group_level": i}
	}

	vr, err := fetchView(ddoc, view, args)

	if err != nil {
		showError(w, r, err.Error(), 404)
		return
	}

	xIndex, err := strconv.Atoi(r.FormValue("x_index"))
	rangeIndex, err := strconv.Atoi(r.FormValue("range_index"))
	ran, err := strconv.Atoi(r.FormValue("range"))

	vr.stackedBar(xIndex, rangeIndex, ran, w)
}
