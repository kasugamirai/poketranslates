package main

import (
	"encoding/json"
	"fmt"

	"pokemon.com/translate"

	"github.com/gin-gonic/gin"
)

func GetTeams(c *gin.Context) []string {
	jsonBlob := []byte(c.PostForm("teams"))
	var pokemons []translate.Pokemon
	var ans []string
	err := json.Unmarshal(jsonBlob, &pokemons)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, ch := range pokemons {
		ans = append(ans, translate.Result(ch))
	}
	return ans
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
