package main

//this isnt working when autostarting due to something wrong with being unable
//to find the module gjson when started outside of ~/config/waybar/modules

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var golemIcon string = "🗿"
var walletIcon string = "👛"
var marketIcon string = "📈"
var trueWalletValIcon string = "💰"

func main() {
	walletVal := getWalletVal()
	marketVal := getMarketVal()
	trueWalletVal := walletVal * marketVal

	json.NewEncoder(os.Stdout).Encode(struct {
		Text    string `json:"text"`
		Alt     string `json:"alt"`
		Tooltip string `json:"tooltip"`
	}{
		fmt.Sprintf(golemIcon + " " +
			getStatus() + " " +
			walletIcon + " " +
			fmt.Sprintf("%.2f", walletVal) + " " +
			marketIcon + " $" +
			fmt.Sprintf("%.2f", marketVal) + " " +
			trueWalletValIcon + " $" +
			fmt.Sprintf("%.2f", trueWalletVal)),
		fmt.Sprintf("desu2"),
		fmt.Sprintf("Click to start/stop golem process."),
	})
}

func getStatus() string {
	args := []string{"-C", "yagna", "-o", "%cpu="}
	psYagna, _ := exec.Command("ps", args...).Output()
	if len(psYagna) == 0 {
		return "Offline"
	} else {
		args := []string{"-C", "vmrt", "-o", "%cpu="}
		psVMRT, _ := exec.Command("ps", args...).Output()
		//floatPsVMRT, _ := strconv.ParseFloat(string(psVMRT), 64)
		if len(psVMRT) == 0 { //Check that vmrt is running & cpu is >85
			return "Online"
		} else {
			return "Working"
		}
	}
}

func getWalletVal() float64 {
	/* I really would have rather just used to zksync API here, but
	they dont have a driver for golang *yet*. This should be replaced
	as soon as they release it... along with a stable API version.*/
	args := []string{"status"}
	gStatus, _ := exec.Command("golemsp", args...).Output()
	words := strings.Fields(string(gStatus))
	s, _ := strconv.ParseFloat(words[58], 64)
	return s
}

func getMarketVal() float64 {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := url.Values{}
	q.Add("slug", "golem")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", "YOUR_API_TOKEN")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request to server")
		os.Exit(1)
	}
	respBody, _ := ioutil.ReadAll(resp.Body)
	respValue := gjson.Get(string(respBody), "data.1455.quote.USD.price")
	return respValue.Float()
}
