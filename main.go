package main

// import (
//     "github.com/gin-gonic/gin"
//     "net/http"
// )

// func main() {
//     router := gin.Default()

//     router.GET("/ping", func(c *gin.Context) {

//         c.String(http.StatusOK, "pong")
//     })

//     router.Run()
// }

import (
	"fmt"
	//"strings"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

var (
	commandPrefix string
	botID         string
)

func main() {

	startPing()
	startDiscord()
}

func startPing() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {

		c.String(http.StatusOK, "pong")
	})

	router.Run()
}

func startDiscord() {
	discord, err := discordgo.New("Bot NTEwNDc1NTk5OTI0NDI4ODMw.Dsd0uw.p13kOVtBrfIsi_bOqK3tQq80W58")
	errCheck("error creating discord session", err)
	user, err := discord.User("@me")
	errCheck("error retrieving account", err)

	botID = user.ID
	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err = discord.UpdateStatus(0, "A friendly helpful bot!")
		if err != nil {
			fmt.Println("Error attempting to set my status")
		}
		servers := discord.State.Guilds
		fmt.Printf("SuperAwesomeOmegaTutorBot has started on %d servers", len(servers))
	})

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	commandPrefix = "!"

	<-make(chan struct{})

}

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	if user.ID == botID || user.Bot {
		//Do nothing because the bot is talking
		return
	}

	// content := message.Content //esta declarado pero no lo usa

	fmt.Printf("Message: %+v || From: %s\n", message.Message, message.Author)
}
