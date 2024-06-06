package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ServiceWeaver/weaver"
	"github.com/pelletier/go-toml"
)

type Config struct {
	ServiceWeaver struct {
		Binary string `toml:"binary"`
	}
	Multi struct {
		Listeners struct {
			Hello struct {
				Address string `toml:"address"`
			} `toml:"hello"`
		} `toml:"listeners"`
	} `toml:"multi"`
}

type app struct {
	weaver.Implements[weaver.Main]
	reverser  weaver.Ref[Reverser]
	addspacer weaver.Ref[AddSpacer]
	hello     weaver.Listener `weaver:"hello"`
	hii       weaver.Listener `weaver:"hii"`
}

func main() {
	// Load configuration from TOML file
	config := loadConfig("weaver.toml")

	// Get the listener address from the configuration
	listenerAddr := config.Multi.Listeners.Hello.Address

	if err := weaver.Run(context.Background(), func(ctx context.Context, app *app) error {
		// Serve the /hello endpoint.
		http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
			name := r.URL.Query().Get("name")
			if name == "" {
				name = "World"
			}
			reversed, err := app.reverser.Get().Reverse(ctx, name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "Hello, %s!\n", reversed)
		})

		// Serve the /hii endpoint.
		http.HandleFunc("/hii", func(w http.ResponseWriter, r *http.Request) {
			name := r.URL.Query().Get("name")
			if name == "" {
				name = "World"
			}
			reversed, err := app.addspacer.Get().AddSpace(ctx, name)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "Hello, %s!\n", reversed)
		})

		return http.ListenAndServe(listenerAddr, nil)
	}); err != nil {
		log.Fatal(err)
	}
}

func loadConfig(filename string) Config {
	var config Config
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening config file: %v", err)
	}
	defer file.Close()

	if err := toml.NewDecoder(file).Decode(&config); err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}

		return config
	}


	
// package main

// import "log"

// type Student struct {
// 	Name   string
// 	rollno string
// }

// func main() {
// var student Student
// student.Name = "Guru"
// student.rollno = "123"
// log.Println(student)
// changename(&student)
// log.Println(student)
// changenamenormal(student)
// log.Println(student)
// }


// func changename(student *Student){
// 	student.Name = "Yameen"
// }

// func changenamenormal(student Student){
// 	student.Name = "Guru"
// }