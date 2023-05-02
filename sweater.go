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

