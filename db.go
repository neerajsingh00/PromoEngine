package main

import (
  "database/sql"
  "fmt"
  "time"
//  "encoding/json"
  _ "github.com/lib/pq"
)


type profile struct {
	consumer_id int
	gender string
	kyc_level int 
	sign_up_date *time.Time 		`json: "sign_up_date"`
	Custom_attributes []interface{} `json: "custom_attributes"`
}

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "Neeraj@123"
  dbname   = "demo"
)


var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
var db, err = sql.Open("postgres", psqlInfo)
var pro = make([]*profile, 0)

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

//function to query 
func fetch() {
	fmt.Println("------------querying----------")
	rows, err := db.Query("SELECT * FROM profile_schema.profile")   
    checkErr(err)
	fmt.Println("consumer_id | gender | kyc_level | sign_up_date ")
    
	for rows.Next() {
		p := new(profile)	
        err = rows.Scan(&p.consumer_id, &p.gender, &p.kyc_level, &p.sign_up_date)
        checkErr(err)
		pro = append(pro,p)
    }
	
	if err = rows.Err(); err != nil {
		panic(err)
	}
	
	for _, p := range pro {
		fmt.Println (p.consumer_id,"\t |",p.gender,"\t |", p.kyc_level,"\t |", p.sign_up_date)
	}
	


	defer rows.Close()
}

//function to insert values
func insert() {

	var lastInsertId int
	err = db.QueryRow("INSERT INTO profile_schema.profile(consumer_id, gender, kyc_level, hipbar ) VALUES ($1,$2,$3,$4);",5,"f",2,"beer").Scan(&lastInsertId)
	err = db.QueryRow("INSERT INTO profile_schema.profile(consumer_id,gender,kyc_level,hipbar) VALUES ($1,$2,$3,$4);",6,"f",0,"beer").Scan(&lastInsertId)
}

//function to update values
func update() {

	sqlStatement := `UPDATE profile_schema.profile SET gender = $1 WHERE consumer_id = $2;`
	
	_, err = db.Exec(sqlStatement, "female", 5)
	checkErr(err)	
	println("-------updated--------")
} 

//function to add column
func alter_add_column() {
	sqlStatement := `ALTER TABLE profile_schema.profile ADD COLUMN hipbar int ;`
	_, err = db.Exec(sqlStatement)
	checkErr(err)
}

//function to rename column
func alter_rename_column() {
	sqlStatement := `ALTER TABLE profile_schema.profile RENAME COLUMN hipbar TO hipster_bar ;`
	_, err = db.Exec(sqlStatement)
	checkErr(err)
}

//function to drop a column
func alter_drop_column() {
	sqlStatement := `ALTER TABLE profile_schema.profile DROP COLUMN hipster_bar ;`
	_, err = db.Exec(sqlStatement)
	checkErr(err)

}

func main() {


	update()
	fetch()
	fmt.Println("-------")
	defer db.Close()

}

