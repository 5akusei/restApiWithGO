package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil" // sirve para manejar las entradas y salidas
	"log"
	"net/http" //para crear el servidor http
	"strconv"  // para convertir

	"github.com/gorilla/mux" //importacion del modulo mux, sirve para definir rutas
)

// Definicion de la tarea
type task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

// Definicion del arreglo para una lista de tareas
type allTasks []task

// esta es la variable que estaremos modificando
var tasks = allTasks{
	{
		ID:      1,
		Name:    "Tarea uno",
		Content: "Descripcion uno",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Esta es la ruta root!")
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Inserta una tarea con datos validos")
	}

	json.Unmarshal(reqBody, &newTask) //asignacion de la info que me llega y se añade a la nueva tarea
	newTask.ID = len(tasks) + 1       // creacion del id
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json") //dice el tipo de contenido que va a devolver a la peticion
	w.WriteHeader(http.StatusCreated)                  //devuelve el status de la accion
	json.NewEncoder(w).Encode(newTask)

}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func getOneTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}

	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask task

	if err != nil {
		fmt.Fprintf(w, "ID invalido")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte datos validos")
	}
	json.Unmarshal(reqBody, &updatedTask)

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)

			updatedTask.ID = t.ID
			tasks = append(tasks, updatedTask)

			// w.Header().Set("Content-Type", "application/json")
			// json.NewEncoder(w).Encode(updatedTask)
			fmt.Fprintf(w, "La tarea con el ID %v se ha actalizado", taskID)
		}
	}

}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "ID invalido")
		return
	}

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "La tarea con el ID %v se ha eliminado", taskID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true) //enrutador

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getOneTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	//para crear el servidor, recibe el puerto donde va a escuchar y el enrutador
	log.Fatal(http.ListenAndServe(":3000", router))
}
