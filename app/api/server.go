package api

import (
	"app/api/server/app/interpreter"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var PORT int = 3500

//Requests Format
type Request struct {
	WithHeader bool   `json:"withHeader"`
	Seperator  string `json:"seperator"`
	Somme      bool   `json:"somme"`
	Moyenne    bool   `json:"moyenne"`
	Median     bool   `json:"median"`
	MaxValue   bool   `json:"maxValue"`
	EntireFile bool   `json:"entireFile"`
	FromLine   string `json:"fromLine"`
	ToLine     string `json:"toLine"`
}

//Responses Format
type Response struct {
	Code    string //Error or Success
	Status  int    //Http Status Code
	Message string //Reponse Message
}

/*
@returns : boolean (true if file stored false otherwise)
, string (err message if false or the filepath if true)
*/

func uploadFile(w http.ResponseWriter, r *http.Request) {

	//Récupération du ficher taille maximale de 15MB
	r.ParseMultipartForm(5 << 15)
	file, _, err := r.FormFile("file")

	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	defer file.Close()

	//Récupération des paramètres et des operations
	var requestParams = r.Form["params"][0]

	request := Request{}
	json.Unmarshal([]byte(requestParams), &request)

	//Store the file's content
	pwd, _ := os.Getwd()
	filepath := pwd + "/app/data"
	temp, err := ioutil.TempFile(filepath, "fichier-*.csv")

	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	defer temp.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	temp.Write(content)

	var operations []string

	if request.Somme {
		operations = append(operations, "somme")
	}
	if request.Moyenne {
		operations = append(operations, "moyenne")
	}
	if request.Median {
		operations = append(operations, "mediane")
	}
	if request.MaxValue {
		operations = append(operations, "maxValue")
	}

	//Envoyer le fichier pour l'interpretation
	var from, _ = strconv.Atoi(request.FromLine)
	var to, _ = strconv.Atoi(request.ToLine)
	results, header, err := interpreter.Interpreter(temp.Name(), request.WithHeader, request.Seperator, request.EntireFile, from, to, operations)

	//Envoi des resultats
	if err != nil {
		fmt.Fprint(w, err.Error())
	} else {
		var response string = "{\n"

		if request.WithHeader {
			response = fmt.Sprintf(response+"\"En-tête\" : %v ,\n", header)
		}

		if request.Somme {
			response = fmt.Sprintf(response+"\"Somme\" : %v,\n", results["somme"])
		}
		if request.Moyenne {
			response = fmt.Sprintf(response+"\"Moyenne\" : %v,\n", results["moyenne"])
		}
		if request.Median {
			response = fmt.Sprintf(response+"\"Median\" : %v,\n", results["mediane"])
		}
		if request.MaxValue {
			response = fmt.Sprintf(response+"\" Valeur Maximale\" : %v", results["maxValue"])
		}

		response = strings.TrimSuffix(response, ",\n") + "\n}"

		fmt.Fprint(w, response)
	}

}

func launchAPI() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
}

func launchWebApp() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
}

func Server() {
	//Lancement de l'API
	go launchAPI()

	//Lancement de l'interface web
	launchWebApp()
}
