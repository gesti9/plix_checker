package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	c := New("7213377601:AAEi0GVe_reXxbYemvaA9GFo4vZ7weqsjYE")

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	checkSites(c)

	for {
		select {
		case <-ticker.C:
			checkSites(c)
		}
	}
}

func checkSites(c *Client) {
	front(c)
	backend(c)
}

func front(c *Client) {
	url := "https://plix.kz"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		c.SendMessage("🔴 Plix.kz  статуc: "+strconv.Itoa(resp.StatusCode), int64(-4245258605))
	}
}

func backend(c *Client) {
	url := "https://plix.kz/api/v1/swagger-ui/index.html"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		c.SendMessage("🔴 https://plix.kz/api/v1/swagger-ui/index.html статус: "+strconv.Itoa(resp.StatusCode), int64(-4245258605))
	}
}

type Client struct {
	bot *tgbotapi.BotAPI
}

func New(apiKey string) *Client {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(err)
	}

	return &Client{
		bot: bot,
	}
}

func (c *Client) SendMessage(text string, chatId int64) error {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = "Markdown"
	_, err := c.bot.Send(msg)
	return err
}
