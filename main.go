package main

import (
        "cicd-test/calculation"
        "github.com/satori/go.uuid"
        log "github.com/sirupsen/logrus"
)

func main() {
        a, b := 5, 2
        log.Printf("values: %v %v\n", a, b)
        log.Println(calculation.Sum(a, b))
        log.Println(calculation.Sub(a, b))
        log.Println("New UUID: ", uuid.NewV4().String())
}
