package translate

import (
	"pokemon.com/data"
)

var Pokemon_en, Abilities_en, Types_en, Natures_en, Items_en, Moves_en []string = data.GetData("en")[0], data.GetData("en")[1], data.GetData("en")[2], data.GetData("en")[3], data.GetData("en")[4], data.GetData("en")[5]

type Language_Set struct {
	Pokemon_set   map[string]int
	Abilities_set map[string]int
	Types_set     map[string]int
	Moves_set     map[string]int
	Natures_set   map[string]int
	Items_set     map[string]int
}

type Pokemon struct {
	Nickname string
	Pokemon  string
	Gender   string
	Item     string
	Ability  string
	Level    string
	Shiny    string
	TeraType string
	IVs      string
	EVs      string
	Nature   string
	Move_1   string
	Move_2   string
	Move_3   string
	Move_4   string
}

var Zh_set Language_Set
var Ja_set Language_Set

func ListToMap(list []string) map[string]int {
	set := make(map[string]int)
	for i, ch := range list {
		set[ch] = i
	}
	return set
}

func Origin_to_Target(en []string, input_map map[string]int, target string) string {
	if p, ok := input_map[target]; ok {
		return en[p]
	}
	return ""
}

func Result(input Pokemon, lang Language_Set) string {
	nickname := input.Nickname
	pokemon := Origin_to_Target(Pokemon_en, lang.Pokemon_set, input.Pokemon)
	gender := input.Gender
	item := Origin_to_Target(Items_en, lang.Items_set, input.Item)
	ability := Origin_to_Target(Abilities_en, lang.Abilities_set, input.Ability)
	level := input.Level
	Shiny := input.Shiny
	TeraType := Origin_to_Target(Types_en, lang.Types_set, input.TeraType)
	IVs := input.IVs
	EVs := input.EVs
	Nature := Origin_to_Target(Natures_en, lang.Natures_set, input.Nature)
	Move_1 := Origin_to_Target(Moves_en, lang.Moves_set, input.Move_1)
	Move_2 := Origin_to_Target(Moves_en, lang.Moves_set, input.Move_2)
	Move_3 := Origin_to_Target(Moves_en, lang.Moves_set, input.Move_3)
	Move_4 := Origin_to_Target(Moves_en, lang.Moves_set, input.Move_4)

	var ans string

	if nickname != "" {
		ans = nickname + " (" + pokemon + ")"
	} else {
		ans = pokemon
	}
	if gender != "" {
		ans += " (" + gender + ")"
	}
	if item != "" {
		ans += " @ " + item
	}
	ans += "\nAbility: " + ability
	if level != "" {
		ans += "\nLevel: " + level
	}
	if Shiny != "" {
		ans += "\nShiny: " + Shiny
	}
	if TeraType != "" {
		ans += "\nTera Type: " + TeraType
	}
	if EVs != "" {
		ans += "\nEVs: " + EVs
	}
	if IVs != "" {
		ans += "\nIVs: " + IVs
	}
	ans += "\n" + Nature + " Nature"
	ans += "\n" + "- " + Move_1 + "\n" + "- " + Move_2 + "\n" + "- " + Move_3 + "\n" + "- " + Move_4 + "\n" + "\n"
	return ans
}

func Create_set(language *Language_Set, date_name string) {
	language.Pokemon_set = ListToMap(data.GetData(date_name)[0])
	language.Abilities_set = ListToMap(data.GetData(date_name)[1])
	language.Types_set = ListToMap(data.GetData(date_name)[2])
	language.Moves_set = ListToMap(data.GetData(date_name)[5])
	language.Natures_set = ListToMap(data.GetData(date_name)[3])
	language.Items_set = ListToMap(data.GetData(date_name)[4])
}

func init() {
	Create_set(&Zh_set, "zh")
	Create_set(&Ja_set, "ja")
}
