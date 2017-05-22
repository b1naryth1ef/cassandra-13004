package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gocql/gocql"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// Constantly update a row based on id
func writeStress(session *gocql.Session, id int) {
	var err error

	for true {
		err = session.Query(`UPDATE guilds SET name=? WHERE id=?`, RandStringBytes(32), id).Exec()
		if err != nil {
			log.Printf("ERROR ON WRITE %v: %v", id, err)
			return
		}
	}
}

// Read and validate a row based on id
func readStress(session *gocql.Session, id int) {
	var err error

	for true {
		row := make(map[string]interface{})
		err = session.Query(`SELECT * FROM guilds WHERE id=?`, id).MapScan(row)
		if err != nil {
			log.Printf("ERROR ON READ %v: %v", id, err)
			time.Sleep(time.Second * 2)
			return
		} else {
			if len(row["name"].(string)) != 32 {
				log.Printf("BAD NAME LEN %v", id)
			}
		}
	}
}

func main() {
	ids := []int{
		295331720679522316,
		295331734109814785,
	}

	// connect to the cluster
	cluster := gocql.NewCluster("10.0.3.148")
	cluster.Keyspace = "test13004"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	for _, id := range ids {
		go readStress(session, id)
		go readStress(session, id)
		go writeStress(session, id)
		go writeStress(session, id)
		go writeStress(session, id)
		go writeStress(session, id)
		go writeStress(session, id)
		go writeStress(session, id)
		go writeStress(session, id)
		go writeStress(session, id)
	}

	time.Sleep(time.Second * 10000)
}
