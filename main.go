package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func getToken() string {
	f, err := os.Open("./tokens.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	token := fileScanner.Text()
	fmt.Println(token)
	//fmt.Println(reflect.TypeOf(token))
	//fmt.Printf("%s\n", token)
	return token

}

func main() {
	sess, err := discordgo.New("Bot " + getToken())
	if err != nil {
		log.Fatal(err)
	}

	//ignore bot's own messages
	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		//respond to message content
		if m.Content == "hello" {
			s.ChannelMessageSend(m.ChannelID, "world!")
			getUId("stagefuel")
			getFavAnime("stagefuel")
			fmt.Printf("\n")
			getFavManga("stagefuel")
			fmt.Printf("\n")
			getFavChara("stagefuel")
		}

	})

	//set intents
	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = sess.Open() //open session
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()
	fmt.Println("戦闘モード起動！！")

	//listen for interrupt and shutdown
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
