package main

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "Neeraj@123"
  dbname   = "demo"
)


//function to check errors
func checkErr(err error) {
	if err != nil {
		panic(err)
   }
}

//establishing connection with the PSQL server
func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)
  	err = db.Ping()
  	checkErr(err)

  	fmt.Println("Connection Established!")
} 

//fetching data from the database
/*func fetch() {
	rows, err := db.Query("SELECT * FROM profile_schema.profile")   
    checkErr(err)
	for rows.Next() {
		var consumer_id, kyc_level int
        var gender,sign_up_date *string	
        err = rows.Scan(&consumer_id, &kyc_level, &gender, &sign_up_date)
        checkErr(err)
        fmt.Println("consumer_id | kyc_level | gender | sign_up_date ")
        fmt.Printf("%11v | %9v | %6v | %12v\n", consumer_id, kyc_level, gender, sign_up_date)
	}
} 


/* inserting data from the database
func insert() {
	sqlStatement := `INSERT INTO profile_schema.profile (CONSUMER_ID, GENDER, KYC_LEVEL, SIGN_UP_DATE) VALUES ($1, $2, $3, $4) RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, 2122, "MALE", "Neeraj", "timestamp_format").Scan(&id)
	if err != nil {
    	panic(err)
  }
  	fmt.Println("New record ID is:", id)
} */ 


// updating data in the database
//func update() {


//}



func main() {
	//establishing connection
	fmt.Println("")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)
  	err = db.Ping()
  	checkErr(err)

	//querying data
	fmt.Println("------------querying----------")
	rows, err := db.Query("SELECT * FROM profile_schema.profile")   
    checkErr(err)
	for rows.Next() {
		var consumer_id, kyc_level int
        var gender,sign_up_date *string	
        err = rows.Scan(&consumer_id, &gender, &kyc_level, &sign_up_date)
        checkErr(err)
        fmt.Println("consumer_id | gender | kyc_level | sign_up_date ")
        fmt.Printf("%11v | %6v | %9v | %12v\n", consumer_id, gender, kyc_level, sign_up_date)
	}
	
	//inserting values



	defer db.Close()
}
