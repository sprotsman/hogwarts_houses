package main

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
)

func usage() {
	_, file := path.Split(os.Args[0])
	fmt.Printf("Please type a Hogwarts house name.\n\n")
	fmt.Printf("    %s [option]\n", file)
	fmt.Printf("\nOptions: Gryffindor | Hufflepuff | Ravenclaw | Slytherin | unknown\n")
}

func main() {

	args := os.Args[1:]
	if len(args) == 0 {
		usage()
		return
	}

	houses := map[string][]string {
		"gryffindor": {"Dumbledore, Albus Percival Wulfric Brian",
			"McGonagall, Minerva", "Hagrid, Rubeus", "Weasley, Arthur",
			"Weasley (née Prewett), Molly", "Black, Sirius", "Evans, Lily",
			"Lupin, Remus", "Pettigrew, Peter", "Potter, James",
			"Weasley, Bill", "Weasley, Charlie", "Coote, Ritchie",
			"Frobisher, Victoria (Vicky)", "Hooper, Geoffrey", "Kirke, Andrew",
			"Robins, Demelza", "Sloper, Jack", "Weasley, Percy Ignatius",
			"Wood, Oliver", "Johnson, Angelina", "Jordan, Lee",
			"Spinnet, Alicia", "Stimpson, Patricia", "Towler, Kenneth",
			"Weasley, Fred", "Weasley, George", "Bell, Katie",
			"McLaggen, Cormac", "Brown, Lavender", "Finnigan, Seamus",
			"Granger, Hermione", "Longbottom, Neville", "Patil, Parvati",
			"Potter, Harry", "Thomas, Dean", "Weasley, Ronald \"Ron\"",
			"Creevey, Colin", "Weasley, Ginny", "Vane, Romilda",
			"Creevey, Dennis", "MacDonald, Natalie", "Peakes, Jimmy",
			"Abercrombie, Euan", "Potter, James Sirius", "Rose, Granger-Weasley"},
		"hufflepuff": {"Hufflepuff, Helga", "Hengist of Woordcroft",
			"The Fat Friar", "Wenlock, Bridget", "Lufkin, Artemisia",
			"Stump, Grogan", "McPhail, Dugald", "Scamander, Newt",
			"Kettleburn, Silvanus", "Sprout, Pomona", "Tonks, Nymphadora",
			"Diggory, Cedric", "Macmillan, Ernie", "Abbott, Hannah",
			"Finch-Fletchley, Justin", "Bones, Susan", "Jones, Megan",
			"Smith, Zacharias", "Cadwallader", "Stebbins", "Summerby", "Summers",
			"Madley, Laura", "Branstone, Eleanor", "Whitby, Kevin",
			"Cauldwell, Owen", "Lupin, Teddy", "Puffett, Eglantine"},
		"ravenclaw": {"Ravenclaw, Helena", "The Grey Lady", "Flitwick, Filius",
			"Warren, Myrtle Elizabeth (Moaning Myrtle)", "Fancourt, Perpetua",
			"Laverne de Montmorency", "Wildsmith, Ignatia", "Bagnold, Millicent",
			"McLaird, Lorcan", "Uric the Oddball", "Lockhart, Gilderoy",
			"Clearwater, Penelope", "Davies, Roger", "Carmichael, Eddie",
			"Chang, Cho", "Edgecombe, Marietta", "Belby, Marcus", "Boot, Terry",
			"Brocklehurst, Mandy", "Corner, Michael", "Cornfoot, Stephen",
			"Entwhistle, Kevin", "Goldstein, Anthony", "Li, Su",
			"McDougal, Morag", "Patil, Padma", "Turpin, Lisa", "Lovegood, Luna",
			"Ackerley, Stewart", "Quirke, Orla", "Hilliard, Robert", "Bradley"},
		"slytherin": {"Slytherin, Salazar", "The Bloody Baron",
			"Merlin \"Prince of Enchanters\"", "Corvinus Gaunt",
			"Slughorn, Horace", "Black, Phineas Nigellus", "Malfoy, Abraxas",
			"Avery (Snape's contemporary)", "Nott Sr.",
			"Mulciber (possible Sr.)", "Riddle, Tom Marvolo (Lord Voldemort)",
			"Umbridge, Dolores", "Lestrange, Rodolphus", "Lestrange, Rabastan",
			"Black, Andromeda", "Black, Bellatrix", "Black, Narcissa",
			"Malfoy, Lucius", "Avery (Riddle's contemporary)", "Mulciber",
			"Rosier, Evan", "Snape, Severus", "Wilkes", "Black, Regulus",
			"Bletchley, Miles", "Higgs, Terence", "Urquhart", "Vaisey",
			"Flint, Marcus", "Bole", "Derrick", "Montague", "Pucey, Adrian",
			"Warrington, C.", "Bulstrode, Millicent", "Crabbe, Vincent",
			"Davis, Tracey (witch)", "Goyle, Gregory", "Greengrass, Daphne",
			"Malfoy, Draco", "Nott, Theodore", "Parkinson, Pansy",
			"Zabini, Blaise", "Harper",
			"Astoria Greengrass (could be Asteria, sister to Daphne)",
			"Baddock, Malcolm", "Pritchard, Graham", "Potter, Albus Severus",
			"Malfoy, Scorpius"},
		"unknown": {"Bundy, K.", "Capper, S.", "Dorny, J.", "Dobbs, Emma",
			"Dunstan, B.", "Gudgeon, Davey", "Hornby, Olive", "Jorkins, Bertha",
			"MacDonald, Mary", "Midgen, Eloise", "Moon", "Perks, Sally-Ann",
			"Derek"},
	}

	house := strings.ToLower(args[0])
	students := houses[strings.ToLower(args[0])]

	if students == nil {
		fmt.Printf("Sorry, Hogwarts doesn't have a house named %q.\n\n", strings.Title(house))
		usage()
		return
	}

	// leave the slices in the map unchanged ... make a copy
	studentsCopy := append([]string(nil), students...)
	sort.Strings(studentsCopy)

	var str string
	if house == "unknown" {
		str = "Students whose house is"
	} else {
		str = "Students belonging to house"
	}

	fmt.Printf("\n——— %s %s ———\n\n", str, strings.Title(house))

	for _, student := range studentsCopy {
		fmt.Printf("\t• %s\n", student)
	}

	fmt.Println()
}
