package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func doPromotion(client *http.Client, url string, token string) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte{}))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to %s: %v", url, err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Promotion %s failed: %s", url, resp.Status)
		return
	}

	log.Printf("Promotion %s success: %s", url, resp.Status)
}

func getPromotionUrl(crewID int, rid int) string {
	return fmt.Sprintf("https://scapi.rockstargames.com/crew/promote?crewId=%d&rockstarIds=%d&newRankOrder=0", crewID, rid)
}

func main() {

	crewID := flag.Int("crewID", 0, "Crew ID")
	firstRid := flag.Int("firstRid", 0, "First Rockstar ID")
	secondRid := flag.Int("secondRid", 0, "Second Rockstar ID")
	bearerToken := flag.String("token", "", "Bearer token")

	flag.Parse()

	if *crewID == 0 || *firstRid == 0 || *secondRid == 0 || *bearerToken == "" {
		log.Fatalf("Missing required arguments: crewID, firstRid, secondRid, and token must all be provided.")
	}

	firstPromotion := getPromotionUrl(*crewID, *firstRid)
	secondPromotion := getPromotionUrl(*crewID, *secondRid)

	client1 := &http.Client{
		Timeout: time.Second * 10,
	}

	client2 := &http.Client{
		Timeout: time.Second * 10,
	}

	go doPromotion(client1, firstPromotion, *bearerToken)
	go doPromotion(client2, secondPromotion, *bearerToken)

	time.Sleep(5 * time.Second)
}
