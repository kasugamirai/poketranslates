package data

import (
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis"
)

func Getcontent(path string) []string {
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

var items = []string{"Pokemon", "Abilities", "Types", "Natures", "Items", "Moves"}

const baseDir = "../data/Resources/"

func Readfiles() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var mp = map[string]PokemonItem{}

	for _, item := range items {
		zh := Getcontent(baseDir + "zh/text_" + item + ".txt")
		en := Getcontent(baseDir + "en/text_" + item + ".txt")
		mp[item] = PokemonItem{zh, en}
	}
	for k, v := range mp {
		rdb.Del(k+"_zh", k+"_en")
		rdb.RPush(k+"_zh", v.Zh)
		rdb.RPush(k+"_en", v.En)
	}
}

func GetDataZh() [][]string {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	Pokemon_zh, _ := rdb.LRange("Pokemon_zh", 0, -1).Result()
	Abilities_zh, _ := rdb.LRange("Abilities_zh", 0, -1).Result()
	Moves_zh, _ := rdb.LRange("Moves_zh", 0, -1).Result()
	Types_zh, _ := rdb.LRange("Types_zh", 0, -1).Result()
	Natures_zh, _ := rdb.LRange("Natures_zh", 0, -1).Result()
	Items_zh, _ := rdb.LRange("Items_zh", 0, -1).Result()
	ans := [][]string{Pokemon_zh, Abilities_zh, Types_zh, Natures_zh, Items_zh, Moves_zh}
	return ans
}

func GetDataEn() [][]string {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	Pokemon_en, _ := rdb.LRange("Pokemon_en", 0, -1).Result()
	Abilities_en, _ := rdb.LRange("Abilities_en", 0, -1).Result()
	Moves_en, _ := rdb.LRange("Moves_en", 0, -1).Result()
	Types_en, _ := rdb.LRange("Types_en", 0, -1).Result()
	Natures_en, _ := rdb.LRange("Natures_en", 0, -1).Result()
	Items_en, _ := rdb.LRange("Items_en", 0, -1).Result()
	ans := [][]string{Pokemon_en, Abilities_en, Types_en, Natures_en, Items_en, Moves_en}
	return ans
}

func Json_zh() map[string][]string {
	ret := map[string][]string{}
	data := []string{"Pokemon", "Ability", "TeraType", "Nature", "Item", "Move"}
	tmp := GetDataZh()
	for i, ch := range data {
		ret[ch] = tmp[i]
	}
	return ret
}
