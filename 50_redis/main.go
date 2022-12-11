package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

type Mahasiswa struct {
	Nama     string `redis:"nama"`
	NPM      string `redis:"npm"`
	IPK      string `redis:"ipk"`
	Semester string `redis:"semester"`
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Panic(err)
	}

	inputDataMahasiswa(conn, "Yulyano Thomas Djaya", "56419764", 4.00, 8)

	mahasiswa := getDataMahasiswa(conn, 1)

	fmt.Printf("Nama mahasiswa = %s\n", mahasiswa.Nama)
}

func inputDataMahasiswa(conn redis.Conn, nama, npm string, ipk float32, semester int) {

	_, err := conn.Do("HSET", "mahasiswa:1", "nama", nama, "npm", npm, "ipk", ipk, "semester", semester)
	if err != nil {
		log.Panic(err)
	}
}

func getDataMahasiswa(conn redis.Conn, n int) Mahasiswa {
	// HGETALL mengembalikan semua field yang ada pada objek tsb
	// redis.Values adalah reply helper yang mengconvert data
	// reply dengan tipe interface ke tipe []interface{}

	strConn := fmt.Sprintf("mahasiswa:%d", n)
	rep, err := redis.Values(conn.Do("HGETALL", strConn))
	if err != nil {
		log.Panic(err)
	}

	var mahasiswa Mahasiswa

	// ScanStruct akan mempopulate semua data reply ke struct
	// mahasiswa, ScanStruct juga otomatis mencasting tipe data
	// sesuai dengan yang didefinisikan di sruct

	err = redis.ScanStruct(rep, &mahasiswa)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%+v\n", mahasiswa)

	return mahasiswa
}
