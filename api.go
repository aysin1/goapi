package main

import (
    "fmt"
    "log"
    "net/http"
    "html/template"
    "regexp"
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)
type userdata struct {
	uname string
	email string
	pwd string
	cnf string

}
var client = options.Client().ApplyURI("mongodb://localhost:27017")
	var conn_ = mongo.connect(context.TODO().client)
	var collection = conn.Database("databasename").collection("collectionname")


func main() {
    http.HandleFunc("/", home)
    http.HandleFunc("/login", login)
    http.HandleFunc("/loginval", loginval)
    http.HandleFunc("signup", signup)
    http.HandleFunc("/signupval", signupval)
    fmt.Println("chalu h..............")
    log.Fatal(http.ListenAndServe(":8080", nil))
    //http.ListenAndServe(":8080", nil)
}


func home(w http.ResponseWriter, r *http.Request) {
    tmpl,err := template.parsefile("./home.html")
    if err != nil {
    	log.Println(err)
    }
    err=tmpl.Execute(w,nil)
    if err != nil {
    	log.Println(err)
    }
}

func login(w http.ResponseWriter, r *http.Request) {
    tmpl,err := template.parsefile("./login.html")
    if err != nil {
    	log.Println(err)
    }
    err=tmpl.Execute(w,nil)
    if err != nil {
    	log.Println(err)
    }
}

func signup(w http.ResponseWriter, r *http.Request) {
    tmpl,err := template.parsefile("./signup.html")
    if err != nil {
    	log.Println(err)
    }
    err=tmpl.Execute(w,nil)
    if err != nil {
    	log.Println(err)
    }
}

func loginval(w http.ResponseWriter, r *http.Request) {
	email = r.formvalue("email")
	pwd = r.formvalue("pwd")

	emailcheck := helper.IsEmpty(email)
	pwdcheck := helper.IsEmpty(pwd)

    if emailcheck || pwdcheck {
    	fmt.Fprint(w,"Error code -10: Please Enter the Data")
    }

    // dbPwd := "1234!*"
    // dbEmail := cihan.ozhan@gmail.com
    collection, err := db.GetDBCollection()
    err = collection.FindOne(context.TODO(), bson.D{{"email", user.email},{"pwd", user.pwd}}).Decode(&result)
   
    dbEmail=result.email
    dbPwd=result.pwd
    
    if email==dbEmail && pwd==dbPwd {
    	fmt.Fprintln(w,"login sucessfull!!!!!!!!!!!!!!!")
    } else {
    	fmt.Fprintln(w,"login Failed!!!!!!!!!!!!!!!")
    }
    
}
func signupval(w http.ResponseWriter, r *http.Request) {
    uname = r.formvalue("username")
    email = r.formvalue("email")
    pwd = r.formvalue("pwd")
    cnf = r.formvalue("confirmpwd")
    
    if(uname!="" && email!="" && pwd!="" && cnf!="") {
        if(pwd==cnf) {
            user := userdata(uname,pwd,cnf,email)
            id,_ := collection.insertone(context.TODO().user)
            fmt.Println("User Id",Id)
            http.Redirect(w,r,"/login.html",302)
            fmt.Println(:sucessfully Registered)
        }
    }
    http.Redirect(w,r,"/signup.html",302)
}