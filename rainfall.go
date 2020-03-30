package main

import (
	"strconv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type neorsdRainfallStruct struct {
	Draw            int `json:"draw"`
	RecordsTotal    int `json:"recordsTotal"`
	RecordsFiltered int `json:"recordsFiltered"`
	Data            []struct {
		TrendDataDay string `json:"trend_data_day"`
		RainTotal    float32 `json:"rain_total,string"`
	} `json:"data"`
}

//Get the NEORSD rainfall info, returning the result as a byte string.
func getRainfall(urlToGet string, yearIndex int, month string, logVerbose bool) ([]byte) {
	if logVerbose {
		fmt.Println("In getRainfall")
	}

formData := url.Values{
	//"draw": {"4"},
	// "columns[0][data]": {"trend_data_day"},
	// "columns[0][name]": {""},
	// "columns[0][searchable]": {"true"},
	// "columns[0][orderable]": {"false"},
	// "columns[0][search][value]": {""},
	// "columns[0][search][regex]": {"false"},
	// "columns[1][data]": {"rain_total"},
	// "columns[1][name]": {""},
	// "columns[1][searchable]": {"true"},
	// "columns[1][orderable]": {"false"},
	// "columns[1][search][value]": {""},
	// "columns[1][search][regex]": {"false"},
	// "start": {"0"},
	// "length": {"10"},
	// "search[value]": {""},
	// "search[regex]": {"false"},
	"startingYear": {strconv.Itoa(yearIndex)},
	"rainfallSite": {"Beachwood"},
	//"day": {"1"},
	"month": {month},
	//"fullDate": {"March 1, 2012"},
}

	resp, err := http.PostForm(urlToGet, formData)

	if logVerbose {
		fmt.Println("NEORSD Response:", resp.Status)
	}
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("Error")
		fmt.Println(resp)
	}
	defer resp.Body.Close()
	response, err := ioutil.ReadAll(resp.Body)

	if logVerbose {
		fmt.Println("End of getRainfall")
	}
	return response
}

func main() {
	fmt.Println("GetRainfall v1.0")
	fmt.Println("Specify year to get with -year= command line argument")
	fmt.Println("Output is Year Month Day, rainfall in inches, rounded to nearest hundreth.")
	fmt.Println("")
	yearToGet := flag.Int("year", 2012, "Year of rainfall data to get")
	logVerbose := flag.Bool("verbose", false, "Writes verbose logs for debugging")
	flag.Parse()

	if *logVerbose {
		fmt.Println("Verbose logging enabled.")
	}

	//Define the APIs base URL
	neorsdBaseURL := "https://www.neorsd.org/Rainfall%20Dashboard/dataTableServerSide.php?rainfallDaily"

	//The NEORSD API defines 2012 as year -7, 2020 is year 1.
	//Yes, I could have used the Go time library, but it'd be more work.
	yearOffset := 2019
	yearIndex := *yearToGet - yearOffset
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	for _,month := range months {
		if *logVerbose {
			fmt.Println("Getting data for", yearToGet)
		}

		//Get the rainfall for the month.
		neorsdRainfallJSON := getRainfall(neorsdBaseURL, yearIndex, month, *logVerbose)

		//Unmarshal the monitors list into a struct
		if *logVerbose {
			fmt.Println("Unmarshalling monitors into struct")
		}
		var neorsdRainfallList neorsdRainfallStruct
		if err := json.Unmarshal(neorsdRainfallJSON, &neorsdRainfallList); err != nil {
			panic(err)
		}

		if *logVerbose {
			fmt.Println(neorsdRainfallList)
		}

		for _,day := range neorsdRainfallList.Data {
			fmt.Printf("%v %v %v,%.2f\n", *yearToGet, month, day.TrendDataDay, day.RainTotal)
		}
	}
}
