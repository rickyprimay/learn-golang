package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type person struct {
    ID       string `json:"id"`
    FullName string `json:"fullname"`
    Age      int    `json:"age"`
}

var persons = []person{
    {ID: "1", FullName: "Ricky Prima", Age: 23},
    {ID: "2", FullName: "Yuda Putra", Age: 22},
}

func getPersons(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, persons)
}

func main() {
    router := gin.Default()
    router.GET("/persons", getPersons)

    router.Run("localhost:8080")
}
