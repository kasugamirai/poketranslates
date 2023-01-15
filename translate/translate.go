package translate

import "pokemon.com/data"

var Pokemon_zh, Abilities_zh, Types_zh, Natures_zh, Items_zh, Moves_zh, Pokemon_en, Abilities_en, Types_en, Natures_en, Items_en, Moves_en []string = data.GetDataZh()[0], data.GetDataZh()[1], data.GetDataZh()[2], data.GetDataZh()[3], data.GetDataZh()[4], data.GetDataZh()[5], data.GetDataEn()[0], data.GetDataEn()[1], data.GetDataEn()[2], data.GetDataEn()[3], data.GetDataEn()[4], data.GetDataEn()[5]
var Pokemon_zh_set map[string]int = ListToMap(Pokemon_zh)
var Abilities_zh_set map[string]int = ListToMap(Abilities_zh)
var Types_zh_set map[string]int = ListToMap(Types_zh)
var Moves_zh_set map[string]int = ListToMap(Moves_zh)
var Natures_zh_set map[string]int = ListToMap(Natures_zh)
var Items_zh_set map[string]int = ListToMap(Items_zh)

type Pokemon struct {
	Nickname string
	Pokemon  string
	Gender   string
	Item     string
	Ability  string
	Level    string
	TeraType string
	EVs      string
	Nature   string
	Move_1   string
	Move_2   string
	Move_3   string
	Move_4   string
}

func ListToMap(list []string) map[string]int {
	set := make(map[string]int)
	for i, ch := range list {
		set[ch] = i
	}
	return set
}

func Zh_to_En(en []string, Zh_map map[string]int, target string) string {
	if Zh_map[target] > -1 {
		return en[Zh_map[target]]
	}
	return ""
}

func Result(input Pokemon) string {
	nickname := input.Nickname
	pokemon := Zh_to_En(Pokemon_en, Pokemon_zh_set, input.Pokemon)
	gender := input.Gender
	item := Zh_to_En(Items_en, Items_zh_set, input.Item)
	ability := Zh_to_En(Abilities_en, Abilities_zh_set, input.Ability)
	level := input.Level
	TeraType := Zh_to_En(Types_en, Types_zh_set, input.TeraType)
	EVs := input.EVs
	Nature := Zh_to_En(Natures_en, Natures_zh_set, input.Nature)
	Move_1 := Zh_to_En(Moves_en, Moves_zh_set, input.Move_1)
	Move_2 := Zh_to_En(Moves_en, Moves_zh_set, input.Move_2)
	Move_3 := Zh_to_En(Moves_en, Moves_zh_set, input.Move_3)
	Move_4 := Zh_to_En(Moves_en, Moves_zh_set, input.Move_4)

	var ans string

	if nickname != "" {
		ans = nickname + "(" + pokemon + ")"
	} else {
		ans = pokemon
	}
	if gender != "" {
		ans += "(" + gender + ")"
	}
	if item != "" {
		ans += " @ " + item
	}
	ans += "\nAbility: " + ability
	if level != "" {
		ans += "\nLevel: " + level
	}
	if TeraType != "" {
		ans += "\nTera Type: " + TeraType
	}
	if EVs != "" {
		ans += "\nEVs: " + EVs
	}
	ans += "\n" + Nature + " Nature"
	ans += "\n" + "- " + Move_1 + "\n" + "- " + Move_2 + "\n" + "- " + Move_3 + "\n" + "- " + Move_4 + "\n" + "\n"
	return ans
}
