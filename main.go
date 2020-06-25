package main

import (
	"CanITravelTo/Controller"
	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
	"os"
)

const PORT = ":443"

func main() {

	creds := os.Args[1]
	router := gin.Default()

	Controller.InitDatabase(creds)
	//Controller.CovidRetrieval()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /travel?destination=Spain&origin=USA

	router.OPTIONS("/travel", Controller.OptionsHandler)
	router.POST("/travel", Controller.PostHandler)
	router.GET("/travel", Controller.GetHandler)
	router.GET("/test", Controller.GetHandlerTest)

	router.RunTLS(PORT, creds+"/https-server.crt", creds+"/https-server.key")

	gocron.Every(1).Day().At("10:30:01").Do(Controller.CovidRetrieval)
	<-gocron.Start()

}
