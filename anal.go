package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	}
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
	m["k10"] = fmt.Sprintf("%-24s", "Return to Main Menu")
	m["k11"] = fmt.Sprintf("%0s", "Quit")

	// print menu options
	fmt.Println(" 1:", m["k1"]+"5:", m["k5"])
	fmt.Println(" 2:", m["k2"]+"6:", m["k6"])
	fmt.Println(" 3:", m["k3"]+"7:", m["k7"])
	fmt.Println(" 4:", m["k4"]+"8:", m["k8"])
	fmt.Println("              "+" 9:", m["k9"])
	fmt.Println("10:", m["k10"]+"11:", m["k11"])

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
