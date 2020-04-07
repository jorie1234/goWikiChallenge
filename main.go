package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
	"github.com/joho/godotenv"

	"github.com/jorie1234/goWikiChallenge/challenge"
	"github.com/jorie1234/goWikiChallenge/confluence"

)

func main() {

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}
	wikiHost := os.Getenv("WIKI_HOST")
	USER := os.Getenv("USER")
	IDList := os.Getenv("ID_LIST")
	PWD := os.Getenv("PWD")
	if len(PWD)==0 {
		log.Printf("Enter PW for user %v: ", USER)
		pwdbyte, _ := terminal.ReadPassword(int(syscall.Stdin))
		PWD = string(pwdbyte)
	}
	c := confluence.NewConfluence(wikiHost, USER, PWD)
	//wikiPages:=[]string{"75613857","75615168"}
	wikiPages:=strings.Split(IDList, ",")
	var data challenge.ChallangeData
	for _,v:=range wikiPages {
		content := c.GetPageById(v)
		w1 := challenge.GetChallangeData(content.Body.Storage.Value)
		data.Days = append(data.Days, w1.Days...)
		data.Sum = append(data.Sum, w1.Sum...)
	}

	fmt.Print(data)
}
