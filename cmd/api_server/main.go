package main

import (
	"MangaHub/internal/auth"
	grpcclient "MangaHub/internal/grpc"
	"MangaHub/internal/manga"
	"MangaHub/internal/progress"
	"MangaHub/internal/tcp"
	"MangaHub/internal/udp"
	"MangaHub/pkg/database"
	"net"

	"log"

	"github.com/gin-gonic/gin"
)

var hub = tcp.NewHub()

func main() {
	go startHTTPServer(hub)
	starttcpserver()

}
func startHTTPServer(hub *tcp.Hub) {
	// create router
	r := gin.Default()

	//connect to the db
	db := database.InitDB("mangahub.db")
	err := database.TableCreate(db)
	if err != nil {
		log.Fatal("Unable to create database", err)
	}

	//call to auth
	r.POST("/auth/register", auth.Register(db))
	r.POST("/auth/login", auth.Login(db))

	//call to find manga

	r.GET("/manga", auth.AuthMiddleware(), manga.Search(db))
	r.GET("/manga/:id", auth.AuthMiddleware(), manga.Detail(db))

	//add and updtae user progress
	udpNotifier := udp.NewNotifier("localhost:9091")

	r.POST("/library/:manga_id", auth.AuthMiddleware(), progress.AddToLibrary(db, hub))
	r.PUT("/progress/:manga_id", auth.AuthMiddleware(), progress.UpdateProgress(db, hub, udpNotifier))
	r.GET("/library", auth.AuthMiddleware(), progress.GetLibrary(db))

	r.GET("/grpc/progress", auth.AuthMiddleware(), func(c *gin.Context) {
		userID := c.GetString("user_id")

		resp, err := grpcclient.GetUserProgress(userID)
		if err != nil {
			c.JSON(500, gin.H{"error": "gRPC error"})
			return
		}

		c.JSON(200, resp)
	})
	//run server
	log.Println("API server is listening and serve on port :8080")
	r.Run(":8080")
	//hTTP SERVER

}
func starttcpserver() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("TCP Sync Server running on :9090")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go tcp.HandleConnection(conn, hub)
	}
}
