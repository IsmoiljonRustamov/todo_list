package main

import (
	"database/sql"
	"fmt"
	"log"
	_"time"

	_ "github.com/lib/pq"
)

func main() {

	var (
		PostgresUser     = "postgres"
		PostrgesPassword = "12345"
		PostgresHost     = "localhost"
		PostgresPort     = 5432
		PostgresDatabase = "todo"
	)
	connStr := fmt.Sprintf("user = %s password = %s host = %s port = %d dbname = %s sslmode = disable", PostgresUser, PostrgesPassword, PostgresHost, PostgresPort, PostgresDatabase)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("falied to open connected: %v", err)
	}
	DBManager := NewDBManager(db)

	// td, err := DBManager.CreateToDo(TODO{
	// 	title:        "My job",
	// 	descriptions: "Middle",
	// 	assignee:     "Zohid Saidov",
	// 	status:       true,
	// 	deadline:     time.Date(2020, 01, 01, 16, 02, 20, 01, time.Local),
	// })

	// if err != nil {
	// 	log.Fatalf("Failed to CreateToDo; %v", err)
	// }
	// fmt.Println(td)

	// upget,err := DBManager.Update(&TODO{
	// 	id:4,
	// 	created_at: time.Now(),
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to Update: %v", err)
	// }
	// fmt.Println(upget)

	// DBManager.Delete(&TODO{
	// 	id: 4,
	// 	deleted_at: time.Now(),
	// })

		getall, err := DBManager.GetAll(&GetAllParam{
		 	limit: 2,
		 	page: 1,
		 	title: "My job",
		 })
		 if err != nil {
		 	log.Fatalf("Failed to GetAll: %v", err)
		}
		 for a, v := range getall {
		 	fmt.Println("key : ", a,"value: ",v )
		}
		
	
	
}
