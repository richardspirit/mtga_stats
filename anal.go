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
		gamebyday()
	case 2:
		analreasonmenu()
	case 3:
		analtimemenu()
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
		analweekday("best", day, "")
	case 2:
		analweekday("worst", day, "")
	case 3:
		println("Deck:")
		in.Scan()
		deck := in.Text()
		deck = validatedeck(deck)
		analweekday(deck, day, "w")
	case 4:
		println("Deck:")
		in.Scan()
		deck := in.Text()
		deck = validatedeck(deck)
		analweekday(deck, day, "l")
	case 5:
		analweekday("all", day, "w")
	case 6:
		analweekday("all", day, "l")
	case 10:
		gamebyday()
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
		if deck == "y" {
			println("Deck: ")
			in.Scan()
			deck = in.Text()
			deck = validatedeck(deck)
		}
		analreason("mana", "w", deck)
	case 2:
		println("Do you want to specify a deck?(y/n)")
		in.Scan()
		deck := in.Text()
		deck = validateuserinput(deck, "confirm")
		if deck == "y" {
			println("Deck: ")
			in.Scan()
			deck = in.Text()
			deck = validatedeck(deck)
		}
		analreason("creature", "w", deck)
	case 3:
		println("Unfinished")
		analreasonmenu()
	case 4:
		println("Deck:")
		in.Scan()
		deck := in.Text()
		deck = validatedeck(deck)
		analreason("", "w", deck)
	case 5:
		println("Do you want to specify a deck?(y/n)")
		in.Scan()
		deck := in.Text()
		deck = validateuserinput(deck, "confirm")
		if deck == "y" {
			println("Deck: ")
			in.Scan()
			deck = in.Text()
			deck = validatedeck(deck)
		}
		analreason("mana", "l", deck)
	case 6:
		println("Do you want to specify a deck?(y/n)")
		in.Scan()
		deck := in.Text()
		deck = validateuserinput(deck, "confirm")
		if deck == "y" {
			println("Deck: ")
			in.Scan()
			deck = in.Text()
			deck = validatedeck(deck)
		}
		analreason("creature", "l", deck)
	case 7:
		println("Unfinished")
		analreasonmenu()
	case 8:
		println("Deck:")
		in.Scan()
		deck := in.Text()
		deck = validatedeck(deck)
		analreason("", "l", deck)
	case 9:
		println("Do you want to specify a deck?(y/n)")
		in.Scan()
		deck := in.Text()
		deck = validateuserinput(deck, "confirm")
		if deck == "y" {
			println("Deck: ")
			in.Scan()
			deck = in.Text()
			deck = validatedeck(deck)
		}
		println("Custom Filter Keyword:")
		in.Scan()
		custom := in.Text()
		println("Wins, loses or all?(w/l/a")
		in.Scan()
		wl := in.Text()
		analreason(custom, wl, deck)
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
		analtime("midnight", "w")
	case 2:
		analtime("morning", "w")
	case 3:
		analtime("noon", "w")
	case 4:
		analtime("night", "w")
	case 5:
		analtime("midnight", "l")
	case 6:
		analtime("morning", "l")
	case 7:
		analtime("noon", "l")
	case 8:
		analtime("night", "l")
	case 9:
		analtime("custom", "")
	case 10:
		anal_menu()
	case 11:
		main()
	case 12:
		os.Exit(0)
	default:
		analtimemenu()
	}
	//analtime()
}
func gamebyday() {
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
		println("Would you like to specify a deck?(y/n)")
		in.Scan()
		deckchoice := in.Text()
		deckchoice = validateuserinput(deckchoice, "confirm")
		if deckchoice == "y" {
			println("Deck: ")
			in.Scan()
			deck := in.Text()
			deck = validatedeck(deck)
			analday(deck, "win")
		} else if deckchoice == "n" {
			analday("n", "win")
		}
	case 2:
		println("Worst Day")
		println("Would you like to specify a deck?(y/n)")
		in.Scan()
		deckchoice := in.Text()
		deckchoice = validateuserinput(deckchoice, "confirm")
		if deckchoice == "y" {
			println("Deck: ")
			in.Scan()
			deck := in.Text()
			deck = validatedeck(deck)
			analday(deck, "lose")
		} else if deckchoice == "n" {
			analday("n", "lose")
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
		gamebyday()
	}
}
func analday(d string, win_lose string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	var (
		deckname string
		max_win  int
		max_lose int
		day      string
	)

	if d != "n" && win_lose == "win" {
		results := db.QueryRow("SELECT deck, MAX(win_count) as max_win, day_of_week FROM mtga.wins_by_day WHERE deck=? GROUP BY deck, day_of_week order by win_count desc limit 1", d)
		err := results.Scan(&deckname, &max_win, &day)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				gamebyday()
			} else {
				panic(err.Error())
			}
		}
		finalstring := fmt.Sprint("The day of most wins for " + deckname + " is " + day + " with " + strconv.Itoa(max_win) + " wins")
		fmt.Println(finalstring)
		println("")
		gamebyday()
	} else if d == "n" && win_lose == "win" {
		results, err := db.Query("SELECT deck, win_count, day_of_week  FROM mtga.most_wbd")
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
			finalstring := fmt.Sprint("The day of most wins for " + deckname + " is " + day + " with " + strconv.Itoa(max_win) + " wins")
			fmt.Println(finalstring)
		}
		gamebyday()
	} else if d != "n" && win_lose == "lose" {
		results := db.QueryRow("SELECT deck, MAX(lose_count) as max_loses, day_of_week FROM mtga.loses_by_day WHERE deck=? group by deck, day_of_week order by lose_count desc limit 1", d)
		err := results.Scan(&deckname, &max_lose, &day)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				gamebyday()
			} else {
				panic(err.Error())
			}
		}
		finalstring := fmt.Sprint("The day of most loses for " + deckname + " is " + day + " with " + strconv.Itoa(max_lose) + " loses")
		fmt.Println(finalstring)
		println("")
		gamebyday()
	} else if d == "n" && win_lose == "lose" {
		results, err := db.Query("SELECT deck, lose_count, day_of_week FROM mtga.most_lbd")
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
			finalstring := fmt.Sprint("The day of most loses for " + deckname + " is " + day + " with " + strconv.Itoa(max_lose) + " loses")
			fmt.Println(finalstring)
		}
		gamebyday()
	}
}
func analweekday(d string, day string, wl string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	var (
		deckname string
		wl_count int
	)

	if d == "best" {
		results := db.QueryRow("SELECT deck, win_count FROM mtga.wins_by_day WHERE day_of_week =? ORDER BY win_count DESC LIMIT 1", day)
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
		finalstring := fmt.Sprint("The the best deck for  " + day + " is " + deckname + " with " + strconv.Itoa(wl_count) + " wins")
		fmt.Println(finalstring)
		println("")
		analdmenu(day)
	} else if d == "worst" {
		results := db.QueryRow("select deck, lose_count from mtga.loses_by_day where day_of_week =? order by lose_count desc limit 1", day)
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
		finalstring := fmt.Sprint("The the worst deck for  " + day + " is " + deckname + " with " + strconv.Itoa(wl_count) + " loses")
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
		finalstring := fmt.Sprint("Wins on " + day + " for " + d + " with " + strconv.Itoa(wl_count) + " wins")
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
		finalstring := fmt.Sprint("Loses on " + day + " for " + d + " with " + strconv.Itoa(wl_count) + " loses")
		fmt.Println(finalstring)
		println("")
		analdmenu(day)
	} else if wl == "w" && d == "all" {
		results, err := db.Query("SELECT deck, win_count FROM mtga.wins_by_day WHERE day_of_week =?", day)
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
			finalstring := fmt.Sprint("Wins for " + deckname + " on " + day + " with " + strconv.Itoa(wl_count) + " wins")
			fmt.Println(finalstring)
		}
		println("")
		analdmenu(day)
	} else if wl == "l" && d == "all" {
		results, err := db.Query("SELECT deck, lose_count FROM mtga.loses_by_day WHERE day_of_week =?", day)
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
			finalstring := fmt.Sprint("Loses for " + deckname + " on " + day + " with " + strconv.Itoa(wl_count) + " loses")
			fmt.Println(finalstring)
		}
		println("")
		analdmenu(day)
	}
}
func analreason(s string, wl string, d string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	defer db.Close()

	var (
		cause   string
		deck    string
		results int
	)

	if wl == "w" {
		results = 0
	} else if wl == "l" {
		results = 1
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
			finalstring := fmt.Sprint("Deck: " + d + " Reasons: " + cause)
			fmt.Println(finalstring)
		}
		println("")
		analreasonmenu()
	} else if d == "n" && wl != "a" {
		results, err := db.Query("SELECT deck, cause FROM mtga.games WHERE cause LIKE CONCAT('%', ?, '%') AND results LIKE CONCAT('%', ?, '%') ORDER BY deck", s, results)

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
			fdeck := fmt.Sprintf("%-50s", "Deck: "+deck)
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
			finalstring := fmt.Sprint("Deck: " + d + " Reasons: " + cause)
			fmt.Println(finalstring)
		}
		println("")
		analreasonmenu()
	} else if d == "n" && wl == "a" {
		results, err := db.Query("SELECT deck, cause FROM mtga.games WHERE cause LIKE CONCAT('%', ?, '%') ORDER BY deck", s)

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
			fdeck := fmt.Sprintf("%-50s", "Deck: "+deck)
			finalstring := fmt.Sprint(fdeck + " Reasons: " + cause)
			fmt.Println(finalstring)
		}
		println("")
		analreasonmenu()
	}
}
func analtime(t string, wl string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	defer db.Close()

	var (
		rcount int
		deck   string
		cause  string
		hour   string
		s      string
		e      string
		iwl    int
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
	}

	if wl == "w" {
		iwl = 0
	} else if wl == "l" {
		iwl = 1
	}

	if confirmchoice == "n" {
		results, err := db.Query("SELECT deck, cause, TIME(`Timestamp`) AS playtime FROM mtga.games WHERE (TIME(`Timestamp`) BETWEEN ? AND ?) AND results =? ORDER BY deck", s, e, iwl)

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
		deckchoice = validatedeck(deckchoice)

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
