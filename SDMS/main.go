package main

// create an api to insert student data into database
// api ==> payload - name,roll_no(primary key)

// read student data  from database
// api (parameter roll_no)

// update student data from database
// api (parameter roll_no)

//delete student data from database
// api (parameter roll_no)
import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "feast"
)

type RequestBody struct {
	RollNo string `json:"roll_no"`
	Name string `json:"name"`
}


func handlePostRequest(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Println("Some err occured")
	}
	fmt.Println("printing body of req")
	fmt.Println(body)


	var requestBody RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil{
		fmt.Println("Some err occured")
	}
	recievedData := fmt.Sprintf("Recieved message, %s", requestBody)
	fmt.Println(recievedData)
	fmt.Println(requestBody.RollNo)
	
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)


	db, err := sql.Open("postgres", connStr)
	if err != nil{
		log.Fatal("error!!!")
	}
	fmt.Println(db)
	defer db.Close()
	err = db.Ping()
	if err != nil{
		log.Fatal("ERROR IN DB")
	}
	fmt.Println("Connected to the database")
	_,err = db.Exec(`CREATE TABLE IF NOT EXISTS STUDENTS (roll_no VARCHAR(50) PRIMARY KEY, name VARCHAR(50))`)
	if err != nil{
		log.Fatal("table creation students failed")

	}
	fmt.Println(" student table created")
	insertStmt := "INSERT INTO STUDENTS (roll_no, name) VALUES($1, $2)"
	_, err = db.Exec(insertStmt, requestBody.RollNo, requestBody.Name)
	
	if err != nil{
		fmt.Println(err)
		fmt.Println("data with given roll no already exists")
	}
	fmt.Println(" data inserted to table")


}

func handleReadRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Println("inside read req")
    
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(body)

    var requestBody RequestBody
    err = json.Unmarshal(body, &requestBody)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("printing data")
    fmt.Println(requestBody.RollNo)

    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("error!!!")
    }
    fmt.Println(db)
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal("ERROR IN DB")
    }
    fmt.Println("Connected to the database")

    readStmt := "SELECT * FROM STUDENTS WHERE roll_no = $1"
    rows, err := db.Query(readStmt, requestBody.RollNo)
    if err != nil {
        fmt.Println(err)
        fmt.Println("Error querying the database")
        http.Error(w, "Error querying the database", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    // Process the query results
    for rows.Next() {

		fmt.Println("inside rows")
        var rollNo string
        var name string

        if err := rows.Scan(&rollNo, &name); err != nil {
            fmt.Println(err)
            fmt.Println("Error scanning row")
            http.Error(w, "Error scanning row", http.StatusInternalServerError)
            return
        }

        // Write the response directly to http.ResponseWriter
        fmt.Fprintf(w, "Roll No: %s, Name: %s\n", rollNo, name)
    }
	panic("this is fatal")
	fmt.Println("done here")
    // Check for errors from iterating over rows (e.g., scanning errors)
    if err := rows.Err(); err != nil {
        fmt.Println(err)
        fmt.Println("Error iterating over rows")
        http.Error(w, "Error iterating over rows", http.StatusInternalServerError)
    }
}

func main(){

	fmt.Println("SDMS creation started")
	http.HandleFunc("/api/posthello", handlePostRequest)

	http.HandleFunc("/api/readhello", handleReadRequest)


	fmt.Println("server listening on :8080")
	http.ListenAndServe(":8080", nil)
}