# 🤖 Are-you-dead-yet?

*The passive-aggressive Telegram bot that cares about your well-being (whether you like it or not)*

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![Telegram](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://telegram.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

## 🎭 What is Are-you-dead-yet?

`Are-You-Dead-Yet?` is the world's most persistent digital companion that refuses to let you disappear into the void! This little Go-powered menace will:

- 📅 Poke you every 3 days like that friend who "just wants to make sure you're alive"
- ⏰ Give you exactly 48 hours to prove your existence 
- 🚨 Tattle on you to your designated emergency contact if you ignore it
- 💾 Remember everything (because it has trust issues and saves state to disk)
- 🏃‍♂️ Run on a potato server (0.5GB RAM? No problem!)

Perfect for helicopter parents, worried friends, or anyone who needs a digital nudge to stay connected!

## 🚀 Features

- **Ultra Low Resource Usage**: Runs happily on servers with 1 core CPU and 0.5GB RAM
- **Persistent Memory**: Saves state to JSON file - survives server restarts like a digital cockroach
- **Smart Notifications**: 3-day check-ins with 48-hour grace periods
- **Escalation Protocol**: Automatically notifies emergency contact when needed
- **Zero Dependencies Drama**: Uses only the most essential Go packages

## 📋 Prerequisites

- Go 1.19 or later
- A Telegram Bot Token (courtesy of [@BotFather](https://t.me/botfather))
- At least 2 people who care about each other's existence
- A server that's barely alive (but hey, aren't we all?)

## 🛠️ Installation

1. **Clone this digital nag:**
   ```bash
   git clone https://github.com/chashtager/are-you-dead-yet.git
   cd are-you-dead-yet
   ```

2. **Get your Bot Token:**
   - Chat with [@BotFather](https://t.me/botfather) on Telegram
   - Create a new bot: `/newbot`
   - Name it something embarrassing like "YourNameDeadCheck"
   - Save that precious token!

3. **Find your Chat IDs:**
   - Message [@userinfobot](https://t.me/userinfobot) to get your chat ID
   - Get IDs for both the person you want to nag and their emergency contact

4. **Set up environment variables:**
   ```bash
   export BOT_TOKEN="your_bot_token_here"
   export PRIMARY_USER_ID="123456789"    # The person getting nagged
   export SECONDARY_USER_ID="987654321"  # The emergency contact/snitch
   ```

5. **Install dependencies and run:**
   ```bash
   go mod tidy
   go run main.go
   ```

## 🎮 Usage

Once running, Are-you-dead-yet will:

1. **Send a check-in message** every 3 days to the primary user
2. **Wait patiently** (but judgmentally) for 48 hours
3. **Tattle to the secondary user** if ignored
4. **Repeat forever** because it has nothing better to do

### Example Conversation:

**Are-you-dead-yet?:** "Hi! This is your 3-day check-in. Please reply to let me know you're okay."

**You:** "I'm alive, stop bothering me!"

**Are-you-dead-yet?:** "Got it! I'll check back with you in 3 days." *evil laugh*

## 📁 File Structure

```
are-you-dead-yet/
├── main.go           # The brain of the operation
├── go.mod           # Go module definition
├── bot_state.json   # Persistent memory (auto-generated)
└── README.md        # This masterpiece
```

## ⚙️ Configuration

All configuration is done via environment variables:

| Variable | Description | Required |
|----------|-------------|----------|
| `BOT_TOKEN` | Your Telegram bot token | ✅ |
| `PRIMARY_USER_ID` | Chat ID of the person to nag | ✅ |
| `SECONDARY_USER_ID` | Chat ID of the emergency contact | ✅ |

## 🔧 Customization

Want to modify the nagging schedule? Edit these values in `main.go`:

- `72 * time.Hour` - Change the 3-day interval
- `48 * time.Hour` - Change the response timeout
- Message texts - Make them more or less passive-aggressive

## 🐛 Troubleshooting

**Bot not responding?**
- Check if your bot token is correct
- Make sure the bot can send messages to both users
- Verify chat IDs are correct (and not your cat's ID)

**Running out of memory?**
- This bot uses ~5MB of RAM. If you're running out of memory, you might want to check what else is running on your server (probably Chrome with 847 tabs open)

**Bot being too annoying?**
- That's a feature, not a bug! 
- But seriously, adjust the time intervals if needed

## 🤝 Contributing

Found a bug? Want to make Are-you-dead-yet even more annoying? Feel free to:

1. Fork this repository
2. Create a feature branch (`git checkout -b feature/more-nagging`)
3. Commit your changes (`git commit -am 'Add even more persistent nagging'`)
4. Push to the branch (`git push origin feature/more-nagging`)
5. Create a Pull Request

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⚠️ Disclaimer

Are-you-dead-yet is not responsible for:
- Damaged relationships due to excessive nagging
- Increased anxiety from constant check-ins
- Your friends blocking the bot (and possibly you)
- Any existential crises caused by digital surveillance

Use responsibly and remember: with great nagging power comes great nagging responsibility.

## 🎉 Acknowledgments

- [@BotFather](https://t.me/botfather) for making bot creation easy
- The Go team for creating a language that compiles to tiny binaries
- Everyone who's ever been worried about someone and needed a digital assistant to do the nagging for them

---

*Made with ❤️ and a healthy dose of digital paranoia*
