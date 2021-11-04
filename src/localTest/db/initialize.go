package db

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UserID  string    `json:"UserID"`
	Name    string    `json:"Name"`
	Created time.Time `json:"Created"`
}

var container string

func runDocker() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		log.Printf("[Docker Container Activating...]")

		args := []string{
			"run", "-d", "--rm",
			"-p", "3306:3306",
			"-e", "MYSQL_ROOT_PASSWORD=rootpass",
			"-e", "MYSQL_DATABASE=mytest",
			"mysql:5.7"}
		cmd := exec.Command("docker", args...)
		// args := []string{"-f", "../db/docker-compose.yaml", "up", "-d"}
		// cmd := exec.Command("docker-compose", args...)

		var stdout bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("[Start Docker Failure] stderr: %s | err: %s\n", stderr.String(), err)
			return
		}

		if !cmd.ProcessState.Success() {
			log.Fatalf("[Start Docker Failure] stderr: %s\n", stderr.String())
			return
		} else {
			container = stdout.String()
			log.Printf("[Start Dokcer Success] stdout: %s\n", stdout.String())
			return
		}
	}()

	// keep waiting all go runtine ready
	waitGroup.Wait()
}

func stopDocker() {
	log.Printf("[Docker Container Killing...]")

	// TODO: I have no idea about that the statement `docker rm {CONTAINER_ID}` doesn't work.
	fmt.Printf("docker rm -fv %s", container)
	// cmd := exec.Command("docker", "rm", "-fv", container)
	cmd := exec.Command("bash", "-c", "docker kill \"$(docker ps -aq)\"")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("[Stop Docker Failure] stderr: %s | err: %s\n", stderr.String(), err)
		return
	}

	if !cmd.ProcessState.Success() {
		log.Fatalf("[Stop Docker Success] stdout: %s \n", stderr.String())
		return
	}
}

func initDBSchema() {
	log.Printf("[MySQL Table Schema Initializing...]")

	dbClient, err := Connect()
	if err != nil {
		log.Fatalf("[Mysql Connection Failure] err: %s\n", err)
	}
	defer dbClient.Close()

	// waiting until mysql container get ready
	for i := 0; i < 40; i++ {
		if err := dbClient.Ping(); err == nil {
			break
		} else {
			log.Println("connecting to MySQL server...")
			<-time.After(time.Duration(1) * time.Second)
		}
	}

	userTableCreateStatement := `
		CREATE TABLE IF NOT EXISTS Users(
			UserID VARCHAR(36) NOT NULL,
			NAME VARCHAR(36) NOT NULL,
			Created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY(UserID)
		)
	`

	_, err = dbClient.Exec(userTableCreateStatement)
	if err != nil {
		log.Fatalf("[Create User Table Failure] err: %s", err)
	}
}

func Init() {

	// execute docker-compose to build dokcer container
	runDocker()

	// prepare DB environment
	initDBSchema()
}

func Teardown() {

	// stop & remove docker container
	stopDocker()
}
