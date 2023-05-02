# Sweater CLI

A simple command-line interface (CLI) written in Go, which outputs coat, sweater or t-shirt depending on your location and weather.

![Sweater CLI (London, UK)](https://user-images.githubusercontent.com/60609268/235652837-32661ee2-78c2-4cc3-94c3-99541d79fda3.png)

Requirements
---

*   Go 1.13 or later
*   An API key for OpenWeatherMap

Usage
---

1.  Clone the repository:

`git clone https://github.com/cemalokten/sweater`

2. Run the program with the following flags:

`go run sweater.go -apikey YOUR_API_KEY -town TOWN_NAME -country COUNTRY_CODE`

For example:

`go run sweater.go -apikey 1234567890 -town London -country UK`

Configuration
---

The program stores the API key, town, and country in a `config.json` file. If you provide these values as flags, the program will update the configuration file accordingly.

You can also edit the `config.json` file manually to set the default API key, town, and country for the application:

```json
{
  "api_key": "123456789",
  "town": "London",
  "country": "UK"
}
```

Replace `YOUR_API_KEY`, `TOWN_NAME`, and `COUNTRY_CODE` with your desired values.
