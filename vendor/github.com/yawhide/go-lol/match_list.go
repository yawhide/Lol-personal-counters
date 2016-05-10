package lol

import (
	"fmt"
)

type Matchlist struct {
    EndIndex    int `json:"endIndex"`
    Matches     []struct {
        Champion    uint64    `json:"champion"`
        Lane    string      `json:"lane"` //Legal values: MID, MIDDLE, TOP, JUNGLE, BOT, BOTTOM
        MatchId uint64    `json:"matchId"`
        PlatformId  string    `json:"platformId"`
        Queue   string      `json:"queue"`//Legal values: TEAM_BUILDER_DRAFT_RANKED_5x5, RANKED_SOLO_5x5, RANKED_TEAM_3x3, RANKED_TEAM_5x5
        Region  string    `json:"region"`
        Role    string      `json:"role`//Legal values: DUO, NONE, SOLO, DUO_CARRY, DUO_SUPPORT
        Season  string      `json:"season"`//Legal values: PRESEASON3, SEASON3, PRESEASON2014, SEASON2014, PRESEASON2015, SEASON2015, PRESEASON2016, SEASON2016
        Timestamp   uint64    `json:"timestamp"`
    }   `json:"matches"`
    StartIndex  int `json:"startIndex"`
    TotalGames  int `json:"totalGames"`
}

func (a *APIEndpoint) formatMatchlistURL(summonerId uint64, options map[string]string) string {
    res := fmt.Sprintf("https://%s/api/lol/%s/v2.2/matchlist/by-summoner/%v?api_key=%s", a.region.url, a.region.code, summonerId, a.key)
    for k, v := range options {
        res = fmt.Sprintf("%s&%s=%s", res, k, v)
    }
    return res
}

func (a *APIEndpoint) GetMatchlist(summonerId uint64, startIndex int) (*Matchlist, error) {
    res := &Matchlist{}
    options := map[string]string{}
    options["beginIndex"] = fmt.Sprintf("%d", startIndex)
    options["endIndex"] = "2147483647"
    options["seasons"] = "SEASON2016,PRESEASON2016"

    url := a.formatMatchlistURL(summonerId, options)
    err := a.g.Get(url, &res)
    if err != nil {
        return nil, err
    }
    return res, nil
}
