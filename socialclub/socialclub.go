package socialclub

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	webClient = &http.Client{
		Timeout: 10 * time.Second,
	}
)

func GetCrewHierarchy(crewID int, token string) (CrewHierarchy, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://scapi.rockstargames.com/crew/ranksWithMembership?crewId=%d&onlineService=sc&searchTerm=&memberCountToRetrieve=5", crewID), nil)
	if err != nil {
		return CrewHierarchy{}, err
	}

	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Authorization", token)

	resp, err := webClient.Do(req)
	if err != nil {
		return CrewHierarchy{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CrewHierarchy{}, err
	}

	var crewHierarchy CrewHierarchy
	err = json.Unmarshal(body, &crewHierarchy)
	if err != nil {
		return CrewHierarchy{}, err
	}

	return crewHierarchy, nil
}

func GetBasicCrewInfo(name string, token string) (CrewBasicInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://scapi.rockstargames.com/crew/byname?name=%s", name), nil)
	if err != nil {
		return CrewBasicInfo{}, err
	}

	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Authorization", token)

	resp, err := webClient.Do(req)
	if err != nil {
		return CrewBasicInfo{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CrewBasicInfo{}, err
	}

	var crewBasicInfo CrewBasicInfo
	err = json.Unmarshal(body, &crewBasicInfo)
	if err != nil {
		return CrewBasicInfo{}, err
	}

	return crewBasicInfo, nil
}
