package controller

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Id         uuid.UUID `json:"id"`
	Author   string    `json:"author"`
	Text       string    `json:"text"`
	CategoryId string    `json:"category_id"`
	PostedAt   time.Time `json:"posted_at"`
}

func (u *UserController) GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//from, err := strconv.Atoi(r.URL.Query()["from"][0])
	//if err != nil {
	//	log.Printf("Cannot convert params from url into int: %v\n", err)
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//
	//number, err := strconv.Atoi(r.URL.Query()["number"][0])
	//if err != nil {
	//	log.Printf("Cannot convert params from url into int: %v\n", err)
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	messages, err := u.getMessages()
	if err != nil {
		log.Printf("Cannot extract messages from database: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(messages)
	if err != nil {
		log.Printf("Error encoding messages to json, error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u *UserController) getMessages() ([]Message, error) {
	rows, err := u.db.Query(
		"SELECT m.id, u.name, m.text, c.name, m.posted_at "+
			"FROM messages m "+
			"INNER JOIN categories c ON m.category_id = c.id "+
			"INNER JOIN users u ON m.author_id = u.id "+
			//"ORDER BY m.posted_at DESC "+
			//"LIMIT $1 OFFSET $2",
			"ORDER BY m.posted_at DESC ",
			//number, from,
		)
	if err != nil {
		return make([]Message, 0), err
	}

	var messages []Message
	for rows.Next() {
		msg := Message{}
		err := rows.Scan(&msg.Id, &msg.Author, &msg.Text, &msg.CategoryId, &msg.PostedAt)
		if err != nil {
			return make([]Message, 0), err
		}
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return make([]Message, 0), err
	}

	if len(messages) == 0 {
		return make([]Message, 0), nil
	}
	return messages, nil
}

func (u *UserController) AddMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
