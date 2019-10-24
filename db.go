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
	fmt.Println("------------------------")
}


/*func update() {
	sqlStatement := `UPDATE profile_schema.profile SET  kyc_level = $2 WHERE consumer_id = $1 RETURNING consumer_id ;`
    	
	var email string
	var id int
	
	err = db.QueryRow(sqlStatement, 101, 2).Scan(&consumer_id, &kyc_level)
	
	checkErr(err)
	
	fmt.Println(consumer_id, kyc_level)

}*/	

//fetching data from the database
/*func fetch() {
	db, err := sql.Open("postgres", psqlInfo)
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
} */


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
	fmt.Println("consumer_id | gender | kyc_level | sign_up_date ")
    
	for rows.Next() {
		var consumer_id, kyc_level int
        var gender,sign_up_date *string	
        err = rows.Scan(&consumer_id, &gender, &kyc_level, &sign_up_date)
        checkErr(err)
        fmt.Printf("%11v | %6v | %9v | %12v\n", consumer_id, gender, kyc_level, sign_up_date)
	}
	
	//inserting values

	var lastInsertId int
    err = db.QueryRow("INSERT INTO profile_schema.profile(consumer_id, gender, kyc_level ) VALUES ($1,$2,$3);",1,"female",2).Scan(&lastInsertId)

	//updating values
	sqlStatement := `UPDATE profile_schema.profile SET kyc_level = $2 WHERE consumer_id = $1;`
	
	_, err = db.Exec(sqlStatement, 100, 2)
	
	if err != nil {
		panic(err)
	}




	defer db.Close()
}
