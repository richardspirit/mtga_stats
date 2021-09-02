package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	menu()
}

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
	Results int    `json:"results"`
	Cause   string `json:"cause"`
	Deck    string `json:"Deck"`
}

type Records struct {
	Deck  string `json:"deck"`
	Wins  int    `json:"wins"`
	Loses int    `json:"loses"`
}

func menu() {
	//main menu
	fmt.Println("MGTA Stats")
	m := make(map[string]string)

	// Set key/value pairs using typical `name[key] = val`
	m["k1"] = fmt.Sprintf("%-20s", "Enter New Deck")
	m["k2"] = fmt.Sprintf("%-20s", "Add New Game")
	m["k3"] = fmt.Sprintf("%-20s", "View Deck Records")
	m["k4"] = fmt.Sprintf("%-20s", "View Game Count")
	m["k5"] = fmt.Sprintf("%-20s", "View Decks")
	m["k6"] = fmt.Sprintf("%-20s", "Top Ten Decks")
	m["k7"] = fmt.Sprintf("%-20s", "Edit Deck")
	m["k10"] = fmt.Sprintf("%20s", "10: Quit")

	// print menu options
	fmt.Println("1:", m["k1"]+"2:", m["k2"])
	fmt.Println("3:", m["k3"]+"4:", m["k4"])
	fmt.Println("5:", m["k5"]+"6:", m["k6"])
	fmt.Println("7:", m["k7"])
	fmt.Println(m["k10"])

	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	choice, _ := strconv.Atoi(in.Text())

	switch choice {
	case 1:
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Deck Name: ")
		name, _ := reader.ReadString('\n')
		fmt.Print("Your new deck name is " + name)
		fmt.Println("Deck Color/s: ")
		color, _ := reader.ReadString('\n')
		fmt.Print("Your new deck color/s is/are " + color)
		fmt.Println("Favorite(y/n): ")
		favorite, _ := reader.ReadString('\n')
		name = strings.TrimSuffix(name, "\r\n")
		color = strings.TrimSuffix(color, "\r\n")
		fmt.Print("Your new deck " + name + " is a favorite: " + favorite)
		favorite = strings.TrimSuffix(favorite, "\r\n")
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
	case 2:
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Results(won/lost): ")
		results, _ := reader.ReadString('\n')
		results = strings.TrimSuffix(results, "\r\n")
		fmt.Print("You " + results + " this game.\n")
		fmt.Println("Why do you think you " + results + " this game?")
		cause, _ := reader.ReadString('\n')
		fmt.Print("Cause: " + cause)
		fmt.Println("What deck were you using?")
		deck, _ := reader.ReadString('\n')
		fmt.Print("You " + results + " your game using " + deck)
		cause = strings.TrimSuffix(cause, "\r\n")
		deck = strings.TrimSuffix(deck, "\r\n")
		results_bin := new(int)
		if results == "won" {
			*results_bin = 0
		} else {
			*results_bin = 1
		}
		//enter into database
		g := Game{
			Results: int(*results_bin),
			Cause:   cause,
			Deck:    deck,
		}
		err := newgame(g)
		if err != nil {
			log.Printf("Insert game failed with error %s", err)
			return
		}
	case 3:
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Would you like to narrow your search?(y/n)")
		deckchoice, _ := reader.ReadString('\n')
		deckchoice = strings.TrimSuffix(deckchoice, "\r\n")
		if deckchoice == "y" {
			fmt.Println("Deck Name: ")
			deckname, _ := reader.ReadString('\n')
			viewrecords(deckname)
		} else {
			viewrecords("n")
		}
	case 4:
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Deck:")
		deckgames, _ := reader.ReadString('\n')
		deckgames = strings.TrimSuffix(deckgames, "\r\n")
		gamecount(deckgames)
	case 5:
		reader := bufio.NewReader(os.Stdin)
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
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Deck: ")
		deckchoice, _ := reader.ReadString('\n')
		editdeck(deckchoice)
	case 10:
		os.Exit(0)
	default:
		//Reset Menu for invalid in put
		fmt.Println("Invalid Selection.")
		menu()
	}
}

func opendb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mgta?parseTime=true")
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
	query := "INSERT INTO mgta.decks(name, colors, favorite, numcards, numlands, numspells, numcreatures) VALUES (?, ?, ?, ?, ?, ?, ?)"
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
	query := "INSERT INTO mgta.games(results, cause, deck) VALUES (?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		panic(err.Error())
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, g.Results, g.Cause, g.Deck)
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
		// Execute the query
		results := db.QueryRow("SELECT deck, wins, loses FROM mgta.record WHERE deck=?", DeckName)
		err := results.Scan(&deckname, &wins, &loses)
		if err != nil {
			panic(err.Error())
		}
		deckname = fmt.Sprintf("%-25s", deckname)
		fwins := fmt.Sprintf("%-10s", "Wins: "+strconv.Itoa(wins))
		floses := fmt.Sprintf("%-5s", "Loses: "+strconv.Itoa(loses))
		finalrecord := fmt.Sprint(deckname + fwins + floses)
		log.Println(finalrecord)
	} else {
		results, err := db.Query("SELECT deck, wins, loses FROM mgta.record ORDER BY wins desc, loses desc")
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
	//DeckName = strings.TrimSuffix(DeckName, "\r\n")
	results := db.QueryRow("SELECT deck, results AS Count FROM mgta.game_count WHERE deck=?", d)
	err := results.Scan(&deckname, &count)
	if err != nil {
		panic(err.Error())
	}
	finalcount := fmt.Sprint(deckname + " Game Count: " + strconv.Itoa(count))
	log.Println(finalcount)

	menu()
}

func viewdecks(DeckName string, edit int) {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	if DeckName != "n" {
		var d Deck
		DeckName = strings.TrimSuffix(DeckName, "\r\n")
		results := db.QueryRow("SELECT name, colors, date_entered, favorite, max_streak, cur_streak, numcards, numlands, numspells, numcreatures, disable FROM mgta.decks WHERE name=?", DeckName)
		err := results.Scan(&d.Name, &d.Colors, &d.Date_Entered, &d.Favorite, &d.Max_Streak, &d.Cur_Streak,
			&d.Num_Cards, &d.Num_Lands, &d.Num_Spells, &d.Num_Creat, &d.Disable)
		if err != nil {
			panic(err.Error())
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
		fdis := fmt.Sprintf("%-5s", strconv.Itoa(d.Disable))
		if fdis == "0" {
			fdis = "Yes"
		} else {
			fdis = "No"
		}
		finalrecord := fmt.Sprint("Name: " + d.Name + "Color/s: " + d.Colors + "Date Entered: " + fdate + "Favorite: " +
			ffav + "\n" + "Max Streak: " + fmax + "Current Streak: " + fcur + "\n" + "Number of Cards: " + fcard + "Number of Lands: " +
			fland + "Number of Spells: " + fspell + "Number of Creatures: " + fcreat + "\n" + "Disabled: " + fdis + " \n")
		log.SetFlags(0)
		log.Println(finalrecord)
	} else {
		results, err := db.Query("SELECT name, colors, date_entered, favorite, max_streak FROM mgta.decks ORDER BY favorite")

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
		menu()
	}
}
func topten() {
	// Open up our database connection.
	db := opendb()
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	results, err := db.Query("SELECT deck, wins, loses FROM mgta.topten")

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
	menu()
}
func editdeck(d string) {
	viewdecks(d, 1)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter section you would like to edit")
	editchoice, _ := reader.ReadString('\n')
	fmt.Println("Edit Section: " + editchoice)
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	echoice := in.Text()
	switch echoice {
	case "name":
		fmt.Println("Name")
	}
}
