package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotState struct {
	LastMessageTime  time.Time `json:"last_message_time"`
	WaitingForReply  bool      `json:"waiting_for_reply"`
	MessageSentTime  time.Time `json:"message_sent_time"`
	NotificationSent bool      `json:"notification_sent"`
}

type Config struct {
	BotToken        string
	PrimaryUserID   int64
	SecondaryUserID int64
	StateFile       string
}

func loadConfig() Config {
	return Config{
		BotToken:        os.Getenv("BOT_TOKEN"),        // Set your bot token
		PrimaryUserID:   parseChatID(os.Getenv("PRIMARY_USER_ID")),   // Primary user chat ID
		SecondaryUserID: parseChatID(os.Getenv("SECONDARY_USER_ID")), // Secondary user chat ID
		StateFile:       "bot_state.json",
	}
}

func parseChatID(s string) int64 {
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Printf("Error parsing chat ID %s: %v", s, err)
		return 0
	}
	return id
}

func loadState(filename string) BotState {
	var state BotState
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Could not read state file, starting fresh: %v", err)
		return BotState{LastMessageTime: time.Now().Add(-72 * time.Hour)} // Start immediately
	}
	
	err = json.Unmarshal(data, &state)
	if err != nil {
		log.Printf("Could not parse state file, starting fresh: %v", err)
		return BotState{LastMessageTime: time.Now().Add(-72 * time.Hour)}
	}
	
	return state
}

func saveState(filename string, state BotState) error {
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	
	return ioutil.WriteFile(filename, data, 0644)
}

func main() {
	config := loadConfig()
	
	if config.BotToken == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}
	
	if config.PrimaryUserID == 0 {
		log.Fatal("PRIMARY_USER_ID environment variable is required")
	}
	
	if config.SecondaryUserID == 0 {
		log.Fatal("SECONDARY_USER_ID environment variable is required")
	}
	
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}
	
	bot.Debug = false
	log.Printf("Bot started as %s", bot.Self.UserName)
	
	// Load state
	state := loadState(config.StateFile)
	
	// Set up update configuration with a timeout
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	
	updates := bot.GetUpdatesChan(u)
	
	// Timer for checking intervals
	ticker := time.NewTicker(1 * time.Minute) // Check every minute
	defer ticker.Stop()
	
	for {
		select {
		case update := <-updates:
			if update.Message != nil {
				handleMessage(bot, update.Message, &state, config)
				saveState(config.StateFile, state)
			}
			
		case <-ticker.C:
			checkAndSendMessages(bot, &state, config)
			saveState(config.StateFile, state)
		}
	}
}

func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message, state *BotState, config Config) {
	// Only handle messages from the primary user
	if message.Chat.ID == config.PrimaryUserID {
		log.Printf("Received message from primary user: %s", message.Text)
		
		// If we were waiting for a reply, mark as replied
		if state.WaitingForReply {
			state.WaitingForReply = false
			state.NotificationSent = false
			log.Println("Primary user replied, resetting wait state")
		}
		
		// Update last message time to reset the 3-day cycle
		state.LastMessageTime = time.Now()
		
		// Send acknowledgment
		msg := tgbotapi.NewMessage(config.PrimaryUserID, "Got it! I'll check back with you in 3 days.")
		bot.Send(msg)
	}
}

func checkAndSendMessages(bot *tgbotapi.BotAPI, state *BotState, config Config) {
	now := time.Now()
	
	// Check if we need to send the initial message (every 3 days)
	if !state.WaitingForReply && now.Sub(state.LastMessageTime) >= 72*time.Hour {
		log.Println("Sending initial message to primary user")
		
		msg := tgbotapi.NewMessage(config.PrimaryUserID, 
			"Hi! This is your 3-day check-in. Please reply to let me know you're okay.")
		
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Error sending message to primary user: %v", err)
			return
		}
		
		state.WaitingForReply = true
		state.MessageSentTime = now
		state.NotificationSent = false
	}
	
	// Check if we need to notify the secondary user (48 hours after initial message)
	if state.WaitingForReply && !state.NotificationSent && 
		now.Sub(state.MessageSentTime) >= 48*time.Hour {
		
		log.Println("Sending notification to secondary user")
		
		msg := tgbotapi.NewMessage(config.SecondaryUserID, 
			fmt.Sprintf("Alert: Primary user (ID: %d) hasn't responded to the check-in message sent 48 hours ago.", 
				config.PrimaryUserID))
		
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Error sending notification to secondary user: %v", err)
			return
		}
		
		state.NotificationSent = true
		
		// Reset the cycle - start counting 3 days from the original message time
		state.LastMessageTime = state.MessageSentTime
		state.WaitingForReply = false
	}
}