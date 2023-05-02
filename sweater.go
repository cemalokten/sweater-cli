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

import "github.com/TwiN/go-color"

type WeatherResponse struct {
	Main struct {
		Temp float64
	}
}

type Config struct {
	APIKey  string `json:"api_key"`
	Town    string `json:"town"`
	Country string `json:"country"`
}

const baseURL = "https://api.openweathermap.org"
const configFileName = "config.json"

func printWithBorder(text string) {
	border := strings.Repeat("#", len(text)+4)
	fmt.Println(border)
	fmt.Printf("# %s #\n", text)
	fmt.Println(border)
}

func main() {
	apiKeyPointer := flag.String("apikey", "", "API Key for OpenWeatherMap")
	townPointer := flag.String("town", "", "Town for the weather forecast")
	countryPointer := flag.String("country", "", "Country for the weather forecast")
	flag.Parse()

	config := readConfig()

	if *apiKeyPointer != "" {
		config.APIKey = *apiKeyPointer
	}
	if *townPointer != "" {
		config.Town = *townPointer
	}
	if *countryPointer != "" {
		config.Country = *countryPointer
	}

	saveConfig(config)

	if config.APIKey == "" {
		log.Fatal("Please provide an Open Weather Map API key using the -APIKey flag")
	}
	if config.Town == "" {
		log.Fatal("Please provide a town using the -town flag")
	}
	if config.Country == "" {
		log.Fatal("Please provide a country using the -country flag")
	}

	URL := fmt.Sprintf("%s/data/2.5/weather?q=%s,%s&appid=%s&units=metric", baseURL, config.Town, config.Country, config.APIKey)

	response, err := http.Get(URL)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code : %d", response.StatusCode)
	}

	var weatherResponse WeatherResponse

	err = json.Unmarshal(body, &weatherResponse)

	if err != nil {
		log.Fatal(err)
	}

	config.Town = strings.ToUpper(config.Town)
	roundedTemp := int(weatherResponse.Main.Temp)
	location := color.With(color.Cyan, "LOCATION:")
	temperature := color.With(color.Red, "TEMPERATURE:")
	wear := color.With(color.Blue, "WEAR:")
	fmt.Println("-------------------------------------")
	fmt.Printf("%s %s, %s \n", location, config.Town, config.Country)
	fmt.Printf("%s %d DEGREES CELSIUS \n", temperature, roundedTemp)
	if weatherResponse.Main.Temp < 0 {
		fmt.Println(wear, "A WARM COAT")
	} else if weatherResponse.Main.Temp < 10 {
		fmt.Println(wear, "LIGHT JACKET")
	} else if weatherResponse.Main.Temp < 18 {
		fmt.Println(wear, "SWEATER")
	} else {
		fmt.Println(wear, "T-SHIRT")
	}
	fmt.Println("-------------------------------------")
}

func saveConfig(config Config) {
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling config: %v", err)
	}

	err = ioutil.WriteFile(configFileName, data, 0644)
	if err != nil {
		log.Fatalf("Error saving config to %s: %v", configFileName, err)
	}
}

func readConfig() Config {
	data, err := ioutil.ReadFile(configFileName)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("Error reading config from %s: %v", configFileName, err)
		}
		return Config{}
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)

	}
	return config
}