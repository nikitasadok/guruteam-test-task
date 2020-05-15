package main

import (
	"net/http"
	"testTaskGuru/globals"
	"testTaskGuru/routes"
	"time"
)


// ?TODO Add cache loading for deposits and transactions
func main() {
	globals.InitApp()
	go triggerEveryTimePeriod(10000 * time.Millisecond, globals.UpdateDatabase)
	router := routes.NewRouter()
	http.ListenAndServe(":8080", router)
	defer globals.Db.Close()

}

func triggerEveryTimePeriod(d time.Duration, f func()) {
	for range time.Tick(d) {
		f()
	}
}
