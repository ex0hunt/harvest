package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func CreateIndexHandler(w http.ResponseWriter, r *http.Request) {
	type JsonData struct {
		IndexName string `json:"index_name"`
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()
	var b JsonData
	err := jsonDecoder.Decode(&b)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = CreateIndex(b.IndexName)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
		return
	}
	err = json.NewEncoder(w).Encode(map[string]string{"msg": "ok"})
	if err != nil{
		log.Fatal(err)
	} else {
		log.Println("Index created", b.IndexName)
	}
	return
}

func DataInsertHandler(w http.ResponseWriter, r *http.Request) {
	type JsonData struct {
		IndexName string `json:"index_name"`
		DataId string `json:"data_id"`
		DataOwner string `json:"data_owner"`
		Data string `json:"data"`
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()
	var b JsonData
	err := jsonDecoder.Decode(&b)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = InsertData(b.IndexName, b.DataId, b.DataOwner, b.Data)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
		return
	}
	err = json.NewEncoder(w).Encode(map[string]string{"msg": "ok"})
	if err != nil{
		log.Fatal(err)
	} else {
		log.Printf("Index: %s; Inserted data: %s\n", b.IndexName, b.DataId)
	}
	return
}

func DataSearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	indexName := r.FormValue("index_name")
	searchQuery := r.FormValue("search_query")

	result, err := SearchData(indexName, searchQuery)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"msg": "ok", "result": result})
	if err != nil{
		log.Fatal(err)
	} else {
		log.Printf("Index:%s; Search request: %s\n", searchQuery)
	}
	return
}

func DataDeleteHandler(w http.ResponseWriter, r *http.Request) {
	type JsonData struct {
		IndexName string `json:"index_name"`
		DataId string `json:"data_id"`
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	jsonDecoder := json.NewDecoder(r.Body)
	jsonDecoder.DisallowUnknownFields()
	var b JsonData
	err := jsonDecoder.Decode(&b)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = DeleteData(b.IndexName, b.DataId)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"msg": err.Error()})
		return
	}
	err = json.NewEncoder(w).Encode(map[string]string{"msg": "ok"})
	if err != nil{
		log.Fatal(err)
	} else {
		log.Printf("Index:%s; Data deleted: %s\n", b.IndexName, b.DataId)
	}
	return
}