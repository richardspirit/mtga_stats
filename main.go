package main

import (
	"fmt"
	"log"
	"time"
	"bufio"
	"os"
	"strconv"
	"context"
	"database/sql"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	menu()
}

/*
 * very simple struct
 */
type Deck struct {
    Name           string       `json:"name"`
	Colors         string       `json:"colors"`
    Date_Entered   time.Time    `json:"date_entered"`
	Favorite       int          `json:"favorite"`
	Max_Streak     int          `json:"max_streak"`
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
	fmt.Println("")
	fmt.Println("1. Enter New Deck")
	fmt.Println("2. Add New Game")
	fmt.Println("3. View Deck Records")
	fmt.Println("4. View Game Count")
	fmt.Println("Q. Quit")
	fmt.Println("")
	fmt.Println("What do you want to do?")
	
	in := bufio.NewScanner(os.Stdin)
    in.Scan()
    choice,_ := strconv.Atoi(in.Text())
	
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
		if ( favorite == "y" ) {
			*favorite_bin = 0;
		} else {
			*favorite_bin = 1;
		}
		//enter into database
		//newdeck(name, color, favorite)
		d := Deck{  
			Name:  name,
			Colors: color,
			Date_Entered: time.Now(),
			Favorite: int(*favorite_bin),
			Max_Streak: 0,
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
		if ( results == "won" ) {
			*results_bin = 0;
		} else {
			*results_bin = 1;
		}
		//enter into database
		g := Game{  
			Results:  int(*results_bin),
			Cause: cause,
			Deck: deck,
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
	default:
		os.Exit(0)
    }
}

func opendb() (*sql.DB) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mgta")
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
	query := "INSERT INTO mgta.decks(name, colors, favorite) VALUES (?, ?, ?)"
    ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
	    panic(err.Error())
    }
    defer stmt.Close()
	res, err := stmt.ExecContext(ctx, d.Name, d.Colors, d.Favorite)
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
		var (deckname string
			 wins int
			 loses int)
		DeckName = strings.TrimSuffix(DeckName, "\r\n")
	    // Execute the query
		results := db.QueryRow("SELECT deck, wins, loses FROM mgta.record WHERE deck=?", DeckName)
		err := results.Scan(&deckname, &wins, &loses)
			if err != nil {
				panic(err.Error())
			}		
		finalrecord := fmt.Sprint(deckname + " Wins: " + strconv.Itoa(wins) + " Loses: " + strconv.Itoa(loses))
		log.Printf(finalrecord)
	}else {
		results, err := db.Query("SELECT deck, wins, loses FROM mgta.record")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	
		for results.Next() {
			var records Records
			// for each row, scan the result into our deck composite object
			err = results.Scan(&records.Deck, &records.Wins, &records.Loses)
		//  if err != nil {
			//    panic(err.Error()) // proper error handling instead of panic in your app
			//}
			// and then print out the tag's Name attribute
			log.SetFlags(0)
			finalrecord := fmt.Sprint(records.Deck + " Wins: " + strconv.Itoa(records.Wins) + " Loses: " + strconv.Itoa(records.Loses))
			log.Printf(finalrecord)
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
	
	var (deckname string
		 count int)
	//DeckName = strings.TrimSuffix(DeckName, "\r\n")	
	results := db.QueryRow("SELECT deck, results AS Count FROM mgta.game_count WHERE deck=?", d)
	err := results.Scan(&deckname, &count)
		if err != nil {
			panic(err.Error())
		}		
	finalcount := fmt.Sprint(deckname + " Game Count: " + strconv.Itoa(count))
	log.Printf(finalcount)
	
	menu()
	//return nil
}