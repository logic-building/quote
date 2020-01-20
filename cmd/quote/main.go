package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gookit/color"

	"quote/pkg/api"
	"quote/pkg/env"
	"quote/pkg/event"
	"quote/pkg/quote"
)

func main() {
	serverRunTimeInMin := env.GetIntWithDefault("SERVER_RUN_DURATION_MIN", 5)
	serverRunTimeInHour := env.GetIntWithDefault("SERVER_RUN_DURATION_HOUR", 2)

	//pic := env.GetBoolWithDefault("PIC", false)
	//img := env.GetBoolWithDefault("IMG", false)
	//img2 := env.GetBoolWithDefault("IMAGE", false)

	quote := quote.QuoteForTheDay()

	wordList := strings.Split(quote, " ")
	wordListLen := len(wordList)

	red := color.FgRed.Render
	green := color.FgGreen.Render
	yellow := color.FgYellow.Render
	blue := color.FgBlue.Render
	cyan := color.FgCyan.Render

	fmt.Println("Quote for the day")
	fmt.Println("-----------------")
	fmt.Printf("\t\"")
	for i := 0; i < wordListLen; i++ {
		if i%15 == 0 {
			fmt.Printf("%s ", cyan(wordList[i]))
		} else if i%10 == 0 {
			fmt.Printf("%s ", blue(wordList[i]))
		} else if i%5 == 0 {
			fmt.Printf("%s ", yellow(wordList[i]))
		} else if i%2 == 0 {
			fmt.Printf("%s ", red(wordList[i]))
		} else {
			fmt.Printf("%s ", green(wordList[i]))
		}
	}
	fmt.Printf("\"")
	fmt.Println("\n")
	fmt.Println("Link for the next quote")
	fmt.Printf("%s", blue("http://localhost:1922/quotes-devotional\n"))
	fmt.Printf("%s", blue("http://localhost:1922/quotes-motivational\n"))

	todayEvents := event.TodayEvents()
	if len(todayEvents) > 0 {
		today := time.Now()
		todayDateStr := today.Format("Mon 2006-01-2")
		fmt.Printf("\nToday %v is auspicious day\n", todayDateStr)
		fmt.Println("-------------")
	}

	for _, event := range todayEvents {
		event.DisplayEvent()
	}

	//if pic || img || img2 {
	//	listDir("/image")
	//	image.DisplayImage("./image/competitionWithMySelf.jpg")
	//}

	currentTime := time.Now()
	currentTime = currentTime.Add(time.Duration(serverRunTimeInHour) * time.Hour)
	currentTime = currentTime.Add(time.Duration(serverRunTimeInMin) * time.Minute)

	const layout = "Jan 2, 2006 at 3:04pm MST"

	fmt.Printf("\n\nServer will be quit in %d hour and %d minutes at %v", serverRunTimeInHour, serverRunTimeInMin, currentTime.Format(layout))
	fmt.Printf("\n or press CTRL+C or CTRL +D to exit and stop docker container - 'quote' using commands- docker ps and docker stop \n")
	const httpPort int = 1922
	apiServer := api.NewServer(httpPort)
	go func() {
		apiServer.Run()
	}()

	time.Sleep(time.Duration(serverRunTimeInMin) * time.Minute)
	time.Sleep(time.Duration(serverRunTimeInHour) * time.Hour)
	apiServer.Close()
}
