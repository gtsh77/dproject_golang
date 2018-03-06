package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

// The person Type (more like an object)
type Msg struct {
    ID  string   `json:"id,omitempty"`
    Msg string   `json:"msg,omitempty"`
}

var msg []Msg

// Display all from the msg var
func GetMessages(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(msg)
}

// Display a single data
func GetMessage(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range msg {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Msg{})
}

//create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var msg2 Msg
    _ = json.NewDecoder(r.Body).Decode(&msg)
    msg2.ID = params["id"]
    msg = append(msg, msg2)
    json.NewEncoder(w).Encode(msg)
}

// main function to boot up everything
func main() {
    router := mux.NewRouter()
    msg = append(msg, Msg{ID: "1", Msg: "PONG_FROM_GOLANG"})
    msg = append(msg, Msg{ID: "2", Msg: "SECOND_PONG_FROM_GOLANG"})
    router.HandleFunc("/api/msg/", GetMessages).Methods("GET")
    router.HandleFunc("/api/msg/{id}/", GetMessage).Methods("GET")
    router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("frontend/dist"))))
    log.Fatal(http.ListenAndServe(":8080", router))
}