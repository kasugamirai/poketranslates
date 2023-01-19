package data

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func GetContent(path string) []string {
	paths := path
	content, err := os.ReadFile(paths)
	if err != nil {
		log.Fatal(err)
	}
	contents := strings.Split(string(content), "\n")
	return contents
}

type PokemonItem struct {
	Zh []string
	En []string
}

var items = []string{"Species", "Abilities", "Types", "Natures", "Items", "Moves"}

const baseDir = "data/Resources/"

func Readfiles() {

	var mp = map[string]PokemonItem{}

	for _, item := range items {
		zh := GetContent(baseDir + "zh/text_" + item + ".txt")
		en := GetContent(baseDir + "en/text_" + item + ".txt")
		mp[item] = PokemonItem{zh, en}
	}
	for k, v := range mp {
		rdb.Del(k+"_zh", k+"_en")
		rdb.RPush(k+"_zh", v.Zh)
		rdb.RPush(k+"_en", v.En)
	}
}

func GetData(name string) [][]string {
	ans := [][]string{}
	for _, item := range items {
		item, _ := rdb.LRange(item+"_"+name, 0, -1).Result()
		ans = append(ans, item)
	}
	return ans
}

func SaveTeams(team string) string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Int63n(time.Now().UnixNano())
	encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(num)))[:8]
	rdb.HSet("teamlist", encoded, team)
	return encoded
}

func RetTeams(link string) string {
	ret, _ := rdb.HGet("teamlist", link).Result()
	return ret
}
