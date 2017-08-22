package main

import (
	"encoding/gob"
	"log"
	"os"
)

var store map[string]interface{}

func InitStore() {
	gob.Register(StoreUserObject{})
	gob.Register(StoreMessageObject{})
}

func initNewEmptyStore() {
	store = make(map[string]interface{})
}

func ReadStoreFromBinary(filePath string) {
	log.Println("Loading store.")
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		log.Println("Could not read store from file. Initializing empty one.")
		initNewEmptyStore()
		return
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&store)
	if err != nil {
		log.Println(err)
	}
}

func FlushStoreToBinary(filePath string) {
	log.Println("Flushing store.")
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		log.Println(err)
	}

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(&store)
	if err != nil {
		log.Println(err)
	}
}

func StoreGet(key string) interface{} {
	return store[key]
}

func StorePut(key string, value interface{}) {
	store[key] = value
}

func StoreDelete(key string) {
	delete(store, key)
}

func StoreGetMap() map[string]interface{} {
	return store
}
