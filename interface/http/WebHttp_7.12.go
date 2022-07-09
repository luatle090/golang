package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var storeList = template.Must(template.New("list").Parse(`
<html>
<body>
	<h1>{{.TotalCount}} mặt hàng</h1>
		<table>
			<tr style='text-align: left'>
				<th>Item</th>
				<th>Price</th>
			</tr>
			{{range $name, $price := .}}
			<tr>
				<td>{{$name}}</td>
				<td>{{$price}}</td>
			</tr>
			{{end}}
		</table>
</body>
</html>
`))

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func main() {
	db := database{"shoes": 50, "socks": 10}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/read/", db.read)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/delete/", db.delete)
	http.HandleFunc("/update/", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	storeList.Execute(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("item")
	price, ok := db[key]
	if !ok {
		msg := fmt.Sprintf("no such item: %q\n", key)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s", price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	name := strings.TrimPrefix(req.URL.Path, "/update/")
	priceForm := req.FormValue("price")

	_, ok := db[name]

	if !ok {
		msg := fmt.Sprintf("no such name: %q\n", name)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	price, error := strconv.ParseFloat(priceForm, 8)

	if error != nil {
		msg := fmt.Sprintf("price has not right format\n")
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	db[name] = dollars(price)
	fmt.Fprintf(w, "item has been update %s : %.2f", name, price)
	w.WriteHeader(http.StatusCreated)
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	name := strings.TrimPrefix(req.URL.Path, "/read/")
	price, ok := db[name]
	if !ok {
		msg := fmt.Sprintf("no such item: %q\n", name)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s : %s", name, price)
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	name := strings.TrimPrefix(req.URL.Path, "/delete/")
	_, ok := db[name]
	if !ok {
		msg := fmt.Sprintf("no such item: %q\n", name)
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	delete(db, name)
	fmt.Fprintf(w, "item has been deleted %q", name)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	priceForm := req.FormValue("price")
	item := req.FormValue("name")
	_, ok := db[item]

	if ok {
		msg := fmt.Sprintf("item has already in database: %q\n", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	price, error := strconv.ParseFloat(priceForm, 8)

	if error != nil {
		msg := fmt.Sprintf("price has not right format\n")
		http.Error(w, msg, http.StatusNotFound)
		return
	}
	db[item] = dollars(price)
	fmt.Fprintf(w, "item has been added %s : %.2f", item, price)
	w.WriteHeader(http.StatusCreated)
}
