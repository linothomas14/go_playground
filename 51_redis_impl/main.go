package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gomodule/redigo/redis"
)

func GetPokemonWithoutRedis(w http.ResponseWriter, r *http.Request) {
	// Ambil query dari url -> (BASE_URL?key=value)
	pokemonName := r.URL.Query()["pokemon"][0]

	client := http.DefaultClient

	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon/"+pokemonName, nil)

	if err != nil {
		log.Panic(err)
	}

	res, err := client.Do(req)

	if err != nil {
		log.Panic(err)
	}

	bd, _ := ioutil.ReadAll(res.Body)
	w.Write(bd)
}

func GetPokemonWithRedis(w http.ResponseWriter, r *http.Request) {

	// Ambil query dari url -> (BASE_URL?key=value)
	pokemonName := r.URL.Query()["pokemon"][0]

	// membuat koneksi redis
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Panic(err)
	}

	// pengecekan pikachu sudah dicache atau belum
	reply, err := redis.Bytes(conn.Do("GET", pokemonName))
	if err == nil {
		w.Write(reply)
		return
	}
	// jika belum lakukan request API
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/pokemon/"+pokemonName, nil)
	if err != nil {
		log.Panic(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	bd, _ := ioutil.ReadAll(res.Body)

	// simpan data response di redis
	_, err = conn.Do("SET", pokemonName, string(bd))
	if err != nil {
		log.Panic(err)
	}
	// write response
	w.Write(bd)

}

func main() {
	http.HandleFunc("/pokemonwithoutredis", GetPokemonWithoutRedis)
	http.HandleFunc("/pokemonwithredis", GetPokemonWithRedis)
	log.Fatal(http.ListenAndServe(":3000", nil))

}
