# Discord College Football Bot

A Discord bot written in Go that provides up-to-date college football scores and schedules. It responds to commands to show recent or upcoming games for specific teams using the [College Football Data API](https://collegefootballdata.com/).

## Features

- Retrieves college football scores by team, including live scores (for Patreon members)
- Provides upcoming game schedules if a game hasn’t started
- Designed to run on a Discord server, responds to commands in real-time

## Commands

- **!s `<team_name>`** — Retrieves the score for the most recent game or the upcoming game schedule if the game hasn’t started.
  
Example:
```
!s Alabama
```

## Setup

### Prerequisites

- **Go** (1.19 or later)
- **College Football Data API key**: Sign up at [College Football Data](https://collegefootballdata.com/) and obtain an API key.
- **Discord Bot Token**: Set up a bot on the [Discord Developer Portal](https://discord.com/developers/applications) to obtain a token.

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/discord-cfb-bot.git
   cd discord-cfb-bot
   ```

2. **Configure API keys**:
   - In `config/config.go`, set the `CFBDAPIKey` with your College Football Data API key and your Discord bot token.

3. **Run the bot**:
   - To test locally:
     ```bash
     go run main.go
     ```
   - To run in the background:
     - Use `nohup`, `screen`, `tmux`, or create a `systemd` service (see below for options).

### Running the Bot in the Background

You can keep the bot running in the background by using tools like `nohup`, `screen`, or `tmux`. For production servers, consider creating a `systemd` service.

Example with `nohup`:
```bash
nohup go run main.go > bot.log 2>&1 &
```

### Deployment with `systemd` (Linux)

1. Create a `systemd` service file at `/etc/systemd/system/discord-cfb-bot.service`:
   ```ini
   [Unit]
   Description=Discord College Football Bot
   After=network.target

   [Service]
   ExecStart=/usr/local/go/bin/go run /path/to/main.go
   WorkingDirectory=/path/to/discord-cfb-bot
   StandardOutput=append:/path/to/bot.log
   StandardError=append:/path/to/bot.log
   Restart=always
   User=yourusername

   [Install]
   WantedBy=multi-user.target
   ```

2. **Enable and start the service**:
   ```bash
   sudo systemctl enable discord-cfb-bot
   sudo systemctl start discord-cfb-bot
   ```

## Usage

Invite the bot to your server and type commands in any text channel where the bot has permissions. For example:
```
!s Texas A&M
```

The bot will reply with the score or the upcoming game schedule for the team.

## Project Structure

- `main.go`: Entry point of the bot
- `config/config.go`: Configuration file for the API keys and other settings
- `clients/cfbd_client.go`: Client for interacting with the College Football Data API
- `bot/bot.go`: Discord bot setup and command handling

## Contributions

Contributions are welcome! Feel free to submit a pull request to improve the bot or add new features.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

