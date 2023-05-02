# Sweater CLI

A simple command-line interface (CLI) application to fetch the weather forecast for a specified town and country using the OpenWeatherMap API.

Prerequisites
-------------

*   Go 1.13 or later
*   An API key for OpenWeatherMap



Usage
-----

1.  Clone the repository:


`git clone https://github.com/cemalokten/sweater`

2. Run the program with the following flags:

`go run sweater.go -apikey YOUR_API_KEY -town TOWN_NAME -country COUNTRY_CODE`

Replace `YOUR_API_KEY` with your OpenWeatherMap API key, `TOWN_NAME` with the name of the town you want the forecast for, and `COUNTRY_CODE` with the 2-letter country code.

For example:

`go run sweater.go -apikey 1234567890 -town London -country UK`

Configuration
-------------

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
