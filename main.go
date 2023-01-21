package main

import (
	"encoding/json"
	"fmt"

	"pokemon.com/data"
	"pokemon.com/translate"

	"github.com/gin-gonic/gin"
)

var links []string

func GetTeams(c *gin.Context) string {
	jsonBlob := []byte(c.PostForm("pokemon"))
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
	r.LoadHTMLGlob("./web/template/*")
	r.Static("/static", "static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "teambuilder.html", gin.H{
			"title": "teambuilder",
		})
	})
	r.GET("/readfiles", func(c *gin.Context) {
		data.Readfiles()
	})
	r.GET("/data_zh", func(c *gin.Context) {
		c.JSON(200, data.GetData("zh"))
	})
	r.POST("/upload", func(c *gin.Context) {
		var pokemons []translate.Pokemon
		c.BindJSON(&pokemons)
		var ans string
		for _, pokemon := range pokemons {
			ans += translate.Result(pokemon)
		}
		link := data.SaveTeams(ans)
		links = append(links, link)
		c.JSON(200, gin.H{
			"link": link,
		})
	})

	r.POST("/search", func(ctx *gin.Context) {
		var link string
		ctx.BindJSON(&link)
		ans := data.RetTeams(link)
		ctx.JSON(200, gin.H{
			"res": ans,
		})
	})

	r.GET("/find", func(ctx *gin.Context) {
		ctx.HTML(200, "search.html", gin.H{
			"title": "search",
		})
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
