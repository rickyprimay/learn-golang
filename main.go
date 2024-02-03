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

func createPersons(c *gin.Context){
    var newPerson person
    if err := c.BindJSON(&newPerson); err != nil{
        return
    }
    persons = append(persons, newPerson)
    c.IndentedJSON(http.StatusCreated, newPerson)
}
func getPersonByID(c *gin.Context) {
    id := c.Param("id")

    for _, p := range persons {
        if p.ID == id {
            c.IndentedJSON(http.StatusOK, p)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
}


func main() {
    router := gin.Default()
    router.GET("/persons", getPersons)
    router.POST("/persons", createPersons)
    router.GET("/persons/:id", getPersonByID)

    router.Run("localhost:8082")
}
