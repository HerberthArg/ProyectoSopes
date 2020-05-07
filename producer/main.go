package main

import (
    "fmt"
    "log"
    "strconv"
    "github.com/streadway/amqp"
)

func failOnError(err error,msg string){
	if err != nil{
		log.Fatalf("%s: %s",msg,err)
	}
}

func main(){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err,"Fallo al  conectar con Rabbit")
	defer conn.Close()

	ch, err:= conn.Channel()
	failOnError(err,"Fallo al abrir el canal")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"cola",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Fallo al crear la cola")

	for i:=0; i<10; i++ {
		numero:=strconv.Itoa(i)
		texto:=fmt.Sprintf("%s %s","mensaje",numero)
		body:= texto
		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType: "text/plain",
				Body: []byte(body),
			},
		)
		failOnError(err,"Fallo al enviar el mensaje")
		log.Printf("[x] Enviado %s",body)
	}
}