package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/gorp.v2"
	"os"
	"os/signal"

	"syscall"
	"time"
)

var (
	dbMap *gorp.DbMap
)

func main() {
	host := flag.String("h", "127.0.0.1:3306", "mysql host:port")
	user := flag.String("u", "root", "mysql user")
	password := flag.String("p", "", "mysql user password, ask if empty")
	repeat := flag.Int64("r", 0, "repeat every N seconds, no repeat if 0")
	secondsOnly := flag.Bool("s", false, "show only seconds behind master")
	flag.Parse()
	if *password == "" {
		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		fmt.Println()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
		*password = string(bytePassword)
	}
	//fmt.Printf("connect to %s@%s\n", *user, *host)
	connString := fmt.Sprintf("%s:%s@tcp(%s)/?tls=skip-verify&parseTime=true", *user, *password, *host)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Printf("Error connecting db: %s\n", err.Error())
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error connecting db: %s\n", err.Error())
		return
	}
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		//fmt.Println("Ctrl+C")
		db.Close()
		os.Exit(1)
	}()
	dbMap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	printSlaveStatus(*secondsOnly)
	if *repeat > 0 {
		t := time.Tick(time.Duration(*repeat) * time.Second)
		for range t {
			printSlaveStatus(*secondsOnly)
		}
	}
	db.Close()
}

func printSlaveStatus(secondsOnly bool) {
	status := SlaveStatus{}
	err := dbMap.SelectOne(&status, "show slave status")

	if err != nil {
		fmt.Printf("cannot get slave status: %s\n", err.Error())
		return
	}
	if secondsOnly {
		if status.Seconds_Behind_Master.Valid {
			fmt.Println(status.Seconds_Behind_Master.Int64)
		} else {
			fmt.Println("Null")
		}
		return
	}
	if status.Slave_SQL_Running.String == "No" {
		fmt.Println("Slave not running")
	} else if status.Seconds_Behind_Master.Valid {
		fmt.Printf("Seconds behind master: %d\n", status.Seconds_Behind_Master.Int64)
	}
	if status.Slave_SQL_Running_State.String != "" {
		fmt.Printf("State: %s\n", status.Slave_SQL_Running_State.String)
	}
	if status.Last_Error.String != "" {
		fmt.Printf("Last error: %s\n", status.Last_Error.String)
	}
}
