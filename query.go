package lasav

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

const (
	Url = "https://api.magicthegathering.io/v1/"
)

func DoRequest() {
	u := "https://api.magicthegathering.io/v1/cards/386616"
	card, err := fetchCards(u)
	if err != nil {
		fmt.Println("error ", err)
	}

	jsonString, _ := json.Marshal(card)
	fmt.Println(string(jsonString))

	// convert json to struct
	s := Card{}
	json.Unmarshal(jsonString, &s)
	fmt.Println(s)

	// for k, v := range card["card"].(map[string]interface{}) {
	// 	fmt.Println("key: ", k)
	// 	fmt.Println("value: ", v)
	// }

}

// func fetchCards(url string) (map[string]interface{}, error) {
// 	var card map[string]interface{}
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return card, err
// 	}

// 	bdy := resp.Body
// 	defer bdy.Close()

// 	if err := json.NewDecoder(bdy).Decode(&card); err != nil {
// 		return card, err
// 	}

// 	return card, nil

// }

func fetchCards(url string) (Card, error) {
	card := cardResponse{}
	resp, err := http.Get(url)
	if err != nil {
		return card.Card, err
	}

	bdy := resp.Body
	defer bdy.Close()

	if err := json.NewDecoder(bdy).Decode(&card); err != nil {
		return card.Card, err
	}
	if !reflect.DeepEqual(card.Card, Card{}) {
		return card.Card, nil
	}
	return card.Cards[0], nil

}

type cardResponse struct {
	Card  Card   `json:"card"`
	Cards []Card `json:"cards"`
}
