package main

import ("fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    fmt.Printf("Hello, world.\n")
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello,World!")
    })
    r.Run()
}
