package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/resend/resend-go/v2"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	mux := http.NewServeMux()
	mux.Handle("/", &reset{})
	http.ListenAndServe(":8080", mux)
}

type reset struct {}

func (h *reset) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if(r.URL.RawQuery != "auth="+os.Getenv("admin_password")){
		w.WriteHeader(403)
	}

	connStr := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("DB ERROR: ", err.Error())
		informErrorViaMail(err.Error())
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("DB PING ERROR: ", err.Error())
		informErrorViaMail(err.Error())
	}

	query := `UPDATE userdetails SET hits = 0 ; UPDATE userkeystatus SET status = 'ok'; `

	_, dbErr := db.Exec(query)

	if dbErr != nil {
		informErrorViaMail(dbErr.Error())
	}

	w.WriteHeader(200);
}

func informErrorViaMail(mes string) {
	apiKey := os.Getenv("RESEND_APIKEY")
	client := resend.NewClient(apiKey)

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{"yakshitchhipa@gmail.com"},
		Subject: "☠️😬 ERROR OCCURED: Daily Database Reset Lambda did'nt worked 😨😱",
		Html:    fmt.Sprintf("<h1> ERROR OCCURED IN RESETING THE DATA BASE: %s </h1>", mes),
	}

	sent, err := client.Emails.Send(params)

	if err != nil {
		fmt.Println("Error in sending the mail:", err.Error())
	}

	fmt.Println("MAIL SENT TO NOTIFY ERROR", sent.Id)
}
