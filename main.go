package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Deck struct {
	Name         string    `json:"name"`
	Colors       string    `json:"colors"`
	Date_Entered time.Time `json:"date_entered"`
	Favorite     int       `json:"favorite"`
	Max_Streak   int       `json:"max_streak"`
	Cur_Streak   int       `json:"cur_streak"`
	Num_Cards    int       `json:"num_cards"`
	Num_Lands    int       `json:"num_lands"`
	Num_Creat    int       `json:"num_creat"`
	Num_Spells   int       `json:"num_spells"`
	Disable      int       `json:"disable"`
}
type Game struct {
	Results       int    `json:"results"`
	Cause         string `json:"cause"`
	Deck          string `json:"deck"`
	Opponent      string `json:"opponent"`
	Level         string `json:"level"`
	CurrentStreak int    `json:"currentstreak"`
	MaxStreak     int    `json:"maxstreak"`
}
type Records struct {
	Deck  string `json:"deck"`
	Wins  int    `json:"wins"`
	Loses int    `json:"loses"`
}

func main() {
	menu()
}
func menu() {
	//main menu
	fmt.Println("mtga Stats")
	m := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	// Set key/value pairs using typical `name[key] = val`
	m["k1"] = fmt.Sprintf("%-25s", "Enter/Delete New Deck")
	m["k2"] = fmt.Sprintf("%0s", "Add New Game")
	m["k3"] = fmt.Sprintf("%-25s", "View Deck Records")
	m["k4"] = fmt.Sprintf("%0s", "View Game Count")
	m["k5"] = fmt.Sprintf("%-25s", "View Decks")
	m["k6"] = fmt.Sprintf("%0s", "Top Ten Decks")
	m["k7"] = fmt.Sprintf("%-25s", "Edit Deck")
	m["k8"] = fmt.Sprintf("%0s", "Win/Lose Percent")
	m["k9"] = fmt.Sprintf("%-25s", "Analysis")
	m["k10"] = fmt.Sprintf("%0s", "Quit")

	// print menu options
	fmt.Println("1:", m["k1"]+" 2:", m["k2"])
	fmt.Println("3:", m["k3"]+" 4:", m["k4"])
	fmt.Println("5:", m["k5"]+" 6:", m["k6"])
	fmt.Println("7:", m["k7"]+" 8:", m["k8"])
	fmt.Println("9:", m["k9"]+"10:", m["k10"])

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	choice, _ := strconv.Atoi(in.Text())

	switch choice {
	case 1:
		//reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter or Delete Deck: ")
		edchoice, _ := reader.ReadString('\n')
		edchoice = strings.TrimSuffix(edchoice, "\r\n")
		//validate user input
		edchoice = validateuserinput(edchoice, "choice")
		if edchoice == "delete" || edchoice == "Delete" {
			fmt.Println("Delete Deck: ")
			ddeck, _ := reader.ReadString('\n')
			ddeck = strings.TrimSuffix(ddeck, "\r\n")
			fmt.Println("Delete Deck: " + ddeck)
			viewdecks(ddeck, 1)
			fmt.Println("Confirm(y/n): ")
			confirm, _ := reader.ReadString('\n')
			confirm = strings.TrimSuffix(confirm, "\r\n")
			//validate confirmation entry
			confirm = validateuserinput(confirm, "confirm")

			if confirm == "y" || confirm == "Y" {
				fmt.Println("Confirm Delete")
				deletedeck(ddeck)
			} else if confirm == "n" || confirm == "N" {
				menu()
			}
		} else if edchoice == "enter" || edchoice == "Enter" {
			fmt.Println("Deck Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSuffix(name, "\r\n")
			fmt.Println("Your new deck name is " + name)
			fmt.Println("Is this deck a multi color deck?(y/n)")
			multi, _ := reader.ReadString('\n')
			multi = strings.TrimSuffix(multi, "\r\n")
			//validate user input
			multi = validateuserinput(multi, "confirm")
			var color string
			if multi == "y" {
				fmt.Println("How many colors?")
				num_col, _ := reader.ReadString('\n')
				num_col = strings.TrimSuffix(num_col, "\r\n")
				fmt.Println("You deck has " + num_col + " of colors")
				fmt.Println("what is your first color?(Black|White|Blue|Red|Green)")
				cols, _ := reader.ReadString('\n')
				cols = strings.TrimSuffix(cols, "\r\n")
				//validate user input
				cols = validateuserinput(cols, "colors")
				count := 1
				snum, _ := strconv.Atoi(num_col)
				for count != snum {
					count++
					fmt.Println("Next Color(Black|White|Blue|Red|Green): ")
					ncol, _ := reader.ReadString('\n')
					ncol = strings.TrimSuffix(ncol, "\r\n")
					validateuserinput(ncol, "colors")
					cols = cols + "," + ncol
				}
				color = cols
			} else if multi == "n" {
				fmt.Println("What color is your deck?(Black|White|Blue|Red|Green)")
				color, _ = reader.ReadString('\n')
				color = strings.TrimSuffix(color, "\r\n")
				validateuserinput(color, "colors")
			}
			fmt.Println("Favorite(y/n): ")
			favorite, _ := reader.ReadString('\n')
			favorite = strings.TrimSuffix(favorite, "\r\n")
			//validate user input
			validateuserinput(favorite, "confirm")
			fmt.Println("Your new deck " + name + " is a favorite: " + favorite)
			//convert string to int equivilant
			favorite_bin := new(int)
			if favorite == "y" {
				*favorite_bin = 0
			} else {
				*favorite_bin = 1
			}
			fmt.Println("Total Number of cards: ")
			numcards, _ := reader.ReadString('\n')
			numcards = strings.TrimSuffix(numcards, "\r\n")
			icards := new(int)
			*icards, _ = strconv.Atoi(numcards)
			fmt.Print("Total number of cards: " + numcards + "\n")
			fmt.Println("Total number of lands: ")
			numlands, _ := reader.ReadString('\n')
			numlands = strings.TrimSuffix(numlands, "\r\n")
			ilands := new(int)
			*ilands, _ = strconv.Atoi((numlands))
			fmt.Print("Total number of lands: " + numlands + "\n")
			fmt.Println("Total number of instant/sorcery/enchantment: ")
			numspells, _ := reader.ReadString('\n')
			numspells = strings.TrimSuffix(numspells, "\r\n")
			ispells := new(int)
			*ispells, _ = strconv.Atoi(numspells)
			fmt.Print("Total number of instant/sorcery/enchantment: " + numspells + "\n")
			fmt.Println("Total number of creatures: ")
			numcreatures, _ := reader.ReadString('\n')
			numcreatures = strings.TrimSuffix(numcreatures, "\r\n")
			icreatures := new(int)
			*icreatures, _ = strconv.Atoi(numcreatures)
			fmt.Print("Total number of creatures: " + numcreatures + "\n")
			//enter into database
			//newdeck(name, color, favorite)
			d := Deck{
				Name:         name,
				Colors:       color,
				Date_Entered: time.Now(),
				Favorite:     int(*favorite_bin),
				Num_Cards:    int(*icards),
				Num_Lands:    int(*ilands),
				Num_Spells:   int(*ispells),
				Num_Creat:    int(*icreatures),
			}
			err := newdeck(d)
			if err != nil {
				log.Printf("Insert deck failed with error %s", err)
				return
			}
		}
	case 2:
		//reader := bufio.NewReader(os.Stdin)
		fmt.Println("Results(won/lost): ")
		results, _ := reader.ReadString('\n')
		results = strings.TrimSuffix(results, "\r\n")
		//validate results input
		results = validateuserinput(results, "results")
		fmt.Print("You " + results + " this game.\n")
		fmt.Println("Why do you think you " + results + " this game?")
		cause, _ := reader.ReadString('\n')
		cause = strings.TrimSuffix(cause, "\r\n")
		fmt.Println("Cause: " + cause)
		fmt.Println("What deck were you using?")
		deck, _ := reader.ReadString('\n')
		deck = strings.TrimSuffix(deck, "\r\n")
		//validate deck name
		deck = validatedeck(deck)
		fmt.Println("You " + results + " your game using " + deck)
		fmt.Println("Who was your opponent?")
		opp, _ := reader.ReadString('\n')
		opp = strings.TrimSuffix(opp, "\r\n")
		fmt.Println("Your opponent was: " + opp)
		fmt.Println("What Level Was the Game?(Bronze, Silver, Gold, Platinum, Diamond, and Mythic)")
		lev, _ := reader.ReadString('\n')
		lev = strings.TrimSuffix(lev, "\r\n")
		//validate level input
		lev = validateuserinput(lev, "level")
		fmt.Println("What Tier Was the Game?(1-4)")
		tier, _ := reader.ReadString('\n')
		tier = strings.TrimSuffix(tier, "\r\n")
		//validate tier input
		tier = validateuserinput(tier, "tier")
		fmt.Println("Your Game was Level " + lev + " and Tier " + tier)
		cmblvl := lev + "-" + tier
		//convert string to int
		results_bin := new(int)
		if results == "won" {
			*results_bin = 0
		} else {
			*results_bin = 1
		}

		//enter into database
		g := Game{
			Results:  int(*results_bin),
			Cause:    cause,
			Deck:     deck,
			Opponent: opp,
			Level:    cmblvl,
		}
		err := newgame(g)
		if err != nil {
			log.Printf("Insert game failed with error %s", err)
			return
		}
	case 3:
		//reader := bufio.NewReader(os.Stdin)
		fmt.Println("Would you like to narrow your search?(y/n)")
		deckchoice, _ := reader.ReadString('\n')
		deckchoice = strings.TrimSuffix(deckchoice, "\r\n")
		validateuserinput(deckchoice, "confirm")
		if deckchoice == "y" {
			fmt.Println("Deck Name: ")
			deckname, _ := reader.ReadString('\n')
			viewrecords(deckname)
		} else {
			viewrecords("n")
		}
	case 4:
		//reader := bufio.NewReader(os.Stdin)
		fmt.Println("Deck:")
		deckgames, _ := reader.ReadString('\n')
		deckgames = strings.TrimSuffix(deckgames, "\r\n")
		gamecount(deckgames)
	case 5:
		//reader := bufio.NewReader(os.Stdin)
		fmt.Println("Would you like to see a specific deck details?(y/n)")
		deckchoice, _ := reader.ReadString('\n')
		deckchoice = strings.TrimSuffix(deckchoice, "\r\n")
		if deckchoice == "y" {
			fmt.Println("Deck Name: ")
			deckname, _ := reader.ReadString('\n')
			viewdecks(deckname, 0)
		} else {
			viewdecks("n", 0)
		}
	case 6:
		topten()
	case 7:
		//reader := bufio.NewReader(os.Stdin)
		fmt.Println("Deck: ")
		deckchoice, _ := reader.ReadString('\n')
		editdeck(deckchoice)
	case 8:
		fmt.Println("Deck: ")
		deck, _ := reader.ReadString('\n')
		deck = strings.TrimSuffix(deck, "\r\n")
		deck = validatedeck(deck)
		fmt.Println("Wins or Loses?")
		wlpct, _ := reader.ReadString('\n')
		wlpct = strings.TrimSuffix(wlpct, "\r\n")
		wlpct = validateuserinput(wlpct, "percent")
		pctvals(wlpct, deck)
	case 9:
		anal_menu()
	case 10:
		os.Exit(0)
	default:
		//Reset Menu for invalid in put
		fmt.Println("Invalid Selection.")
		fmt.Println("")
		menu()
	}
}
func opendb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mtga?parseTime=true")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	return db
}
func newdeck(d Deck) error {

	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	query := "INSERT INTO mtga.decks(name, colors, favorite, numcards, numlands, numspells, numcreatures) VALUES (?, ?, ?, ?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		panic(err.Error())
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, d.Name, d.Colors, d.Favorite, d.Num_Cards, d.Num_Lands, d.Num_Spells, d.Num_Creat)
	if err != nil {
		log.Printf("Error %s when inserting row into deck table", err)
		panic(err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		panic(err.Error())
	}
	log.Printf("%d deck created ", rows)
	fmt.Println("")
	menu()
	return nil
}
func newgame(g Game) error {

	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	query := "INSERT INTO mtga.games(results, cause, deck, opponent, level) VALUES (?, ?, ?,?,?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		panic(err.Error())
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, g.Results, g.Cause, g.Deck, g.Opponent, g.Level)
	if err != nil {
		log.Printf("Error %s when inserting row into deck table", err)
		panic(err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		panic(err.Error())
	}
	log.Printf("%d row added ", rows)
	//determin max and current streak
	streaks(g.Deck)
	fmt.Println("")
	menu()
	return nil
}
func viewrecords(DeckName string) error {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	if DeckName != "n" {
		var (
			deckname string
			wins     int
			loses    int
		)
		DeckName = strings.TrimSuffix(DeckName, "\r\n")
		DeckName = validatedeck(DeckName)

		// Execute the query
		results := db.QueryRow("SELECT deck, wins, loses FROM mtga.record WHERE deck=?", DeckName)
		err := results.Scan(&deckname, &wins, &loses)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				menu()
			} else {
				panic(err.Error())
			}
		}
		deckname = fmt.Sprintf("%-25s", deckname)
		fwins := fmt.Sprintf("%-10s", "Wins: "+strconv.Itoa(wins))
		floses := fmt.Sprintf("%-5s", "Loses: "+strconv.Itoa(loses))
		finalrecord := fmt.Sprint(deckname + fwins + floses)
		log.Println(finalrecord)
		fmt.Println("")
	} else {
		results, err := db.Query("SELECT deck, wins, loses FROM mtga.record ORDER BY wins desc, loses desc")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		for results.Next() {
			var records Records
			// for each row, scan the result into our deck composite object
			err = results.Scan(&records.Deck, &records.Wins, &records.Loses)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)
			records.Deck = fmt.Sprintf("%-25s", records.Deck)
			fwins := fmt.Sprintf("%-10s", "Wins: "+strconv.Itoa(records.Wins))
			floses := fmt.Sprintf("%-5s", "Loses: "+strconv.Itoa(records.Loses))
			finalrecord := fmt.Sprint(records.Deck + fwins + floses)
			log.Println(finalrecord)
		}
	}
	fmt.Println("")
	menu()
	return nil
}
func gamecount(d string) {
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	var (
		deckname string
		count    int
	)
	//validate deck name
	d = validatedeck(d)

	results := db.QueryRow("SELECT deck, results AS Count FROM mtga.game_count WHERE deck=?", d)
	err := results.Scan(&deckname, &count)
	fmt.Println("testing: " + deckname)
	if err != nil {
		panic(err.Error())
	}
	finalcount := fmt.Sprint(deckname + " Game Count: " + strconv.Itoa(count))
	log.Println(finalcount)
	fmt.Println("")
	menu()
}
func viewdecks(DeckName string, edit int) (ret string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	if DeckName != "n" {
		var d Deck
		DeckName = strings.TrimSuffix(DeckName, "\r\n")
		//validate deck name
		DeckName = validatedeck(DeckName)
		results := db.QueryRow("SELECT name, colors, date_entered, favorite, max_streak, cur_streak, numcards, numlands, numspells, numcreatures, disable FROM mtga.decks WHERE name=?", DeckName)
		err := results.Scan(&d.Name, &d.Colors, &d.Date_Entered, &d.Favorite, &d.Max_Streak, &d.Cur_Streak,
			&d.Num_Cards, &d.Num_Lands, &d.Num_Spells, &d.Num_Creat, &d.Disable)
		if err != nil {
			panic(err.Error())
			//menu()
		}
		d.Name = fmt.Sprintf("%-25s", d.Name)
		d.Colors = fmt.Sprintf("%-15s", d.Colors)
		fdate := fmt.Sprintf("%-25s", d.Date_Entered.Format("2006-01-02 15:04:05"))
		ffav := fmt.Sprintf("%-5s", strconv.Itoa(d.Favorite))
		fmax := fmt.Sprintf("%-5s", strconv.Itoa(d.Max_Streak))
		fcur := fmt.Sprintf("%-5s", strconv.Itoa(d.Cur_Streak))
		fcard := fmt.Sprintf("%-5s", strconv.Itoa(d.Num_Cards))
		fland := fmt.Sprintf("%-5s", strconv.Itoa(d.Num_Lands))
		fspell := fmt.Sprintf("%-5s", strconv.Itoa(d.Num_Spells))
		fcreat := fmt.Sprintf("%-5s", strconv.Itoa(d.Num_Creat))
		fdis := d.Disable

		var sdis string
		if fdis == 0 {
			sdis = "Yes"
			//fmt.Println("yes")
		} else {
			sdis = "No"
			//fmt.Println("No")
		}
		finalrecord := fmt.Sprint("Name: " + d.Name + "Color: " + d.Colors + "Date Entered: " + fdate + "Favorite: " +
			ffav + "\n" + "Max Streak: " + fmax + "Current Streak: " + fcur + "\n" + "Number of Cards: " + fcard + "Number of Lands: " +
			fland + "Number of Spells: " + fspell + "Number of Creatures: " + fcreat + "\n" + "Disabled: " + sdis + " \n")
		log.SetFlags(0)
		log.Println(finalrecord)
		ret = d.Name
	} else {
		results, err := db.Query("SELECT name, colors, date_entered, favorite, max_streak FROM mtga.decks ORDER BY favorite")

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		var count int
		for results.Next() {
			var deck Deck
			count++
			// for each row, scan the result into our deck composite object
			err = results.Scan(&deck.Name, &deck.Colors, &deck.Date_Entered, &deck.Favorite, &deck.Max_Streak)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.SetFlags(0)
			mstreak := deck.Max_Streak
			var fav string
			if deck.Favorite == 0 {
				fav = fmt.Sprintf("%-4s", "Yes")
			} else {
				fav = fmt.Sprintf("%-4s", "No")
			}
			//format strings to be more readable
			fcount := fmt.Sprintf("%2s: ", strconv.Itoa(count))
			deck.Name = fmt.Sprintf("%-25s", deck.Name)
			deck.Colors = fmt.Sprintf("%-15s", deck.Colors)
			fdate := fmt.Sprintf("%-20s", deck.Date_Entered.Format("2006-01-02 15:04:05"))
			fmstreak := fmt.Sprintf("%-4s", strconv.Itoa(mstreak))

			finalrecord := fmt.Sprint(fcount + deck.Name + " Colors: " + deck.Colors + " Date Entered: " + fdate +
				" Favorite: " + fav + " Max Streak: " + fmstreak)
			log.Println(finalrecord)
		}
	}
	if edit == 0 {
		fmt.Println("")
		menu()
	}
	return
}
func topten() {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	results, err := db.Query("SELECT deck, wins, loses FROM mtga.topten")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var (
			name  string
			wins  int
			loses int
		)

		// for each row, scan the result into our deck composite object
		err = results.Scan(&name, &wins, &loses)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		log.SetFlags(0)

		//format strings to be more readable
		name = fmt.Sprintf("%-25s", name)
		fwins := fmt.Sprintf("%-5s", strconv.Itoa(wins))
		floses := fmt.Sprintf("%-5s", strconv.Itoa(loses))

		finalrecord := fmt.Sprint(name + " Wins: " + fwins + " Loses: " + floses)
		log.Println(finalrecord)
	}
	fmt.Println("")
	menu()
}
func editdeck(d string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	//show current deck attributes
	d = viewdecks(d, 1)
	//create new deck structure variable
	var deck Deck
	d = strings.TrimSuffix(d, "\r\n")
	results := db.QueryRow("SELECT name, colors, date_entered, favorite, max_streak, cur_streak, numcards, numlands, numspells, numcreatures, disable FROM mtga.decks WHERE name=?", d)
	err := results.Scan(&deck.Name, &deck.Colors, &deck.Date_Entered, &deck.Favorite, &deck.Max_Streak, &deck.Cur_Streak,
		&deck.Num_Cards, &deck.Num_Lands, &deck.Num_Spells, &deck.Num_Creat, &deck.Disable)
	if err != nil {
		panic(err.Error())
	}
	//determine which section is to be edited
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter section you would like to edit or finish to return to main menu")
	editchoice, _ := reader.ReadString('\n')
	fmt.Println("Edit Section: " + editchoice)
	editchoice = strings.TrimSuffix(editchoice, "\r\n")

	switch editchoice {
	case "name":
		fmt.Println("Original Name: " + deck.Name)
		fmt.Print("New Name: ")
		oname := deck.Name
		reader := bufio.NewReader(os.Stdin)
		deck.Name, _ = reader.ReadString('\n')
		fmt.Println("Update Name: " + deck.Name)
		deck.Name = strings.TrimSuffix(deck.Name, "\r\n")
		updatedeck(deck, oname)
	case "color":
		fmt.Println("Original Color: " + deck.Colors)
		fmt.Println("New Colors: ")
		reader := bufio.NewReader(os.Stdin)
		deck.Colors, _ = reader.ReadString('\n')
		fmt.Println("Update Colors: " + deck.Colors)
		deck.Colors = strings.TrimSuffix(deck.Colors, "\r\n")
		updatedeck(deck, deck.Name)
	case "date entered":
		fmt.Println("Cannot change this field.")
		editdeck(d)
	case "favorite":
		fmt.Println("Is This Deck a Favorite(y/n): ")
		reader := bufio.NewReader(os.Stdin)
		sfav, _ := reader.ReadString('\n')
		sfav = strings.TrimSuffix(sfav, "\r\n")
		if sfav == "y" {
			deck.Favorite = 0
		} else {
			deck.Favorite = 1
		}
		fmt.Println(deck.Favorite)
		fmt.Println("Update Favorite: " + sfav)
		updatedeck(deck, deck.Name)
	case "max streak":
		fmt.Println("Cannot change this field")
		editdeck(d)
	case "current streak":
		fmt.Println("Cannot change this field")
		editdeck(d)
	case "number of cards":
		fmt.Println("Original Total Number of Cards: " + strconv.Itoa(deck.Num_Cards))
		fmt.Println("New Total Number of Cards: ")
		reader := bufio.NewReader(os.Stdin)
		scards, _ := reader.ReadString('\n')
		scards = strings.TrimSuffix(scards, "\r\n")
		deck.Num_Cards, _ = strconv.Atoi(scards)
		fmt.Println("Update Total Number of Cards: " + scards)
		updatedeck(deck, deck.Name)
	case "number of lands":
		fmt.Println("Original Total Number of Lands: " + strconv.Itoa(deck.Num_Lands))
		fmt.Println("New Total Number of Lands: ")
		reader := bufio.NewReader(os.Stdin)
		slands, _ := reader.ReadString('\n')
		slands = strings.TrimSuffix(slands, "\r\n")
		deck.Num_Lands, _ = strconv.Atoi(slands)
		fmt.Println("Update Total Number of Lands: " + slands)
		updatedeck(deck, deck.Name)
	case "number of spells":
		fmt.Println("Original Total Number of Instant/Sorcery/Enchantment: " + strconv.Itoa(deck.Num_Spells))
		fmt.Println("New Total Number of Instant/Sorcery/Enchantment: ")
		reader := bufio.NewReader(os.Stdin)
		sspells, _ := reader.ReadString('\n')
		sspells = strings.TrimSuffix(sspells, "\r\n")
		deck.Num_Spells, _ = strconv.Atoi(sspells)
		fmt.Println("Update Total Number of Instant/Sorcery/Enchantment: " + sspells)
		updatedeck(deck, deck.Name)
	case "number of creatures":
		fmt.Println("Original Total Number of creatures: " + strconv.Itoa(deck.Num_Creat))
		fmt.Println("New Total Number of creatures: ")
		reader := bufio.NewReader(os.Stdin)
		screat, _ := reader.ReadString('\n')
		screat = strings.TrimSuffix(screat, "\r\n")
		deck.Num_Creat, _ = strconv.Atoi(screat)
		fmt.Println("Update Total Number of creatures: " + screat)
		updatedeck(deck, deck.Name)
	case "disabled":
		fmt.Println("Do You Want this Deck Disabled(y/n): ")
		reader := bufio.NewReader(os.Stdin)
		sdis, _ := reader.ReadString('\n')
		sdis = strings.TrimSuffix(sdis, "\r\n")
		if sdis == "y" {
			deck.Disable = 0
		} else {
			deck.Disable = 1
		}
		fmt.Println(deck.Disable)
		fmt.Println("Update Disable: " + sdis)
		updatedeck(deck, deck.Name)
	case "finish":
		menu()
	default:
		//Reset Menu for invalid in put
		fmt.Println("Invalid Selection.")
		editdeck(d)
	}
}
func updatedeck(d Deck, oname string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	defer db.Close()

	// perform a db.Query insert
	result, err := db.Exec("UPDATE mtga.decks SET name=?, colors=?, favorite=?, numcards=?, numlands=?, numspells=?, numcreatures=?,disable=? WHERE name=?",
		d.Name, d.Colors, d.Favorite, d.Num_Cards, d.Num_Lands, d.Num_Spells, d.Num_Creat, d.Disable, oname)

	rows, _ := result.RowsAffected()

	fmt.Println(rows)
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		panic(err.Error())
	}
	log.Println("deck updated ")
	editdeck(d.Name)
}
func deletedeck(deck string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	defer db.Close()

	// archive deck record
	query := "INSERT INTO mtga.decks_deleted(name, colors, date_entered, favorite, max_streak, cur_streak, numcards, numlands, numspells, numcreatures, disable) SELECT name, colors, date_entered, favorite, max_streak, cur_streak, numcards, numlands, numspells, numcreatures, disable FROM mtga.decks WHERE name=?"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		panic(err.Error())
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, deck)
	if err != nil {
		log.Printf("Error %s when inserting row into deck table", err)
		panic(err.Error())
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		panic(err.Error())
	}
	log.Printf("%d deck archived ", rows)

	//delete record from deck table
	query = "DELETE FROM mtga.decks WHERE name=?"
	ctx, cancelfunc = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err = db.PrepareContext(ctx, query)
	if err != nil {
		fmt.Printf("Error %s when preparing SQL statement", err)
		panic(err.Error())
	}
	defer stmt.Close()
	res, err = stmt.ExecContext(ctx, deck)
	if err != nil {
		log.Printf("Error %s when deleting row from deck table", err)
		panic(err.Error())
	}
	rows, err = res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		panic(err.Error())
	}
	fmt.Printf("%d deck deleted\n", rows)
	fmt.Println("")
	menu()
}
func validatedeck(deck string) (deckname string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	reader := bufio.NewReader(os.Stdin)
	// Verify Database Name
	//var deckname string
	results := db.QueryRow("SELECT name FROM mtga.decks WHERE name=?", deck)
	err := results.Scan(&deckname)
	for err != nil {
		fmt.Println("Deck Does Not Exist")
		fmt.Println("Deck Name: ")
		deck, _ := reader.ReadString('\n')
		deck = strings.TrimSuffix(deck, "\r\n")
		results := db.QueryRow("SELECT name FROM mtga.decks WHERE name=?", deck)
		err = results.Scan(&deckname)
	}
	return
}
func validateuserinput(s string, u string) (ret string) {
	reader := bufio.NewReader(os.Stdin)
	switch u {
	case "level":
		re, _ := regexp.Compile(`Bronze|Silver|Gold|Platinum|Diamond|Mythic`)
		for !re.MatchString(s) {
			fmt.Println("Invalid Entry. Try Again")
			fmt.Println("What Level Was the Game?(Bronze, Silver, Gold, Platinum, Diamond, and Mythic)")
			s, _ = reader.ReadString('\n')
			s = strings.TrimSuffix(s, "\r\n")
		}
	case "tier":
		re, _ := regexp.Compile(`[1-4]`)
		for !re.MatchString(s) {
			fmt.Println("Invalid Entry. Try Again")
			fmt.Println("What Tier Was the Game?(1-4)")
			s, _ = reader.ReadString('\n')
			s = strings.TrimSuffix(s, "\r\n")
		}
	case "results":
		re, _ := regexp.Compile(`won|lost`)
		for !re.MatchString(s) {
			fmt.Println("Invalid Entry. Try Again")
			fmt.Println("Results(won/lost): ")
			s, _ = reader.ReadString('\n')
			s = strings.TrimSuffix(s, "\r\n")
		}
	case "deck":
		s = validatedeck(s)
	case "choice":
		re, _ := regexp.Compile(`enter|Enter|delete|Delete`)
		for !re.MatchString(s) || len(s) > 6 {
			fmt.Println("Invalid Entry. Try Again")
			fmt.Println("Enter or Delete Deck: ")
			s, _ = reader.ReadString('\n')
			s = strings.TrimSuffix(s, "\r\n")
		}
	case "confirm":
		re, _ := regexp.Compile(`[yn]`)
		for !re.MatchString(s) || len(s) > 1 {
			fmt.Println("Invalid Entry. Try Again")
			fmt.Println("Confirm(y/n): ")
			s, _ = reader.ReadString('\n')
			s = strings.TrimSuffix(s, "\r\n")
		}
	case "colors":
		re, _ := regexp.Compile(`Black|White|Blue|Red|Green|Colorless`)
		for !re.MatchString(s) {
			fmt.Println("Invalid Entry. Try Again")
			fmt.Println("What Color?(Black|White|Blue|Red|Green)")
			s, _ = reader.ReadString('\n')
			s = strings.TrimSuffix(s, "\r\n")
		}
	case "percent":
		re, _ := regexp.Compile(`wins|loses`)
		for !re.MatchString(s) {
			fmt.Println("Invalid Entry")
			fmt.Println("Wins or Loses Percentages?")
			s, _ = reader.ReadString('\n')
			s = strings.TrimSuffix(s, "\r\n")
		}
	}
	ret = s
	return
}
func pctvals(s string, d string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	var (
		deckname string
		pct      float32
		count    int
		games    int
	)

	if s == "wins" {
		results := db.QueryRow("SELECT deck,win_pct,win_count,games FROM mtga.win_percentage WHERE deck =?", d)
		err := results.Scan(&deckname, &pct, &count, &games)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				menu()
			} else {
				panic(err.Error())
			}
		}
		fpct := fmt.Sprintf("%f", pct)
		fpct = fpct[2:4]
		tpct := fpct[0:1]
		println("Print test two: " + tpct)
		if pct == 1 {
			fpct = "100"
			println("test one: " + fpct)
		}
		println("testing: " + fpct)
		finalprint := fmt.Sprint(deckname + "   Win Percentage: " + fpct + "%    Number of Wins: " + strconv.Itoa(count) +
			"    Number of Games: " + strconv.Itoa(games))
		log.Println(finalprint)
		fmt.Println("")
		menu()
	} else if s == "loses" {
		results := db.QueryRow("SELECT deck,lose_pct,lose_count,games FROM mtga.lose_percentage WHERE deck =?", d)
		err := results.Scan(&deckname, &pct, &count, &games)
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				fmt.Println("No Games Recored for this Deck")
				fmt.Println("")
				menu()
			} else {
				panic(err.Error())
			}
		}
		fpct := fmt.Sprintf("%f", pct)
		fpct = fpct[2:4]
		finalprint := fmt.Sprint(deckname + "   Lose Percentage: " + fpct + "%    Number of Loses: " + strconv.Itoa(count) +
			"    Number of Games: " + strconv.Itoa(games))
		log.Println(finalprint)
		fmt.Println("")
		menu()
	}
	menu()
}
func streaks(d string) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	var (
		max    int
		cur    int
		streak int
	)
	println("Deck Name: " + d)
	results, err := db.Query("SELECT deck, results FROM mtga.games WHERE deck=?", d)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		var (
			name   string
			result int
		)

		// for each row, scan the result into our deck composite object
		err = results.Scan(&name, &result)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		//track and store streak values
		if result == 0 {
			if streak == 0 {
				cur++
				if cur > max {
					max = cur
				}
			} else if streak == 1 {
				streak = 0
				cur++
			}
		} else if result == 1 {
			streak = 1
			cur = 0
		}
	}

	// perform a db.Query insert
	upresult, err := db.Exec("UPDATE mtga.decks SET max_streak=?, cur_streak=? where name=?", max, cur, d)

	rows, _ := upresult.RowsAffected()

	//fmt.Println(rows)
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		panic(err.Error())
	}
	log.Println("deck updated ", rows)
}
