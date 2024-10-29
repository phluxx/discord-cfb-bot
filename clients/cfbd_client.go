package clients

import (
    "discord-cfb-bot/config"
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "time"
)

const seasonStartDate = "2024-08-28"

type Game struct {
    HomeTeam      string `json:"home_team"`
    AwayTeam      string `json:"away_team"`
    HomePoints    *int   `json:"home_points"`
    AwayPoints    *int   `json:"away_points"`
    StartDateTime string `json:"start_date"`
}

func getCurrentWeek() int {
    startDate, _ := time.Parse("2006-01-02", seasonStartDate)
    weeks := int(time.Since(startDate).Hours() / (24 * 7))
    return weeks + 1 // +1 to adjust to the 1-based week number
}

func GetGameInfo(teamName string) string {
    currentWeek := getCurrentWeek()
    
    // Use url.QueryEscape to properly encode team names with special characters
    encodedTeamName := url.QueryEscape(teamName)
    url := fmt.Sprintf("https://api.collegefootballdata.com/games?year=%d&week=%d&team=%s", time.Now().Year(), currentWeek, encodedTeamName)

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Bearer "+config.CFBDAPIKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error fetching data:", err)
        return "Error fetching game info."
    }
    defer resp.Body.Close()

    var games []Game
    if err := json.NewDecoder(resp.Body).Decode(&games); err != nil {
        fmt.Println("Error decoding response:", err)
        return "Error parsing game data."
    }

    for _, game := range games {
        startTime, _ := time.Parse(time.RFC3339, game.StartDateTime)
        if time.Now().After(startTime) && game.HomePoints != nil && game.AwayPoints != nil {
            return fmt.Sprintf("%s: %d %s: %d", game.AwayTeam, *game.AwayPoints, game.HomeTeam, *game.HomePoints)
        } else if time.Now().Before(startTime) {
            return fmt.Sprintf("%s @ %s %s Eastern", game.AwayTeam, game.HomeTeam, startTime.Format("03:04 PM"))
        }
    }

    return "No recent or upcoming game data found."
}
