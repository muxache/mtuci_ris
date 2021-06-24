package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/muxache/mtuci_ris/controller"
	"github.com/muxache/mtuci_ris/data"
	"github.com/muxache/mtuci_ris/service"
)

var db *sql.DB

func main() {
	db_data := controller.DBData{
		Server:   "java.ip41.ru",
		Port:     1433,
		User:     "operator_api",
		Password: "******",
		Database: "DEZ",
	}
	db = db_data.ConnectToDB()
	router := mux.NewRouter()
	router.HandleFunc("/orders", OrdersList)
	router.HandleFunc("/", IndexList)
	router.HandleFunc("/edit/{id:[0-9]+}", EditPage).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", EditHandler).Methods("POST")
	// http.HandleFunc("/orders", OrdersList)
	// http.HandleFunc("/", IndexList)
	http.Handle("/", router)

	http.ListenAndServe(":8181", nil)

}

func IndexList(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("html/index.html")
	tmpl.Execute(w, nil)
}

func OrdersList(w http.ResponseWriter, r *http.Request) {

	res, err := service.SelectFromORDERS(db)
	if err != nil {
		log.Println(err)
	}
	// for _, r := range res {
	// 	fmt.Println(r.Order_ID, r.Description, r.Order_date.Format("2006 Jan 2"), r.Close_date.Format("2006 Jan 2"), r.Master_date.Format("2006 Jan 2"))
	// }
	tmpl, _ := template.ParseFiles("html/orders.html")
	tmpl.Execute(w, res)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println(err)
	}

	description := r.FormValue("description")

	master_date_str := r.FormValue("master_date")
	master_date, err := time.Parse("Jan 02, 2006", master_date_str)
	if err != nil {
		log.Println(err)
	}
	close_date_str := r.FormValue("close_date")
	close_date, err := time.Parse("Jan 02, 2006", close_date_str)
	if err != nil {
		log.Println(err)
	}

	res, err := service.UpdateEmployee(id, description, master_date, close_date, db)
	if err != nil {
		log.Println(err)
	}
	log.Println(res)

	http.Redirect(w, r, "/orders", http.StatusMovedPermanently)
}

func EditPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	row := db.QueryRow("SELECT order_id, description, close_date, master_date FROM ORDERS where order_id = $1", id)
	o := data.Orders{}
	err := row.Scan(&o.Order_ID, &o.Description, &o.Master_date, &o.Close_date)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		tmpl, _ := template.ParseFiles("html/edit.html")
		tmpl.Execute(w, o)
	}
}
