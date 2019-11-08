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
	sign_up_date time.Time

}

type coupon struct {

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
	println("")
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

//function to check errors 
func checkErr(err error) {
	if err != nil {
		panic(err)	
   }
}
	

//dropping table
func drop_table() {
	sqlStatement := "DROP TABLE coupon"
	_, err = db.Exec(sqlStatement)
	checkErr(err)
	println("----table deleted----")
}

//creating new table in database
func create_table() {
	sqlStatement := "CREATE TABLE coupon()"
	_, err = db.Exec(sqlStatement)
	checkErr(err)
	println("----table created----")
}

//function to insert values
func insert() {
	
	Initial_Sql_Statement := "INSERT INTO profile_schema.profile() VALUES ($1,$2,$3)"

	col := "consumer_id,gender,kyc_level"
	
	Final_Sql_Statement := Initial_Sql_Statement[:35] + col + Initial_Sql_Statement[35:]
	
	_, err = db.Exec(Final_Sql_Statement,110,"female",1)
	checkErr(err)
	
	fmt.Println(Final_Sql_Statement)
	println("----inserted values----")
}

//function to update values
func update() {

	sqlStatement := "UPDATE profile_schema.profile SET gender = $1 WHERE consumer_id = $2;"
	
	_, err = db.Exec(sqlStatement, "female", 5)
	checkErr(err)	
	println("-------updated values--------")
} 

//function to add column
func alter_add_column() {
	sqlStatement := `ALTER TABLE profile_schema.profile ADD COLUMN hipbar int ;`
	_, err = db.Exec(sqlStatement)
	checkErr(err)
	println("----column added----")
}

//function to rename column
func alter_rename_column() {
	sqlStatement := `ALTER TABLE profile_schema.profile RENAME COLUMN hipbar TO hipster_bar ;`
	_, err = db.Exec(sqlStatement)
	checkErr(err)
	println("----column renamed---")
}

//function to drop a column
func alter_drop_column() {
	sqlStatement := `ALTER TABLE profile_schema.profile DROP COLUMN name`
	_, err = db.Exec(sqlStatement)
	checkErr(err)
	println("----column deleted-----")
}

//function to update and alter column simultaneously
func alter_and_update_column() {
	sqlStatement1 := `ALTER TABLE profile_schema.profile ADD column name varchar(100)` 
	_, err = db.Exec(sqlStatement1)
	checkErr(err)	
	
	sqlStatement2 := `UPDATE profile_schema.profile SET name = $1 WHERE consumer_id = $2`
	_, err = db.Exec(sqlStatement2,"neeraj",100)
	checkErr(err)
	println("-----Altered & Updated simultaneously-----")
}

func main() {
	
	insert()
	
	fetch()	
	
	defer db.Close()

}

