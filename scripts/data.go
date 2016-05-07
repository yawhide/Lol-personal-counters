package main

import (
    "encoding/json"
    "fmt"
    "github.com/spf13/viper"
    "gopkg.in/pg.v4"
    "io/ioutil"
    "net/http"
)

var db *pg.DB

func main() {

    viper.SetConfigName("config")
    viper.AddConfigPath("../")
    viper.SetConfigType("json")
    err := viper.ReadInConfig()
    if err != nil { // Handle errors reading the config file
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
    username := viper.GetString("postgres.username")
    password := viper.GetString("postgres.password")

    db = pg.Connect(&pg.Options{
        User: username,
        Password: password,
    })

    _, err = db.Exec(`DROP TABLE IF EXISTS champion_matchups`)
    if err != nil {
        panic(err)
    }

    err = createSchema(db)
    if err != nil {
        panic(err)
    }


    // get all win rates between champions
    for _, champion := range CHAMPIONS {
        getWinrateForChampion(champion, db)
    }

}

func getWinrateForChampion(champion string, db *pg.DB) (matchups []ChampionMatchup, err error) {
    var matchupInfo []ChampionMatchupWinrate
    url := fmt.Sprintf(
        "http://api.champion.gg/champion/%v/matchup?api_key=%v",
        champion,
        viper.GetString("championgg.key"))
    fmt.Println(url)
    err = requestAndUnmarshal(url, &matchupInfo)
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, data := range matchupInfo {
        role := data.Role
        for _, matchup := range data.Matchups {
            c := ChampionMatchup{
                Champion: CHAMPION_KEYS[champion],
                Enemy: CHAMPION_KEYS[NormalizeChampion(matchup.Enemy)],
                Games: matchup.Games,
                Role: role,
                StatScore: matchup.StatScore,
                WinRate: matchup.WinRate}
            err = db.Create(&c)
            if err != nil {
                fmt.Println("Failed to create champion matchup", err)
                return
            }
            matchups = append(matchups, c)
        }
    }
    // fmt.Println(matchups)
    return
}


func requestAndUnmarshal(requestURL string, v interface{}) (err error) {
    resp, err := http.Get(requestURL)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return RiotError{StatusCode: resp.StatusCode}
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    err = json.Unmarshal(body, v)
    if err != nil {
        fmt.Println(err)
        return
    }
    return
}
