package main

import (
	"fmt"
	"github.com/vanilla/go-jwt-crud/api/common"
	"github.com/vanilla/go-jwt-crud/api/router"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	var err error

	err = common.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Printf("Listening on port %s\n", common.Config.Port)

	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", common.Config.Port), r))
}
