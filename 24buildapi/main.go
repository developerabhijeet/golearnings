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

// model for courses -file
type Course struct {
	CourseId    string  `json: "courseid"`
	CourseName  string  `json: "coursename"`
	CoursePrice int     `json: "price"`
	Author      *Author `json: "author"`
}
type Author struct {
	FullName string `json: "fullname"`
	Website  string `json: "website"`
}

// fake DB
var courses []Course

// middleware or helper
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.Coursename=""
	return c.CourseName == ""
}

func main() {
	fmt.Println("APIs for courses")

	r := mux.NewRouter()

	// seeding
	courses = append(courses, 
		Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, 
		Author: &Author{FullName: "Abhijeet", Website: "lco.dev"}})
	courses = append(courses, 
		Course{CourseId: "3", CourseName: "MERN", CoursePrice: 299, 
		Author: &Author{FullName: "Abhijeet", Website: "lco.dev"}})

	// routing
	r.HandleFunc("/", server).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")



	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))

}

// Controllers - file

//  serve home route

func server(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Server is up 🚀</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	// loop through the courses, find the matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")
	
	// if data body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data in body")
	}
	
	
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	for _, exist := range courses {
		if exist.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course already exist with this name")
			return
		}
	}
	
	

	rand.NewSource(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from req

	params := mux.Vars(r)

	// loop , id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the given id")
	return

}
func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("Delete successfully")
	return

}
