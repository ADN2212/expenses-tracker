package main

import (
	"expenses-tracker.com/DB"
	"expenses-tracker.com/endpoints"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func init() {
	DB.InitDataBase()
}

func main() {

	envErr := godotenv.Load("../.env")

	if envErr != nil {
		panic(".env file not found")
	} else {
		fmt.Println(".env file loaded succesfully")
	}

	router := gin.Default()
	//Transaction endpoints:
	router.GET("transactions", enpoints.GetTransactions)
	router.POST("transactions", enpoints.AddTransaction)
	router.DELETE("transactions/:id", enpoints.DeleteTransaction)

	port := os.Getenv("PORT")

	if len(port) == 0 {
		fmt.Println("PORT not found, setting to 8080")
		port = "8080"
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			res := origin == "http://localhost:5173"
			fmt.Println("-----------------------------")
			fmt.Println(origin)			
			fmt.Println(res)
			fmt.Println("-----------------------------")
			return res
		},
		MaxAge: 10 * time.Second,
	}))

	router.Run(fmt.Sprintf("localhost:%s", port))
}
