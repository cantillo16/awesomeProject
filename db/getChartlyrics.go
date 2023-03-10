package db

import (
	"awesomeProject/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetChartlyrics(query string) []byte {
	url := "http://api.chartlyrics.com/apiv1.asmx/SearchLyricDirect?artist=" + query
	fmt.Println("url:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("ERROR EN EL REQUEST")
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var resultado models.Song
	json.Unmarshal(body, &resultado)
	for _, v := range resultado.Results.([]interface{}) {
		InsertoRegistro(v, url)
	}
	if err != nil {
		fmt.Println("ERROR EN EL UNMARSHAL")
		log.Fatal(err)

	}
	return body
}
