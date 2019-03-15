package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/*
* Entry
 *
 * The Entry struct is the repeating part of the ToDo struct.
 *
 * The newEntry function loads passed data into the struct format.
*/
type Entry struct {
	Name   string
	Email  string
	Mobile string
}

func newEntry(name string, email string, mobile string) Entry {
	return Entry{
		Name:   name,
		Email:  email,
		Mobile: mobile,
	}
}

/*
 * Contacts
 *
 * The Contacts struct is the list parent and contains the Entries.
 */
type Contacts struct {
	User string
	List []Entry
}

func getData() Contacts {
	var contacts Contacts

	contacts.User = "Kenn"
	contacts.List = append(contacts.List, newEntry("Jane Dough", "janedough@example.com", "123-456-7890"))
	contacts.List = append(contacts.List, newEntry("Jon Dough", "jondough@example.com", "234-567-8901"))
	contacts.List = append(contacts.List, newEntry("Sour Dough", "sourdough@example.com", "345-678-9012"))

	return contacts
}

func contactsMain(w http.ResponseWriter, r *http.Request) {
	contacts := getData()

	// Files are provided as a slice of strings.
	paths := []string{
		"templates/contact-main.tmpl",
		"templates/contact-header.tmpl",
		"templates/contact-footer.tmpl",
	}

	contactTmpl, _ := template.ParseFiles(paths...)

	err := contactTmpl.ExecuteTemplate(w, "layout", contacts)
	if err != nil {
		panic(err)
	}
}

/*
 *
 */
func main() {
	r := mux.NewRouter()

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./public/js"))))

	r.HandleFunc("/", contactsMain)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
