package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}

var courses []Course //DB
func (c *Course) IsEmpty() bool { //middle ware
	return c.CourseName == ""
}

func main() {
	fmt.Println("Hey Server Working!!")
	r := mux.NewRouter()
	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Name: "AJAY", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Name: "AJAY", Website: "go.dev"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":4000" , r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the API"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	fmt.Println("Params : " , params)
	fmt.Printf("Params Type %T : " , params)

	for _ , course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course with this ID found")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please Send some data")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Values not appropriate")
	}
	
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses , course)
	json.NewEncoder(w).Encode(courses)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	id := r.URL.Query().Get("id")
	var updatedCourse Course
	_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
	for index , course := range courses {
		if course.CourseId == id {
			courses = append(courses[:index] , courses[index+1:]...)    //delete the previous entry
			updatedCourse.CourseId = course.CourseId
			courses = append(courses , updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}

	json.NewEncoder(w).Encode("Course with this ID is not found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index , course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index] , courses[index+1:]...)    
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course with this ID is not found")
}