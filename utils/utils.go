package utils

import (
	"assignment-3/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	cron "github.com/robfig/cron/v3"
)

const max = 100
const min = 1
const fileName = "data.json"

func WriteRandomData() models.Data {
	var data models.Data
	rand.Seed(time.Now().UnixNano())
	data.Status.Water = uint(rand.Intn(max-min) + min)
	data.Status.Wind = uint(rand.Intn(max-min) + min)
	return data
}

func WriteJSON(){
	fmt.Println("write json called")
	data := WriteRandomData()

	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(fileName, file, 0644)
}

func ReadJSON() {
	var data models.Data
	jsonFile, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully opened file")

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &data)
	fmt.Println(data)

	defer jsonFile.Close()
}

func CronJob() {
	fmt.Println("cron job called")
	WriteJSON()
	c := cron.New()

	defer c.Stop()

	c.AddFunc("@every 15s", WriteJSON)
	c.Start()
	time.Sleep(time.Minute * 15)
}
