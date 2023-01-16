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

func Readfiles() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	var Pokemon_zh = Getcontent("../data/Resources/zh/text_Species_zh.txt")
	var Abilities_zh = Getcontent("../data/Resources/zh/text_Abilities_zh.txt")
	var Moves_zh = Getcontent("../data/Resources/zh/text_Moves_zh.txt")
	var Types_zh = Getcontent("../data/Resources/zh/text_Types_zh.txt")
	var Natures_zh = Getcontent("../data/Resources/zh/text_Natures_zh.txt")
	var Items_zh = Getcontent("../data/Resources/zh/text_Items_zh.txt")
	var Pokemon_en = Getcontent("../data/Resources/en/text_Species_en.txt")
	var Abilities_en = Getcontent("../data/Resources/en/text_Abilities_en.txt")
	var Moves_en = Getcontent("../data/Resources/en/text_Moves_en.txt")
	var Types_en = Getcontent("../data/Resources/en/text_Types_en.txt")
	var Natures_en = Getcontent("../data/Resources/en/text_Natures_en.txt")
	var Items_en = Getcontent("../data/Resources/en/text_Items_en.txt")
	var mp = map[string][]string{
		"Pokemon_zh":   Pokemon_zh,
		"Abilities_zh": Abilities_zh,
		"Moves_zh":     Moves_zh,
		"Types_zh":     Types_zh,
		"Natures_zh":   Natures_zh,
		"Items_zh":     Items_zh,
		"Pokemon_en":   Pokemon_en,
		"Abilities_en": Abilities_en,
		"Moves_en":     Moves_en,
		"Types_en":     Types_en,
		"Natures_en":   Natures_en,
		"Items_en":     Items_en,
	}
	for k, v := range mp {
		rdb.Del(k)
		rdb.RPush(k, v)
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
