package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func anal_menu() {
	//analysis main menu
	fmt.Println("Analysis")
	m := make(map[string]string)
	//reader := bufio.NewReader(os.Stdin)

	// Set key/value pairs using typical `name[key] = val`
	m["k1"] = fmt.Sprintf("%-25s", "Games by Day")
	m["k2"] = fmt.Sprintf("%-25s", "Games by Reason")
	m["k3"] = fmt.Sprintf("%-25s", "Games by Time of Day")
	m["k4"] = fmt.Sprintf("%-25s", "Games by Level-Tier")
	m["k5"] = fmt.Sprintf("%0s", "Deck Recommended for Deleting")
	m["k6"] = fmt.Sprintf("%0s", "Decks by Number of Cards")
	m["k7"] = fmt.Sprintf("%0s", "Decks by Number of Creatures")
	m["k8"] = fmt.Sprintf("%0s", "Decks by Number of Non-Creatures")
	m["k9"] = fmt.Sprintf("%24s", "Decks by Number of Lands")
	m["k10"] = fmt.Sprintf("%-24s", "Return to Main Menu")
	m["k11"] = fmt.Sprintf("%0s", "Quit")

	// print menu options
	fmt.Println(" 1:", m["k1"]+"5:", m["k5"])
	fmt.Println(" 2:", m["k2"]+"6:", m["k6"])
	fmt.Println(" 3:", m["k3"]+"7:", m["k7"])
	fmt.Println(" 4:", m["k4"]+"8:", m["k8"])
	fmt.Println("                            "+" 9:", m["k9"])
	fmt.Println("10:", m["k10"]+"11:", m["k11"])

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	choice, _ := strconv.Atoi(in.Text())

	switch choice {
	case 1:
		gamebydaymenu()
	case 2:
		analreasonmenu()
	case 3:
		analtimemenu()
	case 4:
		analvltiermenu()
	case 5:
		recdelete()
	case 6:
		deckbycardmenu("card")
	case 7:
		deckbycardmenu("creature")
	case 8:
		deckbycardmenu("non-creature")
	case 9:
		deckbycardmenu("land")
	case 10:
		main()
	case 11:
		os.Exit(0)
	}
}
func analdmenu(day string) {
	m := make(map[string]string)
	fmt.Println("Choose Options Below")
	m["k1"] = fmt.Sprintf("%-30s", "Best Deck for "+day)
	m["k2"] = fmt.Sprintf("%-30s", "Worst Deck for "+day)
	m["k3"] = fmt.Sprintf("%-30s", "Wins for a deck on "+day)
	m["k4"] = fmt.Sprintf("%-25s", "Loses for a deck on "+day)
	m["k5"] = fmt.Sprintf("%0s", "Wins for all decks on "+day)
	m["k6"] = fmt.Sprintf("%0s", "Loses for all decks on "+day)
	m["k10"] = fmt.Sprintf("%0s", "Return to Previous Menu")
	m["k11"] = fmt.Sprintf("%0s", "Return to Main Menu")
	m["k12"] = fmt.Sprintf("%0s", "Quit")

	// print menu options
	fmt.Println(" 1:", m["k1"]+"4:", m["k4"])
	fmt.Println(" 2:", m["k2"]+"5:", m["k5"])
	fmt.Println(" 3:", m["k3"]+"6:", m["k6"])
	println("")
	fmt.Println("10:", m["k10"])
	fmt.Println("11:", m["k11"])
	fmt.Println("12:", m["k12"])

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	choice, _ := strconv.Atoi(in.Text())

	switch choice {
	case 1:
		wkdychoice := historic()
		analweekday("best", day, "", wkdychoice)
	case 2:
		wkdychoice := historic()
		analweekday("worst", day, "", wkdychoice)
	case 3:
		wkdychoice := historic()
		println("Deck:")
		in.Scan()
		deck := in.Text()
		deck = validatedeck(deck, wkdychoice)
		analweekday(deck, day, "w", wkdychoice)
	case 4:
		wkdychoice := historic()
		println("Deck:")
		in.Scan()
		deck := in.Text()
		deck = validatedeck(deck, wkdychoice)
		analweekday(deck, day, "l", wkdychoice)
	case 5:
		wkdychoice := historic()
		analweekday("all", day, "w", wkdychoice)
	case 6:
		wkdychoice := historic()
		analweekday("all", day, "l", wkdychoice)
	case 10:
		gamebydaymenu()
	case 11:
		main()
	case 12:
		os.Exit(0)
	default:
		analdmenu(day)
	}
}
func analreasonmenu() {
	//analysis main menu
	fmt.Println("Reason")
	m := make(map[string]string)

	// Set key/value pairs using typical `name[key] = val`
	m["k1"] = fmt.Sprintf("%-50s", "Win Reasons Related to Mana")
	m["k2"] = fmt.Sprintf("%-50s", "Win Reasons Related to Creatures")
	m["k3"] = fmt.Sprintf("%-50s", "Win Reasons Related to Opponent Deck Type")
	m["k4"] = fmt.Sprintf("%-50s", "Win Reasons Related to Specific Deck")
	m["k5"] = fmt.Sprintf("%0s", "Lost Reasons Related to Mana")
	m["k6"] = fmt.Sprintf("%0s", "Lost Reasons Related to Creatures")
	m["k7"] = fmt.Sprintf("%0s", "Lost Reasons Related to Opponent Deck Type")
	m["k8"] = fmt.Sprintf("%24s", "Lost Reasons Related to Specific Deck")
	m["k9"] = fmt.Sprintf("%0s", "Customized Reasons")
	m["k10"] = fmt.Sprintf("%-50s", "Return to Previous Menu")
	m["k11"] = fmt.Sprintf("%0s", "Return to Main Menu")
	m["k12"] = fmt.Sprintf("%0s", "Quit")

	// print menu options
	fmt.Println(" 1:", m["k1"]+"5:", m["k5"])
	fmt.Println(" 2:", m["k2"]+"6:", m["k6"])
	fmt.Println(" 3:", m["k3"]+"7:", m["k7"])
	fmt.Println(" 4:", m["k4"]+"8:", m["k8"])
	fmt.Println("                            "+" 9:", m["k9"])
	fmt.Println("10:", m["k10"]+"11:", m["k11"])
	fmt.Println("12:", m["k12"])

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	choice, _ := strconv.Atoi(in.Text())

	switch choice {
	case 1:
		println("Do you want to specify a deck?(y/n)")
		in.Scan()
		deck := in.Text()
		deck = validateuserinput(deck, "confirm")
		rsnchoice := historic()
		if deck == "y" {
			println("Deck: ")
			in.Scan()
			deck = in.Text()
			deck = validatedeck(deck, rsnchoice)
		}
		analreason("mana", "w", deck, rsnchoice)
	case 2:
		println("Do you want to specify a deck?(y/n)")
		in.Scan()
		deck := in.Text()
		deck = validateuserinput(deck, "confirm")
		rsnchoice := historic()
		if deck == "y" {
			println("Deck: ")
			in.Scan()
			deck = in.Text()
			deck = validatedeck(deck, rsnchoice)
		}
		analreason("creature", "w", deck, rsnchoice)
	case 3:
		println("Unfinished")
		analreasonmenu()
	case 4:
		rsnchoice := historic()
		println("Deck:")
		in.Scan()
		deck := in.Text()
		deck = validatedeck(deck, rsnchoice)
		analreason("", "w", deck, rsnchoice)
	case 5:
		println("Do you want to specify a deck?(y/n)")
		in.Scan()
		deck := in.Text()
		deck = validateuserinput(deck, "confirm")
		rsnchoice := historic()
		if deck == "y" {
			println("Deck: ")
			in.Scan()
			deck = in.Text()
			deck = validatedeck(deck, rsnchoice)
		}
		analreason("mana", "l", deck, rsnchoice)
	case 6:
		println("Do you want to specify a deck?(y/n)")
		in.Scan()
		deck := in.Text()
		deck = validateuserinput(deck, "confirm")
		rsnchoice := historic()
		if deck == "y" {
			println("Deck: ")
			in.Scan()
			deck = in.Text()
			deck = validatedeck(deck, rsnchoice)
		}
		analreason("creature", "l", deck, rsnchoice)
	case 7:
		println("Unfinished")
		analreasonmenu()
	case 8:
		rsnchoice := historic()
		println("Deck:")
		in.Scan()
		deck := in.Text()
		deck = validatedeck(deck, rsnchoice)
		analreason("", "l", deck, rsnchoice)
	case 9:
		println("Do you want to specify a deck?(y/n)")
		in.Scan()
		deck := in.Text()
		deck = validateuserinput(deck, "confirm")
		rsnchoice := historic()
		if deck == "y" {
			println("Deck: ")
			in.Scan()
			deck = in.Text()
			deck = validatedeck(deck, rsnchoice)
		}
		println("Custom Filter Keyword:")
		in.Scan()
		custom := in.Text()
		println("Wins, loses or all?(w/l/a)")
		in.Scan()
		wl := in.Text()
		analreason(custom, wl, deck, rsnchoice)
	case 10:
		anal_menu()
	case 11:
		main()
	case 12:
		os.Exit(0)
	default:
		analreasonmenu()
	}
}
func analtimemenu() {
	//analysis time of day menu
	fmt.Println("Reason")
	m := make(map[string]string)

	// Set key/value pairs using typical `name[key] = val`
	m["k1"] = fmt.Sprintf("%-50s", "Wins Between Midnight and 6am")
	m["k2"] = fmt.Sprintf("%-50s", "Wins Between 6am and Noon")
	m["k3"] = fmt.Sprintf("%-50s", "Wins Between Noon and 6pm")
	m["k4"] = fmt.Sprintf("%-50s", "Wins Between 6pm and Midnight")
	m["k5"] = fmt.Sprintf("%0s", "Loses Between Midnight and 6am")
	m["k6"] = fmt.Sprintf("%0s", "Loses Between 6am and Noon")
	m["k7"] = fmt.Sprintf("%0s", "Loses Between Noon and 6pm")
	m["k8"] = fmt.Sprintf("%24s", "Losese Between 6pm and Midnight")
	m["k9"] = fmt.Sprintf("%0s", "Customized Start/End Time")
	m["k10"] = fmt.Sprintf("%-50s", "Return to Previous Menu")
	m["k11"] = fmt.Sprintf("%0s", "Return to Main Menu")
	m["k12"] = fmt.Sprintf("%0s", "Quit")

	// print menu options
	fmt.Println(" 1:", m["k1"]+"5:", m["k5"])
	fmt.Println(" 2:", m["k2"]+"6:", m["k6"])
	fmt.Println(" 3:", m["k3"]+"7:", m["k7"])
	fmt.Println(" 4:", m["k4"]+"8:", m["k8"])
	fmt.Println("                            "+" 9:", m["k9"])
	fmt.Println("10:", m["k10"]+"11:", m["k11"])
	fmt.Println("12:", m["k12"])

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	choice, _ := strconv.Atoi(in.Text())

	switch choice {
	case 1:
		tmchoice := historic()
		analtime("midnight", "w", tmchoice)
	case 2:
		tmchoice := historic()
		analtime("morning", "w", tmchoice)
	case 3:
		tmchoice := historic()
		analtime("noon", "w", tmchoice)
	case 4:
		tmchoice := historic()
		analtime("night", "w", tmchoice)
	case 5:
		tmchoice := historic()
		analtime("midnight", "l", tmchoice)
	case 6:
		tmchoice := historic()
		analtime("morning", "l", tmchoice)
	case 7:
		tmchoice := historic()
		analtime("noon", "l", tmchoice)
	case 8:
		tmchoice := historic()
		analtime("night", "l", tmchoice)
	case 9:
		tmchoice := historic()
		analtime("custom", "", tmchoice)
	case 10:
		anal_menu()
	case 11:
		main()
	case 12:
		os.Exit(0)
	default:
		analtimemenu()
	}
}
func analvltiermenu() {
	//analysis main menu
	fmt.Println("Analysis Level/Tier")
	m := make(map[string]string)
	//reader := bufio.NewReader(os.Stdin)

	// Set key/value pairs using typical `name[key] = val`
	m["k1"] = fmt.Sprintf("%-25s", "Wins by Level")
	m["k2"] = fmt.Sprintf("%-25s", "Loses by Level")
	m["k3"] = fmt.Sprintf("%-25s", "Wins by Specific Level-Tier")
	m["k4"] = fmt.Sprintf("%-25s", "Loses by Specific Level-Tier")
	m["k10"] = fmt.Sprintf("%0s", "Return to Previous Menu")
	m["k11"] = fmt.Sprintf("%0s", "Return to Main Menu")
	m["k12"] = fmt.Sprintf("%0s", "Quit")

	// print menu options
	fmt.Println(" 1:", m["k1"])
	fmt.Println(" 2:", m["k2"])
	fmt.Println(" 3:", m["k3"])
	fmt.Println(" 4:", m["k4"])
	fmt.Println("10:", m["k10"])
	fmt.Println("11:", m["k11"])
	fmt.Println("12:", m["k12"])

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	choice, _ := strconv.Atoi(in.Text())

	switch choice {
	case 1:
		println("Level: ")
		in.Scan()
		lvl := in.Text()
		lvl = validateuserinput(lvl, "level")
		analvltier(lvl, "", "w")
	case 2:
		println("Level: ")
		in.Scan()
		lvl := in.Text()
		lvl = validateuserinput(lvl, "level")
		analvltier(lvl, "", "l")
	case 3:
		println("Level: ")
		in.Scan()
		lvl := in.Text()
		lvl = validateuserinput(lvl, "level")
		println("Tier: ")
		in.Scan()
		tier := in.Text()
		tier = validateuserinput(tier, "tier")
		analvltier(lvl, tier, "w")
	case 4:
		println("Level: ")
		in.Scan()
		lvl := in.Text()
		lvl = validateuserinput(lvl, "level")
		println("Tier: ")
		in.Scan()
		tier := in.Text()
		tier = validateuserinput(tier, "tier")
		analvltier(lvl, tier, "l")
	case 10:
		anal_menu()
	case 11:
		main()
	case 12:
		os.Exit(0)
	default:
		analvltiermenu()
	}
}
func gamebydaymenu() {
	m := make(map[string]string)
	fmt.Println("Pick Best Day, Worst Day or Specific Day")
	m["k1"] = fmt.Sprintf("%-25s", "Best Day")
	m["k2"] = fmt.Sprintf("%-25s", "Worst Day")
	m["k3"] = fmt.Sprintf("%-25s", "Monday")
	m["k4"] = fmt.Sprintf("%-25s", "Tuesday")
	m["k5"] = fmt.Sprintf("%0s", "Wednesday")
	m["k6"] = fmt.Sprintf("%0s", "Thursday")
	m["k7"] = fmt.Sprintf("%0s", "Friday")
	m["k8"] = fmt.Sprintf("%0s", "Saturday")
	m["k9"] = fmt.Sprintf("%0s", "Sunday")
	m["k10"] = fmt.Sprintf("%0s", "Return to Previous Menu")
	m["k11"] = fmt.Sprintf("%0s", "Return to Main Menu")
	m["k12"] = fmt.Sprintf("%0s", "Quit")

	// print menu options
	fmt.Println(" 1:", m["k1"]+" 5:", m["k5"])
	fmt.Println(" 2:", m["k2"]+" 6:", m["k6"])
	fmt.Println(" 3:", m["k3"]+" 7:", m["k7"])
	fmt.Println(" 4:", m["k4"]+" 8:", m["k8"])
	fmt.Println("             9:", m["k9"])
	fmt.Println("10:", m["k10"])
	fmt.Println("11:", m["k11"])
	fmt.Println("12:", m["k12"])

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	choice, _ := strconv.Atoi(in.Text())

	switch choice {
	case 1:
		println("Best Day")
		gbdchoice := historic()
		println("Would you like to specify a deck?(y/n)")
		in.Scan()
		deckchoice := in.Text()
		deckchoice = validateuserinput(deckchoice, "confirm")
		if deckchoice == "y" {
			println("Deck: ")
			in.Scan()
			deck := in.Text()
			deck = validatedeck(deck, gbdchoice)
			println("Deck test: " + deck)
			analday(deck, "win", gbdchoice)
		} else if deckchoice == "n" {
			analday("n", "win", gbdchoice)
		}
	case 2:
		println("Worst Day")
		gbdchoice := historic()
		println("Would you like to specify a deck?(y/n)")
		in.Scan()
		deckchoice := in.Text()
		deckchoice = validateuserinput(deckchoice, "confirm")
		if deckchoice == "y" {
			println("Deck: ")
			in.Scan()
			deck := in.Text()
			deck = validatedeck(deck, gbdchoice)
			analday(deck, "lose", gbdchoice)
		} else if deckchoice == "n" {
			analday("n", "lose", gbdchoice)
		}
	case 3:
		println("Monday")
		analdmenu("Monday")
	case 4:
		println("Tuesday")
		analdmenu("Tuesday")
	case 5:
		println("Wednesday")
		analdmenu("Wednesday")
	case 6:
		println("Thursday")
		analdmenu("Thursday")
	case 7:
		println("Friday")
		analdmenu("Friday")
	case 8:
		println("Saturday")
		analdmenu("Saturday")
	case 9:
		println("Sunday")
		analdmenu("Sunday")
	case 10:
		anal_menu()
	case 11:
		main()
	case 12:
		os.Exit(0)
	default:
		gamebydaymenu()
	}
}
func deckbycardmenu(cardtype string) {
	m := make(map[string]string)
	fmt.Println("Choose Options Below")
	m["k1"] = fmt.Sprintf("%-30s", "Most wins by total "+cardtype+" number")
	m["k2"] = fmt.Sprintf("%-30s", "Most loses by total "+cardtype+" number")
	m["k3"] = fmt.Sprintf("%-30s", "Wins for all decks by total "+cardtype+" number")
	m["k4"] = fmt.Sprintf("%-25s", "Loses for all decks by total "+cardtype+" number")
	m["k10"] = fmt.Sprintf("%0s", "Return to Previous Menu")
	m["k11"] = fmt.Sprintf("%0s", "Return to Main Menu")
	m["k12"] = fmt.Sprintf("%0s", "Quit")

	// print menu options
	fmt.Println(" 1:", m["k1"])
	fmt.Println(" 2:", m["k2"])
	fmt.Println(" 3:", m["k3"])
	fmt.Println(" 4:", m["k4"])
	println("")
	fmt.Println("10:", m["k10"])
	fmt.Println("11:", m["k11"])
	fmt.Println("12:", m["k12"])

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	choice, _ := strconv.Atoi(in.Text())

	var cardchoice string
	switch cardtype {
	case "card":
		cardchoice = "numcards"
	case "creature":
		cardchoice = "numcreatures"
	case "non-creature":
		cardchoice = "numspells"
	case "land":
		cardchoice = "numlands"
	}

	switch choice {
	case 1:
		analdeckcard("wins", "", cardchoice)
	case 2:
		analdeckcard("loses", "", cardchoice)
	case 3:
		analdeckcard("wins", "all", cardchoice)
	case 4:
		analdeckcard("loses", "all", cardchoice)
	case 10:
		anal_menu()
	case 11:
		main()
	case 12:
		os.Exit(0)
	default:
		deckbycardmenu(cardtype)
	}
}
func analday(d string, win_lose string, h string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	var (
		deckname           string
		max_win            int
		max_lose           int
		day                string
		win_day_query      string
		win_day_all_query  string
		lose_day_query     string
		lose_day_all_query string
	)

	if h == "y" {
		win_day_query = "SELECT deck, MAX(win_count) as max_win, day_of_week FROM mtga.wins_by_day WHERE deck=? GROUP BY deck, day_of_week ORDER BY win_count DESC LIMIT 1"
		win_day_all_query = "SELECT deck, win_count, day_of_week FROM mtga.most_wbd ORDER BY FIELD(day_of_week , 'MONDAY', 'TUESDAY', 'WEDNESDAY', 'THURSDAY', 'FRIDAY', 'SATURDAY', 'SUNDAY'), win_count DESC;"
		lose_day_query = "SELECT deck, MAX(lose_count) as max_loses, day_of_week FROM mtga.loses_by_day WHERE deck=? GROUP BY deck, day_of_week ORDER BY lose_count DESC LIMIT 1"
		lose_day_all_query = "SELECT deck, lose_count, day_of_week FROM mtga.most_lbd ORDER BY FIELD(day_of_week , 'MONDAY', 'TUESDAY', 'WEDNESDAY', 'THURSDAY', 'FRIDAY', 'SATURDAY', 'SUNDAY'), lose_count DESC"
	} else {
		win_day_query = "SELECT deck, MAX(win_count) as max_win, day_of_week FROM mtga.wins_by_day WHERE deck=? AND deck IN (SELECT name FROM mtga.decks) GROUP BY deck, day_of_week ORDER BY win_count DESC LIMIT 1"
		win_day_all_query = "SELECT deck, win_count, day_of_week FROM mtga.most_wbd WHERE deck IN (SELECT name FROM mtga.decks) ORDER BY FIELD(day_of_week , 'MONDAY', 'TUESDAY', 'WEDNESDAY', 'THURSDAY', 'FRIDAY', 'SATURDAY', 'SUNDAY'), win_count DESC;"
		lose_day_query = "SELECT deck, MAX(lose_count) as max_loses, day_of_week FROM mtga.loses_by_day WHERE deck=? AND deck IN (SELECT name FROM mtga.decks) GROUP BY deck, day_of_week ORDER BY lose_count DESC LIMIT 1"
		lose_day_all_query = "SELECT deck, lose_count, day_of_week FROM mtga.most_lbd WHERE deck IN (SELECT name FROM mtga.decks) ORDER BY FIELD(day_of_week , 'MONDAY', 'TUESDAY', 'WEDNESDAY', 'THURSDAY', 'FRIDAY', 'SATURDAY', 'SUNDAY'), lose_count DESC"
	}

	if d != "n" && win_lose == "win" {
		results := db.QueryRow(win_day_query, d)
		err := results.Scan(&deckname, &max_win, &day)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				gamebydaymenu()
			} else {
				panic(err.Error())
			}
		}
		fdeck := fmt.Sprintf("%-30s", "Deck: "+deckname)
		fday := fmt.Sprintf("%-25s", "Most Wins: "+day)
		finalstring := fmt.Sprint(fdeck + fday + " Wins: " + strconv.Itoa(max_win))
		fmt.Println(finalstring)
		println("")
		gamebydaymenu()
	} else if d == "n" && win_lose == "win" {
		results, err := db.Query(win_day_all_query)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		for results.Next() {
			// for each row, scan the result into our deck composite object
			err = results.Scan(&deckname, &max_win, &day)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)
			fdeck := fmt.Sprintf("%-30s", "Deck: "+deckname)
			fday := fmt.Sprintf("%-25s", "Most Wins: "+day)
			finalstring := fmt.Sprint(fdeck + fday + " Wins: " + strconv.Itoa(max_win))
			fmt.Println(finalstring)
		}
		gamebydaymenu()
	} else if d != "n" && win_lose == "lose" {
		results := db.QueryRow(lose_day_query, d)
		err := results.Scan(&deckname, &max_lose, &day)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				gamebydaymenu()
			} else {
				panic(err.Error())
			}
		}
		fdeck := fmt.Sprintf("%-30s", "Deck: "+deckname)
		fday := fmt.Sprintf("%-25s", "Most Loses: "+day)
		finalstring := fmt.Sprint(fdeck + fday + " Loses: " + strconv.Itoa(max_lose))
		fmt.Println(finalstring)
		println("")
		gamebydaymenu()
	} else if d == "n" && win_lose == "lose" {
		results, err := db.Query(lose_day_all_query)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		for results.Next() {
			// for each row, scan the result into our deck composite object
			err = results.Scan(&deckname, &max_lose, &day)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)
			fdeck := fmt.Sprintf("%-30s", "Deck: "+deckname)
			fday := fmt.Sprintf("%-25s", "Most Loses: "+day)
			finalstring := fmt.Sprint(fdeck + fday + " Loses: " + strconv.Itoa(max_lose))
			fmt.Println(finalstring)
		}
		gamebydaymenu()
	}
}
func analweekday(d string, day string, wl string, h string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	var (
		deckname       string
		wl_count       int
		wkdy_best      string
		wkdy_worst     string
		wkdy_all_best  string
		wkdy_all_worst string
	)

	if h == "y" {
		wkdy_best = "SELECT deck, win_count FROM mtga.wins_by_day WHERE day_of_week =? ORDER BY win_count DESC LIMIT 1"
		wkdy_worst = "SELECT deck, lose_count FROM mtga.loses_by_day WHERE day_of_week =? ORDER BY lose_count DESC LIMIT 1"
		wkdy_all_best = "SELECT deck, win_count FROM mtga.wins_by_day WHERE day_of_week =? ORDER BY win_count"
		wkdy_all_worst = "SELECT deck, lose_count FROM mtga.loses_by_day WHERE day_of_week =? ORDER BY lose_count"
	} else {
		wkdy_best = "SELECT deck, win_count FROM mtga.wins_by_day WHERE day_of_week =? AND deck IN (SELECT name FROM mtga.decks) ORDER BY win_count DESC LIMIT 1"
		wkdy_worst = "SELECT deck, lose_count FROM mtga.loses_by_day WHERE day_of_week =? AND deck IN (SELECT name FROM mtga.decks) ORDER BY lose_count DESC LIMIT 1"
		wkdy_all_best = "SELECT deck, win_count FROM mtga.wins_by_day WHERE day_of_week =? AND deck IN (SELECT name FROM mtga.decks) ORDER BY win_count"
		wkdy_all_worst = "SELECT deck, lose_count FROM mtga.loses_by_day WHERE day_of_week =? AND deck IN (SELECT name FROM mtga.decks) ORDER BY lose_count"
	}

	if d == "best" {
		results := db.QueryRow(wkdy_best, day)
		err := results.Scan(&deckname, &wl_count)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				analdmenu(day)
			} else {
				panic(err.Error())
			}
		}
		fday := fmt.Sprintf("%-20s", "Day: "+day)
		fdeck := fmt.Sprintf("%-30s", "Deck: "+deckname)
		finalstring := fmt.Sprint(fday + fdeck + " Wins: " + strconv.Itoa(wl_count))
		fmt.Println(finalstring)
		println("")
		analdmenu(day)
	} else if d == "worst" {
		results := db.QueryRow(wkdy_worst, day)
		err := results.Scan(&deckname, &wl_count)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				analdmenu(day)
			} else {
				panic(err.Error())
			}
		}
		fday := fmt.Sprintf("%-20s", "Day: "+day)
		fdeck := fmt.Sprintf("%-30s", "Deck: "+deckname)
		finalstring := fmt.Sprint(fday + fdeck + " Loses: " + strconv.Itoa(wl_count))
		fmt.Println(finalstring)
		println("")
		analdmenu(day)
	} else if wl == "w" && d != "all" {
		results := db.QueryRow("SELECT win_count FROM mtga.wins_by_day WHERE day_of_week =? AND deck =?", day, d)
		err := results.Scan(&wl_count)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Wins Recored for this Deck and this Day")
				fmt.Println("")
				analdmenu(day)
			} else {
				panic(err.Error())
			}
		}
		fday := fmt.Sprintf("%-20s", "Day: "+day)
		fdeck := fmt.Sprintf("%-30s", "Deck: "+d)
		finalstring := fmt.Sprint(fday + fdeck + "Wins: " + strconv.Itoa(wl_count))
		fmt.Println(finalstring)
		println("")
		analdmenu(day)
	} else if wl == "l" && d != "all" {
		results := db.QueryRow("SELECT lose_count FROM mtga.loses_by_day WHERE day_of_week =? AND deck =?", day, d)
		err := results.Scan(&wl_count)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Loses Recored for this Deck and this Day")
				fmt.Println("")
				analdmenu(day)
			} else {
				panic(err.Error())
			}
		}
		fday := fmt.Sprintf("%-20s", "Day: "+day)
		fdeck := fmt.Sprintf("%-30s", "Deck: "+d)
		finalstring := fmt.Sprint(fday + fdeck + "Loses: " + strconv.Itoa(wl_count))
		fmt.Println(finalstring)
		println("")
		analdmenu(day)
	} else if wl == "w" && d == "all" {
		results, err := db.Query(wkdy_all_best, day)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Wins Recored for this Day")
				fmt.Println("")
				analdmenu(day)
			} else {
				panic(err.Error())
			}
		}

		for results.Next() {
			// for each row, scan the result into our deck composite object
			err = results.Scan(&deckname, &wl_count)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)
			fday := fmt.Sprintf("%-20s", "Day: "+day)
			fdeck := fmt.Sprintf("%-30s", "Deck: "+deckname)
			finalstring := fmt.Sprint(fdeck + fday + "Wins: " + strconv.Itoa(wl_count))
			fmt.Println(finalstring)
		}
		println("")
		analdmenu(day)
	} else if wl == "l" && d == "all" {
		results, err := db.Query(wkdy_all_worst, day)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Loses Recored for this Day")
				fmt.Println("")
				analdmenu(day)
			} else {
				panic(err.Error())
			}
		}

		for results.Next() {
			// for each row, scan the result into our deck composite object
			err = results.Scan(&deckname, &wl_count)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)
			fday := fmt.Sprintf("%-20s", "Day: "+day)
			fdeck := fmt.Sprintf("%-30s", "Deck: "+deckname)
			finalstring := fmt.Sprint(fdeck + fday + "Loses: " + strconv.Itoa(wl_count))
			fmt.Println(finalstring)
		}
		println("")
		analdmenu(day)
	}
}
func analreason(s string, wl string, d string, h string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	defer db.Close()

	var (
		cause         string
		deck          string
		results       int
		rsn_query     string
		rsn_all_query string
	)

	if wl == "w" {
		results = 0
	} else if wl == "l" {
		results = 1
	}

	if h == "y" {
		rsn_query = "SELECT deck, cause FROM mtga.games WHERE cause LIKE CONCAT('%', ?, '%') AND results LIKE CONCAT('%', ?, '%') ORDER BY deck"
		rsn_all_query = "SELECT deck, cause FROM mtga.games WHERE cause LIKE CONCAT('%', ?, '%') ORDER BY deck"
	} else {
		rsn_query = "SELECT deck, cause FROM mtga.games WHERE cause LIKE CONCAT('%', ?, '%') AND results LIKE CONCAT('%', ?, '%') AND deck IN (SELECT name FROM mtga.decks) ORDER BY deck"
		rsn_all_query = "SELECT deck, cause FROM mtga.games WHERE cause LIKE CONCAT('%', ?, '%') AND deck IN (SELECT name FROM mtga.decks) ORDER BY deck"
	}

	if d != "n" && wl != "a" {
		results, err := db.Query("SELECT cause FROM mtga.games WHERE cause LIKE CONCAT('%', ?, '%') AND deck =? AND results LIKE CONCAT('%', ?, '%')", s, d, results)
		//err := results.Scan(&cause)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				analreasonmenu()
			} else {
				panic(err.Error())
			}
		}
		for results.Next() {
			// for each row, scan the result into our deck composite object
			err = results.Scan(&cause)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)
			fdeck := fmt.Sprintf("%-30s", "Deck: "+d)
			finalstring := fmt.Sprint(fdeck + " Reasons: " + cause)
			fmt.Println(finalstring)
		}
		println("")
		analreasonmenu()
	} else if d == "n" && wl != "a" {
		results, err := db.Query(rsn_query, s, results)

		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				analreasonmenu()
			} else {
				panic(err.Error())
			}
		}
		for results.Next() {
			// for each row, scan the result into our deck composite object
			err = results.Scan(&deck, &cause)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)
			fdeck := fmt.Sprintf("%-30s", "Deck: "+deck)
			finalstring := fmt.Sprint(fdeck + " Reasons: " + cause)
			fmt.Println(finalstring)
		}
		println("")
		analreasonmenu()
	} else if d != "n" && wl == "a" {
		results, err := db.Query("SELECT cause FROM mtga.games WHERE cause LIKE CONCAT('%', ?, '%') AND deck =?", s, d)
		//err := results.Scan(&cause)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				analreasonmenu()
			} else {
				panic(err.Error())
			}
		}
		for results.Next() {
			// for each row, scan the result into our deck composite object
			err = results.Scan(&cause)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)
			fdeck := fmt.Sprintf("%-30s", "Deck: "+deck)
			finalstring := fmt.Sprint(fdeck + " Reasons: " + cause)
			fmt.Println(finalstring)
		}
		println("")
		analreasonmenu()
	} else if d == "n" && wl == "a" {
		results, err := db.Query(rsn_all_query, s)

		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				analreasonmenu()
			} else {
				panic(err.Error())
			}
		}
		for results.Next() {
			// for each row, scan the result into our deck composite object
			err = results.Scan(&deck, &cause)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)
			fdeck := fmt.Sprintf("%-30s", "Deck: "+deck)
			finalstring := fmt.Sprint(fdeck + " Reasons: " + cause)
			fmt.Println(finalstring)
		}
		println("")
		analreasonmenu()
	}
}
func analtime(t string, wl string, h string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	defer db.Close()

	var (
		rcount   int
		deck     string
		cause    string
		hour     string
		s        string
		e        string
		iwl      int
		tm_query string
	)

	in := bufio.NewScanner(os.Stdin)
	println("Do you want to specify a deck?(y/n)")
	in.Scan()
	confirmchoice := in.Text()
	confirmchoice = validateuserinput(confirmchoice, "confirm")

	if t == "midnight" {
		s = "00:00:00"
		e = "06:00:00"
	} else if t == "morning" {
		s = "06:00:00"
		e = "12:00:00"
	} else if t == "noon" {
		s = "12:00:00"
		e = "18:00:00"
	} else if t == "night" {
		s = "18:00:00"
		e = "23:00:59"
	} else if t == "custom" {
		println("Custom")
		println("Won or Lost")
		in.Scan()
		wl = in.Text()
		wl = validateuserinput(wl, "results")
		runes := []rune(wl)
		// ... Convert back into a string from rune slice.
		wl = string(runes[0:1])
		println("Starting Time(Format: 00:00:00)")
		in.Scan()
		s = in.Text()
		println("Ending Time(Format: 00:00:00)")
		in.Scan()
		e = in.Text()
	}

	if wl == "w" {
		iwl = 0
	} else if wl == "l" {
		iwl = 1
	}

	if h == "y" {
		tm_query = "SELECT deck, cause, TIME(`Timestamp`) AS playtime FROM mtga.games WHERE (TIME(`Timestamp`) BETWEEN ? AND ?) AND results =? ORDER BY deck"
	} else {
		tm_query = "SELECT deck, cause, TIME(`Timestamp`) AS playtime FROM mtga.games WHERE (TIME(`Timestamp`) BETWEEN ? AND ?) AND results =? AND deck IN (SELECT name FROM mtga.decks) ORDER BY deck"
	}

	if confirmchoice == "n" {
		results, err := db.Query(tm_query, s, e, iwl)

		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				analtimemenu()
			} else {
				panic(err.Error())
			}
		}
		for results.Next() {
			// for each row, scan the result into our deck composite object
			err = results.Scan(&deck, &cause, &hour)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)

			layout1 := "03:04:05 PM"
			layout2 := "15:04:05"
			t, err := time.Parse(layout2, hour)
			if err != nil {
				fmt.Println(err)
				return
			}
			fdeck := fmt.Sprintf("%-30s", "Deck: "+deck)
			fhour := fmt.Sprintf("%-25s", "Hour: "+t.Format(layout1))
			finalstring := fmt.Sprint(fdeck + fhour + " Reasons: " + cause)
			rcount++
			fmt.Println(finalstring)
		}
		println("Total Row Count: " + strconv.Itoa(rcount))
		println("")
		analtimemenu()
	} else if confirmchoice == "y" {
		println("Deck:")
		in.Scan()
		deckchoice := in.Text()
		deckchoice = validatedeck(deckchoice, h)

		results, err := db.Query("SELECT deck, cause, TIME(`Timestamp`) AS playtime FROM mtga.games WHERE (TIME(`Timestamp`) BETWEEN ? AND ?) AND results =? AND deck =?", s, e, iwl, deckchoice)

		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				analtimemenu()
			} else {
				panic(err.Error())
			}
		}
		for results.Next() {
			// for each row, scan the result into our deck composite object
			err = results.Scan(&deck, &cause, &hour)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)

			layout1 := "03:04:05 PM"
			layout2 := "15:04:05"
			t, err := time.Parse(layout2, hour)
			if err != nil {
				fmt.Println(err)
				return
			}
			fdeck := fmt.Sprintf("%-30s", "Deck: "+deck)
			fhour := fmt.Sprintf("%-25s", "Hour: "+t.Format(layout1))
			finalstring := fmt.Sprint(fdeck + fhour + " Reasons: " + cause)
			rcount++
			fmt.Println(finalstring)
		}
		println("Total Row Count: " + strconv.Itoa(rcount))
		println("")
		analtimemenu()
	}
}
func analvltier(lvl string, tr string, wl string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	defer db.Close()

	var (
		deck  string
		level string
		iwl   int
		opp   string
		cause string
	)

	in := bufio.NewScanner(os.Stdin)
	println("Do you want to specify a deck?(y/n)")
	in.Scan()
	confirmchoice := in.Text()
	confirmchoice = validateuserinput(confirmchoice, "confirm")

	if wl == "w" {
		iwl = 0
	} else if wl == "l" {
		iwl = 1
	}

	if confirmchoice == "n" {
		results, err := db.Query("SELECT deck, opponent, `level`, cause FROM mtga.games WHERE results=? AND level LIKE CONCAT('%',?,'%',?,'%') ORDER BY deck",
			iwl, lvl, tr)
		//err := results.Scan(&cause)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				analreasonmenu()
			} else {
				panic(err.Error())
			}
		}
		for results.Next() {
			// for each row, scan the result into our deck composite object
			err = results.Scan(&deck, &opp, &level, &cause)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)
			fdeck := fmt.Sprintf("%-30s", "Deck: "+deck)
			fopp := fmt.Sprintf("%-35s", "Opponent: "+opp)
			finalstring := fmt.Sprint(fdeck + fopp + " Level: " + level + " Reasons: " + cause)
			fmt.Println(finalstring)
		}
		println("")
		analvltiermenu()
	} else if confirmchoice != "n" {
		println("Deck:")
		in.Scan()
		deckchoice := in.Text()
		deckchoice = validatedeck(deckchoice, "n")

		results, err := db.Query("SELECT deck, opponent, `level`, cause FROM mtga.games WHERE results=? AND level LIKE CONCAT('%',?,'%',?,'%') AND deck =? ORDER BY deck ", iwl, lvl, tr, deckchoice)

		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				analreasonmenu()
			} else {
				panic(err.Error())
			}
		}
		for results.Next() {
			// for each row, scan the result into our deck composite object
			err = results.Scan(&deck, &opp, &level, &cause)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)
			fdeck := fmt.Sprintf("%-30s", "Deck: "+deck)
			fopp := fmt.Sprintf("%-35s", "Opponent: "+opp)
			finalstring := fmt.Sprint(fdeck + fopp + " Level: " + level + " Reasons: " + cause)
			fmt.Println(finalstring)
		}
		println("")
		analvltiermenu()
	}
}
func recdelete() {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	defer db.Close()

	var (
		deck         string
		date_entered time.Time
		win_pct      float32
		win_count    int
		games        int
	)

	results, err := db.Query("SELECT name, date_entered, win_pct, win_count, games FROM mtga.decks d JOIN mtga.win_percentage wp ON d.name = wp.deck WHERE win_pct <= .20")

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			fmt.Println("No Games Recored for this Deck")
			fmt.Println("")
			anal_menu()
		} else {
			panic(err.Error())
		}
	}
	for results.Next() {
		// for each row, scan the result into our deck composite object
		err = results.Scan(&deck, &date_entered, &win_pct, &win_count, &games)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.SetFlags(0)
		fdeck := fmt.Sprintf("%-30s", "Deck: "+deck)
		fdate := fmt.Sprintf("%-30s", "Date Entered: "+date_entered.Format("2006-01-02"))
		cwin_pct := fmt.Sprintf("%f", win_pct)
		cwin_pct = cwin_pct[2:4]
		fwin_pct := fmt.Sprintf("%-20s", "Win Percent: "+cwin_pct+"%")
		fwin_ct := fmt.Sprintf("%-15s", "Win Count: "+strconv.Itoa(win_count))
		finalstring := fmt.Sprint(fdeck + fdate + fwin_pct + fwin_ct + "Number of Games: " + strconv.Itoa(games))
		fmt.Println(finalstring)
	}
	println("")
	anal_menu()
}
func analdeckcard(wl string, all string, card string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	defer db.Close()

	var (
		cardtype    int
		winlose     int
		finalstring string
		stmt        string
	)

	if all == "" {
		stmt = "SELECT " + card + ", " + wl + " FROM mtga.decks d JOIN mtga.record r ON d.name = r.deck GROUP BY " + card + ", " + wl + " ORDER BY " + wl + " DESC, " + card + " DESC LIMIT 10"
	} else if all == "all" {
		stmt = "SELECT " + card + ", " + wl + " FROM mtga.decks d JOIN mtga.record r ON d.name = r.deck ORDER BY " + wl + " DESC, " + card + " DESC"
	}

	results, err := db.Query(stmt)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			fmt.Println("No Games Recored for this Deck")
			fmt.Println("")
			anal_menu()
		} else {
			panic(err.Error())
		}
	}
	for results.Next() {
		// for each row, scan the result into our deck composite object
		err = results.Scan(&cardtype, &winlose)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.SetFlags(0)
		switch card {
		case "numspells":
			card = "Total Non-Creatures: "
		case "numcards":
			card = "Total Cards: "
		case "numcreatures":
			card = "Total Creatures: "
		case "numlands":
			card = "Total Lands: "
		}
		switch wl {
		case "wins":
			wl = "Wins: "
		case "loses":
			wl = "Loses: "
		}
		fcardtype := fmt.Sprintf("%-30s", card+strconv.Itoa(cardtype))
		fwinlose := fmt.Sprintf("%0s", wl+strconv.Itoa(winlose))
		finalstring = fmt.Sprint(fcardtype + fwinlose)
		fmt.Println(finalstring)
	}
	println()
	switch card {
	case "Total Non-Creatures: ":
		deckbycardmenu("non-creature")
	case "Total Cards: ":
		deckbycardmenu("card")
	case "Total Creatures: ":
		deckbycardmenu("creature")
	case "Total Lands: ":
		deckbycardmenu("land")
	}
}
