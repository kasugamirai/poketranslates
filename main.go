package main

import (
	"encoding/json"
	"fmt"

	"pokemon.com/data"
	"pokemon.com/translate"

	"github.com/gin-gonic/gin"
)

func GetTeams(c *gin.Context) string {
	jsonBlob := []byte(c.PostForm("teams"))
	var pokemons []translate.Pokemon
	var ans string
	err := json.Unmarshal(jsonBlob, &pokemons)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, ch := range pokemons {
		ans += translate.Result(ch)
	}
	return ans
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("web/template/*")
	r.Static("/static", "static")
	r.GET("/teambulider", func(c *gin.Context) {
		c.HTML(200, "teambulider.html", gin.H{
			"title": "teambulider",
		})
	})
	r.GET("/readfiles", func(c *gin.Context) {
		data.Readfiles()
	})
	r.GET("/datazh", func(c *gin.Context) {
		c.JSON(200, data.GetDataZh())
	})
	r.POST("/upload", func(c *gin.Context) {
		ret := GetTeams(c)
		c.String(200, string(ret))

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
