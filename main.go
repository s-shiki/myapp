package main

import {
  "net/http"
}

func main() {
  // handle static assets
  mux := http.NewServerMux()
  files := http.FileServr(http.Dir(config.Static))
  mux.Handle("/static/", http.StripPrefix("/static/", files))

  // all route patterns matched here
  // route handler functions defined in other files

  //index
  mux.HandleFunc("/", index)
  //error
  mux.HandleFunc("/err", err)

  // defined in route_auth.go
  mux.HandleFunc("/login", login)
  mux.HandleFunc("/logout", logout)
  mux.HandleFunc("/signup", signup)
  mux.HandleFunc("/signup_account", signupAccount)
  mux.HandleFunc("/authenticate", authenticate)

  // difined in route_thread.go
  mux.HandleFunc("/thread/new", newThread)
  mux.HandleFunc("/thread/create", createThread)
  mux.HandleFunc("/thread/post", postThread)
  mux.HandleFunc("/thread/read", readThread)

  // starting up the server
  server := &http.Server{
    Addr:     "0.0.0.0:8080",
    Handler:  mux,
  }
  server.ListenAndServe()
}
