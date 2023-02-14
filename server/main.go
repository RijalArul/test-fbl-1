package main

import (
	databases "test-fbl-1/server/db"
	"test-fbl-1/server/routers"
)

// import (
// 	"encoding/csv"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// )

// type M map[string]interface{}

// func handleListFiles(w http.ResponseWriter, r *http.Request) {
// 	files := []M{}
// 	basePath, _ := os.Getwd()
// 	filesLocation := filepath.Join(basePath, "files")

// 	err := filepath.Walk(filesLocation, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		if info.IsDir() {
// 			return nil
// 		}

// 		files = append(files, M{"filename": info.Name(), "path": path})
// 		return nil
// 	})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	res, err := json.Marshal(files)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(res)
// }

// func handleDownload(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	path := r.FormValue("path")
// 	fmt.Println(path)
// 	f, err := os.Open(path)
// 	if f != nil {
// 		defer f.Close()
// 	}
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
// 	w.Header().Set("Content-Disposition", contentDisposition)

// 	if _, err := io.Copy(w, f); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// type Employee struct {
// 	ID  string
// 	Age int
// }

// func main() {
// 	records := []Employee{
// 		{"E01", 25},
// 		{"E02", 26},
// 		{"E03", 23},
// 		{"E04", 26},
// 		{"E05", 27},
// 	}
// 	file, err := os.Create("records.csv")
// 	defer file.Close()
// 	if err != nil {
// 		log.Fatalln("failed to open file", err)
// 	}
// 	w := csv.NewWriter(file)
// 	defer w.Flush()
// 	var data [][]string
// 	for _, record := range records {
// 		row := []string{record.ID, strconv.Itoa(record.Age)}
// 		data = append(data, row)
// 	}
// 	w.WriteAll(data)

// 	http.HandleFunc("/list-files", handleListFiles)
// 	http.HandleFunc("/download", handleDownload)

// 	fmt.Println("server started at localhost:9000")
// 	http.ListenAndServe(":9000", nil)
// }

func main() {
	databases.StartDB()
	routers.Routes()
}
