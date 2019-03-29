package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	<div role=\"note\" class=\"hatnote\">For other uses, see <a href=\"\/wiki\/Pizza_(disambiguation)\" class=\"mw-disambig\" title=\"Pizza (disambiguation)\">Pizza (disambiguation)<\/a>.<\/div>\n<table class=\"infobox hrecipe adr\" style=\"width:22em\">\n<caption class=\"fn\"><span>Pizza<\/span><\/caption>\n<tr>\n<td colspan=\"2\" style=\"text-align:center\"><a href=\"\/wiki\/File:Pepperoni_pizza.jpg\" class=\"image\"><img alt=\"Pepperoni pizza.jpg\" src=\"\/\/upload.wikimedia.org\/wikipedia\/commons\/thumb\/d\/d1\/Pepperoni_pizza.jpg\/220px-Pepperoni_pizza.jpg\" width=\"220\" height=\"139\" srcset=\"\/\/upload.wikimedia.org\/wikipedia\/commons\/thumb\/d\/d1\/Pepperoni_pizza.jpg\/330px-Pepperoni_pizza.jpg 1.5x, \/\/upload.wikimedia.org\/wikipedia\/commons\/thumb\/d\/d1\/Pepperoni_pizza.jpg\/440px-Pepperoni_pizza.jpg 2x\" data-file-width=\"959\" data-file-height=\"606\" \/><\/a>\n<div style=\"padding-bottom:0.25em;border-bottom:1px solid #aaa;\">Pizza topped with <a href=\"\/wiki\/Pepperoni\" title=\"Pepperoni\">pepperoni<\/a><\/div>\n<\/td>\n<\/tr>\n<tr>\n<th scope=\"row\" style=\"padding-top:0.245em;line-height:1.15em; padding-right:0.65em;\">Type<\/th>\n<td><a href=\"\/wiki\/Flatbread\" title=\"Flatbread\">Flatbread<\/a><\/td>\n<\/tr>\n<tr>\n<th scope=\"row\" style=\"padding-top:0.245em;line-height:1.15em; padding-right:0.65em;\">Course<\/th>\n<td>Lunch or dinner<\/td>\n<\/tr>\n<tr class=\"note\">\n<th scope=\"row\" style=\"padding-top:0.245em;line-height:1.15em; padding-right:0.65em;\">Place of origin<\/th>\n<td class=\"country-name\"><a href=\"\/wiki\/Naples\" title=\"Naples\">Naples<\/a>, <a href=\"\/wiki\/Campania\" title=\"Campania\">Campania<\/a>, <a href=\"\/wiki\/Italy\" title=\"Italy\">Italy<\/a><\/td>\n<\/tr>\n<tr>\n<th scope=\"row\" style=\"padding-top:0.245em;line-height:1.15em; padding-right:0.65em;\">Serving temperature<\/th>\n<td>Hot or warm<\/td>\n<\/tr>\n<tr>\n<th scope=\"row\" style=\"padding-top:0.245em;line-height:1.15em; padding-right:0.65em;\">Main ingredients<\/th>\n<td class=\"ingredient\">Dough, often <a href=\"\/wiki\/Tomato_sauce\" title=\"Tomato sauce\">tomato sauce<\/a>, <a href=\"\/wiki\/Cheese\" title=\"Cheese\">cheese<\/a><\/td>\n<\/tr>\n<tr>\n<th scope=\"row\" style=\"padding-top:0.245em;line-height:1.15em; padding-right:0.65em;\">Variations<\/th>\n<td><a href=\"\/wiki\/Calzone\"
	`
	s := strings.Split(text, ">")

	for i := 0; i < len(s); i++ {
		if strings.HasPrefix(s[i], "<") {
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}

	var count int

	for i := 0; i < len(s); i++ {
		tmpSlice := strings.Split(s[i], "<")
		tmpItem := strings.ToLower(tmpSlice[0])
		count += strings.Count(tmpItem, "pizza")
	}
	fmt.Println("*** count = ", count)
}
