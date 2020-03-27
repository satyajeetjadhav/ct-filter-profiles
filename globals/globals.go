package globals

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

var CSVFilePath *string
var JSONFilePath *string
var AccountID *string
var AccountPasscode *string
var AccountToken *string
var Region *string
var DryRun *bool
var StartTs *float64

//var AutoConvert *bool

type arrayFlags []string

func (i *arrayFlags) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func Init() bool {
	CSVFilePath = flag.String("csv", "", "Absolute path to the csv file")
	JSONFilePath = flag.String("json", "", "Absolute path to the json file")
	StartTs = flag.Float64("startTs", 0, "Start timestamp for events upload")
	AccountID = flag.String("id", "", "CleverTap Account ID")
	AccountPasscode = flag.String("p", "", "CleverTap Account Passcode")
	AccountToken = flag.String("tk", "", "CleverTap Account Token")
	Region = flag.String("r", "eu", "The account region, either eu, in, sk, or sg, defaults to eu")
	DryRun = flag.Bool("dryrun", false, "Do a dry run, process records but do not upload")
	//AutoConvert = flag.Bool("autoConvert", false, "automatically covert property value type to number for number entries")
	flag.Parse()

	if *Region != "eu" && *Region != "in" && *Region != "sk" && *Region != "sg" {
		log.Println("Region can be either eu, in, sk, or sg")
		return false
	}

	return true
}

var Schema map[string]string

func ParseSchema(file *os.File) bool {
	/**
	{
		"key": "Float",
		"key 1": "Integer",
		"key 2": "Number",
		"key 3": "Float[]",
		"key 4": "Integer[]",
		"key 5": "String[]",
		"key 6": "Boolean[]"
	}
	*/
	err := json.NewDecoder(file).Decode(&Schema)
	if err != nil {
		log.Println(err)
		log.Println("Unable to parse schema file")
		return false
	}
	return true
}
