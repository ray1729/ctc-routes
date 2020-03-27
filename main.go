package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"text/template"
)

const TMPL = `
static {
    {{range .Routes}}
    routes.add(new Route(
        {{.Name | printf "%q"}},
        {{.Distance | printf "%q"}},
        {{.Author | printf "%q"}},
        {{.Description | printf "%q"}},
        {{.Link | printf "%q"}},
        {{.Start | printf "%q"}},
        ));
    {{end}}
}
`

var tmpl = template.Must(template.New("java").Parse(TMPL))

func main() {
	log.SetFlags(0)
	if len(os.Args) != 3 {
		log.Fatalf("Usage: %s ROUTES.csv OUTPUT\n", os.Args[0])
	}
	routes, err := read_routes_csv(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	out, err := os.OpenFile(os.Args[2], os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	err = tmpl.Execute(out, routes)
	if err != nil {
		log.Fatal(err)
	}
}

func read_routes_csv(path string) (*Routes, error) {
	var routes []Route
	rdr, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer rdr.Close()
	r := csv.NewReader(rdr)
	skip_header := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if skip_header {
			skip_header = false
			continue
		}
		routes = append(routes, Route{
			Name:        record[3],
			Distance:    record[0],
			Author:      record[2],
			Description: record[5],
			Link:        record[4],
			Start:       record[1],
		})
	}
	return &Routes{routes}, nil
}

type Route struct {
	Name, Distance, Author, Description, Link, Start string
}

type Routes struct {
	Routes []Route
}
