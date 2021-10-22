package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"sample.com/database/operations"
)

var template_html = template.Must(template.ParseGlob("templates/*"))

func Home(writer http.ResponseWriter, reader *http.Request) {
	// var template_html *template.Template
	// template_html = template.Must(template.ParseFiles("main.html"))
	customer_list := operations.GetCustomers()
	log.Println(customer_list)
	template_html.ExecuteTemplate(writer, "Home", customer_list)
}

func Create(writer http.ResponseWriter, reader *http.Request) {
	template_html.ExecuteTemplate(writer, "Create", nil)
}
func healthz(writer http.ResponseWriter, reader *http.Request) {
	fmt.Fprintf(writer, "Yo!!")
}

func update(writer http.ResponseWriter, reader *http.Request) {
	customerIdStr := reader.FormValue("id")
	val, _ := strconv.ParseInt(customerIdStr, 10, 64)
	customerID := int(val)
	customer := operations.GetCustomerByID(customerID)
	template_html.ExecuteTemplate(writer, "Update", customer)
}

func InsertCustomer(writer http.ResponseWriter, reader *http.Request) {
	var customer operations.Customer
	customer.CustomerName = reader.FormValue("customername")
	customer.SSN = reader.FormValue("ssn")
	log.Println(customer)
	operations.InsertCustomer(customer)
	customer_list := operations.GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customer_list)
}

func alter(writer http.ResponseWriter, reader *http.Request) {
	var customer operations.Customer
	customer.CustomerName = reader.FormValue("customername")
	customer.SSN = reader.FormValue("ssn")
	log.Println(customer)
	operations.UpdateCustomer(customer)
	customer_list := operations.GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customer_list)
}
func delete(writer http.ResponseWriter, reader *http.Request) {
	customerIdStr := reader.FormValue("id")
	val, _ := strconv.ParseInt(customerIdStr, 10, 64)
	customerID := int(val)
	operations.DeleteCustomerByID(customerID)
	customer_list := operations.GetCustomers()
	template_html.ExecuteTemplate(writer, "Home", customer_list)
}

func view(writer http.ResponseWriter, reader *http.Request) {
	customerIdStr := reader.FormValue("id")
	val, _ := strconv.ParseInt(customerIdStr, 10, 64)
	customerID := int(val)
	customer := operations.GetCustomerByID(customerID)
	var customers []operations.Customer
	customers = append(customers, customer)
	log.Println("Trying to view", customers)
	template_html.ExecuteTemplate(writer, "View", customers)
}

func main() {
	log.Println("Server started on port 8000")
	http.HandleFunc("/", Home)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/insert", InsertCustomer)
	http.HandleFunc("/update", update)
	http.HandleFunc("/alter", alter)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/view", view)
	http.ListenAndServe(":8000", nil)
}
