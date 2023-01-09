package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * Complete the 'postHandler', 'deleteHandler' and 'getHandler' functions below.
 *
 * All functions are expected to be void.
 * All functions accept http.ResponseWriter w and *http.Request req as parameters.
 */
func postHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("> POST handler")

	body, err := io.ReadAll(req.Body)
	checkError(err)

	var l Lake
	err = json.Unmarshal(body, &l)
	checkError(err)

	// Empty body
	if (l == Lake{}) {
		// w.WriteHeader(http.StatusNoContent)
		return
	}

	// TODO: protect with a Mutex
	store[l.Id] = l

	log.Println("  - STORE: ", store)
}

func deleteHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("> DELETE handler")

	id := req.URL.Query().Get("id")

	if (store[id] == Lake{}) {
		http.Error(w, "Lake ID not found: "+id, http.StatusNotFound)
		return
	}

	delete(store, id)
	log.Println("  - STORE: ", store)
}

func getHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("> GET handler")

	id := req.URL.Query().Get("id")

	if (store[id] == Lake{}) {
		// w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Lake ID not found: "+id, http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(store[id])
	checkError(err)
	w.Write(resp)

	log.Println("  - STORE: ", store)
}

func main() {
	// reader  := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
	// --- Custom testing
	file, err := os.Open("./tests/01-input")
	if err != nil {
		panic("File access error!")
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	// --- export desired output target
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	actionsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	log.Println("> Incoming requests: ", actionsCount)
	checkError(err)

	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/delete", deleteHandler)

	go http.ListenAndServe(portSuffix, nil)
	time.Sleep(100 * time.Millisecond)

	var actions []string

	for i := 0; i < int(actionsCount); i++ {
		actionsItem := readLine(reader)
		actions = append(actions, actionsItem)
	}

	for _, actionStr := range actions {
		var action Action
		err := json.Unmarshal([]byte(actionStr), &action)
		checkError(err)
		switch action.Type {

		case "post":
			resp, err := http.Post(Address+"/post", "application/json", strings.NewReader(action.Payload))
			getRespHeaders(resp)
			checkError(err)

		case "delete":
			client := &http.Client{}
			req, err := http.NewRequest("DELETE", Address+"/delete?id="+action.Payload, nil)
			checkError(err)

			resp, err := client.Do(req)
			getRespHeaders(resp)
			checkError(err)

			if resp.StatusCode != 200 {
				fmt.Fprintf(writer, "%s\n", resp.Status)
				continue
			}

		case "get":
			resp, err := http.Get(Address + "/get?id=" + action.Payload)
			getRespHeaders(resp)
			checkError(err)
			if resp.StatusCode != 200 {
				fmt.Fprintf(writer, "%s\n", resp.Status)
				continue
			}
			var lake Lake
			err = json.NewDecoder(resp.Body).Decode(&lake)
			checkError(err)
			fmt.Fprintf(writer, "%s\n", lake.Name)
			fmt.Fprintf(writer, "%d\n", lake.Area)
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

const portSuffix = ":3333"

var Address = "http://localhost" + portSuffix

type Lake struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Area int32  `json:"area"`
}

type Action struct {
	Type    string
	Payload string
}

var store = map[string]Lake{}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getRespHeaders(resp *http.Response) {
	fmt.Printf(">> Server response headers: %v\n   Status Code: %v\n", resp.Header, resp.StatusCode)
}
