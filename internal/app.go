package app

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/yosa12978/MyShirts/internal/api"
	"github.com/yosa12978/MyShirts/internal/config"
)

func Run() {
	router := api.Init()
	conf := config.GetConfig()
	rand.Seed(time.Now().UnixNano())
	server := http.Server{
		Addr:    conf.Address + ":" + strconv.Itoa(conf.Port),
		Handler: router,
	}
	server.ListenAndServe()
}
