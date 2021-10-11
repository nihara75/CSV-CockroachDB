package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {

	csvFile, err := os.Open("table.csv")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(csvLines[1][0])

	err1 := godotenv.Load(".env")

	if err1 != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err2 := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if _, err := db.Exec("TRUNCATE TABLE student;"); err != nil {
		log.Fatal(err)
	}

	if err2 != nil {
		log.Fatal("error connecting to the database: ", err2)
	}
	fmt.Println("Database student opened and ready.")
	//defer db.Close()

	sqlStatement := `insert into inputTable (id,category,prod_name,description,mrp) values ($1,$2,$3,$4,$5)`
	for i := 1; i < len(csvLines); i++ {
		id, er := strconv.Atoi(csvLines[i][0])
		if er != nil {
			log.Fatal(er)
		}

		cate := csvLines[i][1]
		prod := csvLines[i][2]
		desc := csvLines[i][3]
		mr := csvLines[i][4]

		_, err3 := db.Exec(sqlStatement, id, cate, prod, desc, mr)
		if err3 != nil {
			log.Fatalf("error")
		}
	}

}
