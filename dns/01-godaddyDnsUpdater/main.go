package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// getPublicIP will get the public ip from ipify.org
func getPublicIP() (string, error) {
	url := "https://api.ipify.org?format=text" // we are using a pulib IP API, we're using ipify here, below are some others
	// https://www.ipify.org
	// http://myexternalip.com
	// http://api.ident.me
	// http://whatismyipaddress.com/api
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed getting public ip %v", err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed reading public ip body %v", err)
	}

	return string(ip), nil
}

func getGodaddyCurrentIP(key string, secret string) (string, error) {
	httpClient := &http.Client{}
	// Create a get request
	req, err := http.NewRequest("GET", "https://api.godaddy.com/v1/domains/erter.org/records/A/dev", nil)
	if err != nil {
		return "", fmt.Errorf("failed creating request: %v", err)
	}

	// Set the correct header for the request.
	req.Header.Set("Authorization", "sso-key "+key+":"+secret)
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed getting response: %v", err)
	}

	defer resp.Body.Close()

	var body []byte
	if resp.StatusCode == http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed reading body : %v", err)
		}

		body = b
	}

	GData := []goDaddyData{}
	json.Unmarshal(body, &GData)

	return GData[0].Data, err

}

func setGodaddyCurrentIP(key string, secret string, apiURL string, gdData string) error {
	httpClient := &http.Client{}
	// Create a new POST request, and prepare it with the POST data.
	req, err := http.NewRequest("PUT", apiURL, strings.NewReader(gdData))
	if err != nil {
		return fmt.Errorf("failed creating request: %v", err)
	}

	req.Header.Set("Authorization", "sso-key "+key+":"+secret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed doing POST: %v", err)
	}

	// "https://api.godaddy.com/v1/domains/${mydomain}/records/A/${myhostname}"

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error: failed reading the response body of the POST: ", err)
	}

	log.Println("POST resp body = ", string(b))

	return nil

}

type goDaddyData struct {
	Data string `json:"data"`
	TTL  int    `json:"ttl"`
}

// run will orchestrate the checks for finding out if ip's are changed,
// and change it at godaddy if changed.
func run(key string, secret string) error {
	// Get the current public ip of your connection.
	pIP, err := getPublicIP()
	if err != nil {
		log.Println("error: public ip: ", err)
		return err
	}
	log.Printf("My IP is:%s\n", pIP)

	// get current ip registered at godaddy.
	gIP, err := getGodaddyCurrentIP(key, secret)
	log.Printf("Current godaddy ip for dev.erter.org = %v\n", gIP)

	// If the current public ip and the registered dns ip at godaddy are not the same,
	// change the value in the godaddy dns record.
	if pIP != gIP {
		log.Println("* The ip's are different")
		gd := goDaddyData{
			Data: pIP,
			TTL:  600,
		}

		// Create the data for header that will be changed
		gdArray := []goDaddyData{gd}
		gdJSON, err := json.Marshal(gdArray)
		if err != nil {
			log.Println("error: json marshal failed")
		}

		apiURL := "https://api.godaddy.com/v1/domains/erter.org/records/A/dev"

		// do the api call to set the new ip
		err = setGodaddyCurrentIP(key, secret, apiURL, string(gdJSON))
		if err != nil {
			log.Println("error: setGodaddyCurrent ip = ", err)
		}

		return nil
	}

	log.Println("The ip's where the same, doing nothing")

	return nil
}

func main() {
	auth := flag.String("auth", "env", `Use "env" or "flag" for way to get key and secret.\n
	if value chosen is "flag", use the -key and -secret flags.\
	if value chosen is "env", set the env variables "goddaddykey" and "godaddysecret"
	`)
	key := flag.String("key", "", "the key you got at https://developer.godaddy.com/keys")
	secret := flag.String("secret", "", "the secret you got at https://developer.godaddy.com/keys")
	flag.Parse()

	switch *auth {
	case "env":
		*key = os.Getenv("godaddykey")
		*secret = os.Getenv("godaddysecret")
		if *key == "" || *secret == "" {
			log.Println("method env chosen, and you need to set key and secret")
			return
		}
	case "flag":
		if *key == "" || *secret == "" {
			log.Println("method flag chosen, and you need to set key and secret")
			return
		}
	}

	// Run the checking, and eventually edit dns record at godaddy.
	run(*key, *secret)

}
