package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

type CollectionData struct {
	Collection Collection `json:"collection"`
}
type Collection struct {
	Size   int     `json:"size"`
	Traits []Trait `json:"traits"`
}

type Character struct {
	Name       string      `json:"name"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type Trait struct {
	TraitType string   `json:"trait_type"`
	Values    []string `json:"values"`
}

func buildCharacters(t []Trait) []Character {
	var character []Character
	build(t, 0, &character, nil)
	return character
}
func build(t []Trait, idx int, ch *[]Character, c []Attribute) {
	if idx == len(t) {
		*ch = append(*ch, Character{Attributes: append([]Attribute{}, c...)}) //deep copy
		return
	}
	for _, val := range t[idx].Values {
		attr := Attribute{
			TraitType: t[idx].TraitType,
			Value:     val,
		}
		build(t, idx+1, ch, append(c, attr))
	}
}
func main() {
	rawInput := "{\n  \"collection\": {\n    \"size\": 5,\n    \"traits\": [\n      {\n        \"trait_type\": \"ear\",\n        \"values\": [\n          \"small\",\n          \"big\",\n          \"pointy\",\n          \"round\"\n        ]\n      },\n      {\n        \"trait_type\": \"mouth\",\n        \"values\": [\n          \"smile\",\n          \"frown\",\n          \"neutral\",\n          \"open\"\n        ]\n      },\n      {\n        \"trait_type\": \"nose\",\n        \"values\": [\n          \"small\",\n          \"big\",\n          \"sharp\",\n          \"flat\"\n        ]\n      }\n    ]\n  }\n}"
	c := CollectionData{}
	err := json.Unmarshal([]byte(rawInput), &c)
	if err != nil {
		return
	}

	characters := buildCharacters(c.Collection.Traits)
	//var characters []Character
	//for i := 0; i < c.Collection.Size; i++ {
	//	var char Character
	//	char.Name = fmt.Sprintf("Character: %d", i)
	//	for _, trait := range c.Collection.Traits {
	//		attr := Attribute{
	//			TraitType: trait.TraitType,
	//			Value:     pickRandom(trait.Values),
	//		}
	//		char.Attributes = append(char.Attributes, attr)
	//	}
	//	characters = append(characters, char)
	//}
	output, err := json.MarshalIndent(characters, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}
func pickRandom(s []string) string {
	n := len(s)
	return s[rand.Intn(n)]
}
