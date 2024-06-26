/*
Copyright © 2024 EpykLab contact@epyklab.com
*/

package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/EpykLab/chx/cmd/utils/configs"
	"github.com/charmbracelet/log"
)

const (
	abuseipdbURL string = "https://api.abuseipdb.com/api/v2/check"
)

func GetIPInfo(ip string) {

	var maxAgeInDays string = "90"

	// if len(os.Args) < 2 {
	// 	fmt.Println("Provide IP address as first argument\n\nEXAMPLE: chx 170.205.29.2")
	// 	os.Exit(1)
	// }

	// if len(os.Args) == 3 {
	// 	maxAgeInDays = os.Args[2]
	// } else {
	// 	maxAgeInDays = "90"
	// }

	// ipAddr := os.Args[1]

	key, err := configs.ReadConfig()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	apiKey := key.ABUSE_DB_KEY

	req, _ := http.NewRequest("GET", abuseipdbURL, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Key", apiKey)

	query := req.URL.Query()
	query.Add("ipAddress", ip)
	query.Add("maxAgeInDays", maxAgeInDays)

	req.URL.RawQuery = query.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Unable to make request: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Unable to parse response: %v", err)
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal(err)
	}

	jsByte, err := json.MarshalIndent(m, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	output := string(jsByte)
	fmt.Println(output)
}
