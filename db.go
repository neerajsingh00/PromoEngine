package main

import (
  "database/sql"
  "fmt"

  _ "github.com/lib/pq"
)

//const --
const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "Neeraj@123"
  dbname   = "demo"
)


var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
var db, err = sql.Open("postgres", psqlInfo)


//function to check errors
func checkErr(err error) {
	if err != nil {
		panic(err)
   }
}
	
//establishing connection with the PSQL server
func init() {
	checkErr(err)
  	err = db.Ping()
  	checkErr(err)

  	fmt.Println("Connection Established!")
	fmt.Println("------------------------")
}


func fetch() {
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


}

func insert() {

	var lastInsertId int
    err = db.QueryRow("INSERT INTO profile_schema.profile(consumer_id, gender, kyc_level ) VALUES ($1,$2,$3);",5,"f",2).Scan(&lastInsertId)

}

func update() {

	sqlStatement := `UPDATE profile_schema.profile SET kyc_level = $2 WHERE consumer_id = $1;`
	
	_, err = db.Exec(sqlStatement, 5, 1)
	checkErr(err)	
		

} 


func main() {

	fetch()
	fmt.Println("----------------")
	insert()
	fetch()
	fmt.Println("-----------------")
	update()
	fetch()
	fmt.Println("--------------------")
	defer db.Close()

}
