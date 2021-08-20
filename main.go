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

/*
 * Deck... - a very simple struct
 */
type Deck struct {
    Name           string       `json:"name"`
	Colors         string       `json:"colors"`
    Date_Entered   time.Time    `json:"date_entered"`
	Favorite       int          `json:"favorite"`
	Max_Streak     int          `json:"max_streak"`
}

func newdeck(d Deck) error {

    // Open up our database connection.
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mgta")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()
	
	// perform a db.Query insert
    //insert, err := db.Query("INSERT INTO mgta.decks (Name, Colors, Date_Entered, Favorite, Max_Streak) VALUES('White Flame', 'White,Red', 1, NULL);")
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
	return nil
}

func main() {
    fmt.Println("MGTA Stats")
	fmt.Println("")
	fmt.Println("1. Enter New Deck")
	fmt.Println("2. Add New Game")
	fmt.Println("3. View Deck Records")
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
        fmt.Println("Your choice is two")
    case 3:
        fmt.Println("Your choice is three")
	default:
		fmt.Println("No Choices.")
    }
	
    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mgta")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()
	
	    // Execute the query
    results, err := db.Query("SELECT name, colors, date_entered, favorite, max_streak FROM mgta.decks")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	
    for results.Next() {
        var deck Deck
        // for each row, scan the result into our deck composite object
        err = results.Scan(&deck.Name, &deck.Colors, &deck.Date_Entered, &deck.Favorite, &deck.Max_Streak)
      //  if err != nil {
        //    panic(err.Error()) // proper error handling instead of panic in your app
        //}
                // and then print out the tag's Name attribute
		log.SetFlags(0)
        //log.Printf(deck.Name)
    }	

}