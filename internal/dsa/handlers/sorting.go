package handlers

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gofiber/contrib/websocket"
)

const (
	NUM_OF_NUMS   = 10
	NEW_LIST_MSG  = "new"
	ASC_LIST_MSG  = "asc"
	DESC_LIST_MSG = "desc"
)

var (
	returnHolder = ValueHolderThing{}
)

type ValueHolderThing struct {
	Values        []int `json:"values"`
	I_PlaceHolder int   `json:"i"`
	J_PlaceHolder int   `json:"j"`
}

type SortMessage struct {
	Message    string `json:"message"`
	Order      string `json:"order"`
	LISTLENGTH int    `json:"list_length"`
}

func BubbleSort(c *websocket.Conn) {

	for {
		fmt.Println("ITERATION")
		var incomingMessage SortMessage
		if err := c.ReadJSON(&incomingMessage); err != nil {
			log.Fatal(err)
		}

		switch incomingMessage.Message {
		case NEW_LIST_MSG:
			returnHolder.Values = getNewList(incomingMessage.LISTLENGTH)
			for i := 0; i < len(returnHolder.Values)-1; i++ {
				for j := 0; j < len(returnHolder.Values)-i-1; j++ {
					var greaterThanCheck bool = returnHolder.Values[j+1] < returnHolder.Values[j]
					switch incomingMessage.Order {
					case ASC_LIST_MSG:
						break
					case DESC_LIST_MSG:
						greaterThanCheck = !greaterThanCheck
						break
					}

					if greaterThanCheck {
						returnHolder.Values[j], returnHolder.Values[j+1] = returnHolder.Values[j+1], returnHolder.Values[j]
					}
					if err := c.WriteJSON(returnHolder); err != nil {
						log.Println(err)
					}
				}
			}
			break
		default:
			log.Println("Invalid message")
			if err := c.WriteMessage(1, []byte("Invalid message")); err != nil {
				log.Println(err)
			}
			break
		}
	}
}

func getNewList(list_length int) []int {

	if list_length > 1000 {
		list_length = NUM_OF_NUMS
	}

	var list []int = make([]int, list_length)
	for i := range list_length {
		list[i] = rand.Intn(list_length)
	}
	return list
}
