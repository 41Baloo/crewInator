package main

import (
	"crewInator/socialclub"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/atotto/clipboard"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/valyala/fasthttp"
)

func doPromotion(wg *sync.WaitGroup, url string, token string) {
	defer wg.Done()

	if url == "" {
		return
	}

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")
	req.Header.Set("Authorization", token)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	if err := fasthttp.Do(req, resp); err != nil {
		log.Printf("Error sending request to %s: %v", url, err)
		return
	}

	statusCode := resp.StatusCode()
	body := resp.Body()

	if statusCode != 200 {
		log.Printf("Promotion %s failed: Status Code: %d, Response: %s", url, statusCode, body)
		return
	}
}

func getPromotionUrl(crewID int, rid int) string {
	return fmt.Sprintf("https://scapi.rockstargames.com/crew/promote?crewId=%d&rockstarIds=%d&newRankOrder=0", crewID, rid)
}

func handleCrewPromotion(crewName string, bearerToken string) error {
	crewName = strings.ReplaceAll(crewName, " ", "_")
	bearerToken = "Bearer " + bearerToken

	basicCrewInfo, err := socialclub.GetBasicCrewInfo(crewName)
	if err != nil || !basicCrewInfo.Status {
		return errors.New("failed to convert crew name to crewID")
	}

	crewHierarchy, err := socialclub.GetCrewHierarchy(basicCrewInfo.CrewID)
	if err != nil || !crewHierarchy.Status {
		return errors.New("failed to fetch crew hierarchy")
	}

	crewRanks := crewHierarchy.CrewRanks
	if len(crewRanks) < 2 {
		return errors.New("failed to find commissioner rank")
	}

	comissionerMembers := crewRanks[1].RankMembers

	// Doing this dirty shit so goroutines don't have to do any additional progressing and send at nearly the exact same time
	promotion1 := ""
	promotion2 := ""
	promotion3 := ""
	promotion4 := ""
	promotion5 := ""

	lengthRids := len(comissionerMembers)

	if lengthRids > 0 {
		promotion1 = getPromotionUrl(basicCrewInfo.CrewID, comissionerMembers[0].RockstarID)
	}
	if lengthRids > 1 {
		promotion2 = getPromotionUrl(basicCrewInfo.CrewID, comissionerMembers[1].RockstarID)
	}
	if lengthRids > 2 {
		promotion3 = getPromotionUrl(basicCrewInfo.CrewID, comissionerMembers[2].RockstarID)
	}
	if lengthRids > 3 {
		promotion4 = getPromotionUrl(basicCrewInfo.CrewID, comissionerMembers[3].RockstarID)
	}
	if lengthRids > 4 {
		promotion5 = getPromotionUrl(basicCrewInfo.CrewID, comissionerMembers[4].RockstarID)
	}

	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(5)

		go doPromotion(&wg, promotion1, bearerToken)
		go doPromotion(&wg, promotion2, bearerToken)
		go doPromotion(&wg, promotion3, bearerToken)
		go doPromotion(&wg, promotion4, bearerToken)
		go doPromotion(&wg, promotion5, bearerToken)
	}

	wg.Wait()

	return nil
}

func main() {
	jsCode := `function gT(){let t=decodeURIComponent(document.cookie).split(";"),e="BearerToken=";for(let o of t)if((o=o.trim()).startsWith(e)){let i=o.substring(e.length,o.length);navigator.clipboard.writeText(i).then(()=>{alert("Copied To Clipboard")}).catch(t=>{console.log("Copy this:",i),alert("Copy this: "+i)});return}alert("You don't seem to be logged in")}gT();`

	if runtime.GOOS == "windows" { // windows don't support colors, for the most part :[
		color.NoColor = true
	}

	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	fmt.Println(bold("Welcome to CrewInator! This tool helps you promote leaders in your Social Club crew."))
	fmt.Println()
	fmt.Println("How to use CrewInator:")
	fmt.Printf("%s. Log in to your Social Club account with leader privileges.\n", green("1"))
	fmt.Printf("%s. Promote up to 5 people to 'commissioner' rank.\n", green("2"))
	fmt.Printf("%s. Go to your crew's page and get the crew name from the URL.\n", green("3"))
	fmt.Printf("%s. Open the browser console (right-click -> Inspect -> Console) and paste this JS code:\n", green("4"))
	fmt.Println()
	fmt.Println(blue(jsCode))
	fmt.Println()
	if clipboardContent, _ := clipboard.ReadAll(); strings.TrimSpace(clipboardContent) == "" {
		clipboard.WriteAll(jsCode)
		fmt.Println(green("(The JavaScript code has been copied to your clipboard.)"))
	} else {
		fmt.Println(blue("(JavaScript code not copied because your clipboard contains other data.)"))
	}
	fmt.Println()
	fmt.Printf("%s. Copy the output Bearer Token, and enter it into the prompt here.\n", green("5"))
	fmt.Println()

	prompt := promptui.Prompt{
		Label: "Enter the crew name",
	}

	crewName, err := prompt.Run()
	if err != nil {
		log.Fatalf(red("Prompt failed %v\n"), err)
	}

	tokenPrompt := promptui.Prompt{
		Label: "Enter the bearer token",
		Mask:  '*',
	}

	bearerToken, err := tokenPrompt.Run()
	if err != nil {
		log.Fatalf(red("Prompt failed %v\n"), err)
	}

	if !strings.HasPrefix(bearerToken, "eyJ") { // 'eyJ' is base64 for '{"'
		fmt.Println(red("Your input does not seem to be a BearerToken. Make sure you copied the JavaScript's output correctly. BearerTokens start with 'eyJ'."))
		os.Exit(1)
	}

	fmt.Println()

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()

	err = handleCrewPromotion(crewName, bearerToken)

	s.Stop()

	if err != nil {
		fmt.Println(red(fmt.Sprintf("Failed to do promotions: %s", err.Error())))
		os.Exit(1)
	}

	fmt.Println(green("Promotion requests sent. Depending on timing, multiple leaders may be promoted."))
	fmt.Println("If you encounter issues, try the hosted service at https://crew.dudx.info/")

	select {}
}
