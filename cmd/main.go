package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"main/infra/kafka"
	repository2 "main/infra/repository"
	usecase2 "main/usecase"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/fullcycle")

	if err != nil {
		log.Fatalln(err)
	}

	repository := repository2.CourseMySQLRepository{DB: db}
	usecase := usecase2.CreateCourse{Repository: repository}

	var msgChannel = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "appgo",
	}
	topics := []string{"courses"}
	consumer := kafka.NewConsumer(configMapConsumer, topics)

	go consumer.Consume(msgChannel)

	for msg := range msgChannel {
		var input usecase2.CreateCourseInputDTO
		json.Unmarshal(msg.Value, &input)

		output, err := usecase.Execute(input)

		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("Output: ", output)
		}
	}
}

// {"name":"Curso Full Cycle","description":"Full Cycle 3.0","status":"Pending"}
