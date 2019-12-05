package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Note struct {
    Content     string `json:"content"`
}

var notes []Note

func getNoteHandler(w http.ResponseWriter, r *http.Request) {
   //Convert the "notes" variable to json
   noteListBytes, err := json.Marshal(notes)

   // If there is an error, print it to the console, and return a server
   // error response to the user
   if err != nil {
       fmt.Println(fmt.Errorf("Error: %v", err))
       w.WriteHeader(http.StatusInternalServerError)
       return
   }
   // If all goes well, write the JSON list of birds to the response
   w.Write(noteListBytes)
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
   // Create a new instance of Note
   note := Note{}

   // We send all our data as HTML form data
   // the `ParseForm` method of the request, parses the
   // form values
   err := r.ParseForm()

   // In case of any error, we respond with an error to the user
   if err != nil {
       fmt.Println(fmt.Errorf("Error: %v", err))
       w.WriteHeader(http.StatusInternalServerError)
       return
   }

   // Get the information about the note from the form info
   note.Content = r.Form.Get("content")

   fmt.Println("note: %v", note)

   // Append our existing list of birds with a new entry
   notes = append(notes, note)

   //Finally, we redirect the user to the original HTMl page (located at `/assets/`)
   http.Redirect(w, r, "/assets/", http.StatusFound)
}