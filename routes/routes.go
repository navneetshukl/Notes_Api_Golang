package routes

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/navneetshukl/Golang_notes_API/database"
	"github.com/navneetshukl/Golang_notes_API/models"
)
type Session_Id struct{
	session string
	email string
}
var sessionId map[int]Session_Id


func generateSessionID() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b), nil
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var user models.Signup
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "400 Bad Request")
		return
	}
	fmt.Println(user.Name, user.Email, user.Password)
	name := user.Name
	email := user.Email
	password := user.Password
	database.InsertIntoUser(name, email, password)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "200 OK")

}

func Login(w http.ResponseWriter, r *http.Request) {

	var user models.Login
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "400 Bad Request")
		return
	}
	email := user.Email
	password := user.Password
	fmt.Println(email , password)
	userExists, err := database.CheckUser(email, password)
	if err != nil {
		log.Fatalf("Error checking user: %v", err)
	}
	if !userExists {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "401 Unauthorized")
		return

	}
	session, err := generateSessionID()
	if err != nil {
		log.Fatal("Session id not generated")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "200 Ok")
	sessionId=make(map[int]Session_Id)
	sess:=Session_Id{
		session:session,
		email:email,
	}
	sessionId[0]=sess
	fmt.Println(session)
	fmt.Println("Login Successfully")

}

func GetNote(w http.ResponseWriter, r *http.Request) {
	cnt := 0
	notes, err := database.GetDataFromNotes()
	if err!=nil{
		log.Fatalf("No data found")
		return
	}
	length := len(notes)
	mail:=sessionId[0].email
	if  mail==""{
		w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "401 Unauthorized")
			return

	}

	for i := 0; i < length; i++ {
		if notes[i].Email == mail {
			cnt += 1
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(models.StudentNotes[i].Title)
			fmt.Println(notes[i].Title)
		}

	}

	if cnt == 0 {
		fmt.Println("No User Found")
		return
	}

}

func CreateNote(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	req := string(body)
	fmt.Println(req)

	var note models.Notes
	json.Unmarshal(body, &note)
	email:=sessionId[0].email
	title:=note.Title
	if  email==""{
		w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "401 Unauthorized")
			return

	}

	err=database.SaveData(title,email)
	if err!=nil{
		log.Fatalf("Data is not saved")
	}
	// Send a 201 created response

	w.Header().Add("COntent-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Added Successfully")

}

func DeleteNote(w http.ResponseWriter, r *http.Request) {

	email:=sessionId[0].email
	if  email==""{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "401 Unauthorized")
			return

	}

	err:=database.DeleteData(email)
	if err!=nil{
		log.Fatalf("Data is not deleted")
		return;
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "200 Ok")
	fmt.Println("Data Deleted Successfully")

}
