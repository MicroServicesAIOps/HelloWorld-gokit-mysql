package api

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type IMyService interface {
	Health() []Health
	Factorial(n int) Factorial
}

type MyService struct{}

type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}

type Factorial struct {
	FacVal uint64 `json:"facval"`
}

func (s MyService) Health() []Health {
	var health []Health
	app := Health{"demo-mysql", "OK", time.Now().String()}

	health = append(health, app)

	return health
}

func (s MyService) Factorial(n int) Factorial {
	var value uint64 = 1
	log.Println("num is ", n)
	for i := 1; i <= n; i++ {
		value *= uint64(i)
	}
	log.Println(n, "!   is ", value)
	return Factorial{value}
}
