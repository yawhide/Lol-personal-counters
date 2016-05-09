package main

import (
    "errors"
    "fmt"
    "github.com/yawhide/go-lol"
    "github.com/spf13/viper"
    "gopkg.in/pg.v4"
    "strconv"
    "strings"
    "time"
)

// var PLATFORM_IDS = map[string]string{
//     "br": "BR1",
//     "eune": "EUN1",
//     "euw": "EUW1",
//     "kr": "KR",
//     "lan": "LA1",
//     "las": "LA2",
//     "na": "NA1",
//     "oce": "OC1",
//     "tr": "TR1",
//     "ru": "RU",
// }
var RIOT_REGIONS = map[string]bool{
    "br": true,
    "eune": true,
    "euw": true,
    "kr": true,
    "lan": true,
    "las": true,
    "na": true,
    "oce": true,
    "tr": true,
    "ru": true,
}
const LOL_API_HOST_SUFFIX = "api.pvp.net"
var CHAMPIONS = [...]string { "shaco","drmundo","rammus","anivia","irelia","yasuo","sona","kassadin","zac","gnar","karma","corki","gangplank","janna","jhin","kindred","braum","ashe","tryndamere","jax","morgana","zilean","singed","evelynn","twitch","galio","velkoz","olaf","annie","karthus","leblanc","urgot","amumu","xinzhao","chogath","twistedfate","fiddlesticks","vladimir","warwick","teemo","tristana","sivir","soraka","ryze","sion","masteryi","alistar","missfortune","nunu","rengar","volibear","fizz","graves","ahri","shyvana","lux","xerath","thresh","shen","kogmaw","jinx","tahmkench","riven","talon","malzahar","kayle","kalista","reksai","illaoi","leona","lulu","gragas","poppy","fiora","ziggs","udyr","viktor","sejuani","varus","nautilus","draven","bard","mordekaiser","ekko","yorick","pantheon","ezreal","garen","akali","kennen","vayne","jayce","lissandra","cassiopeia","rumble","khazix","darius","hecarim","skarner","lucian","heimerdinger","nasus","zed","nidalee","syndra","jarvaniv","quinn","renekton","maokai","aurelionsol","nocturne","katarina","leesin","monkeyking","azir","brand","diana","elise","nami","aatrox","orianna","zyra","trundle","veigar","taric","caitlyn","blitzcrank","malphite","vi","swain" }
var CHAMPION_KEYS = map[string]string{ "aatrox":"266","ahri":"103","akali":"84","alistar":"12","amumu":"32","anivia":"34","annie":"1","ashe":"22","aurelionsol":"136","azir":"268","bard":"432","blitzcrank":"53","brand":"63","braum":"201","caitlyn":"51","cassiopeia":"69","chogath":"31","corki":"42","darius":"122","diana":"131","draven":"119","drmundo":"36","ekko":"245","elise":"60","evelynn":"28","ezreal":"81","fiddlesticks":"9","fiora":"114","fizz":"105","galio":"3","gangplank":"41","garen":"86","gnar":"150","gragas":"79","graves":"104","hecarim":"120","heimerdinger":"74","illaoi":"420","irelia":"39","janna":"40","jarvaniv":"59","jax":"24","jayce":"126","jhin":"202","jinx":"222","kalista":"429","karma":"43","karthus":"30","kassadin":"38","katarina":"55","kayle":"10","kennen":"85","khazix":"121","kindred":"203","kogmaw":"96","leblanc":"7","leesin":"64","leona":"89","lissandra":"127","lucian":"236","lulu":"117","lux":"99","malphite":"54","malzahar":"90","maokai":"57","masteryi":"11","missfortune":"21","monkeyking":"62","wukong":"62","mordekaiser":"82","morgana":"25","nami":"267","nasus":"75","nautilus":"111","nidalee":"76","nocturne":"56","nunu":"20","olaf":"2","orianna":"61","pantheon":"80","poppy":"78","quinn":"133","rammus":"33","reksai":"421","renekton":"58","rengar":"107","riven":"92","rumble":"68","ryze":"13","sejuani":"113","shaco":"35","shen":"98","shyvana":"102","singed":"27","sion":"14","sivir":"15","skarner":"72","sona":"37","soraka":"16","swain":"50","syndra":"134","tahmkench":"223","talon":"91","taric":"44","teemo":"17","thresh":"412","tristana":"18","trundle":"48","tryndamere":"23","twistedfate":"4","twitch":"29","udyr":"77","urgot":"6","varus":"110","vayne":"67","veigar":"45","velkoz":"161","vi":"254","viktor":"112","vladimir":"8","volibear":"106","warwick":"19","xerath":"101","xinzhao":"5","yasuo":"157","yorick":"83","zac":"154","zed":"238","ziggs":"115","zilean":"26","zyra":"143"}
var CHAMPION_KEYS_BY_KEY = map[string]string{ "1":"annie", "2":"olaf", "3":"galio", "4":"twistedfate", "5":"xinzhao", "6":"urgot", "7":"leblanc", "8":"vladimir", "9":"fiddlesticks", "10":"kayle", "11":"masteryi", "12":"alistar", "13":"ryze", "14":"sion", "15":"sivir", "16":"soraka", "17":"teemo", "18":"tristana", "19":"warwick", "20":"nunu", "21":"missfortune", "22":"ashe", "23":"tryndamere", "24":"jax", "25":"morgana", "26":"zilean", "27":"singed", "28":"evelynn", "29":"twitch", "30":"karthus", "31":"chogath", "32":"amumu", "33":"rammus", "34":"anivia", "35":"shaco", "36":"drmundo", "37":"sona", "38":"kassadin", "39":"irelia", "40":"janna", "41":"gangplank", "42":"corki", "43":"karma", "44":"taric", "45":"veigar", "48":"trundle", "50":"swain", "51":"caitlyn", "53":"blitzcrank", "54":"malphite", "55":"katarina", "56":"nocturne", "57":"maokai", "58":"renekton", "59":"jarvaniv", "60":"elise", "61":"orianna", "62":"monkeyking", "63":"brand", "64":"leesin", "67":"vayne", "68":"rumble", "69":"cassiopeia", "72":"skarner", "74":"heimerdinger", "75":"nasus", "76":"nidalee", "77":"udyr", "78":"poppy", "79":"gragas", "80":"pantheon", "81":"ezreal", "82":"mordekaiser", "83":"yorick", "84":"akali", "85":"kennen", "86":"garen", "89":"leona", "90":"malzahar", "91":"talon", "92":"riven", "96":"kogmaw", "98":"shen", "99":"lux", "101":"xerath", "102":"shyvana", "103":"ahri", "104":"graves", "105":"fizz", "106":"volibear", "107":"rengar", "110":"varus", "111":"nautilus", "112":"viktor", "113":"sejuani", "114":"fiora", "115":"ziggs", "117":"lulu", "119":"draven", "120":"hecarim", "121":"khazix", "122":"darius", "126":"jayce", "127":"lissandra", "131":"diana", "133":"quinn", "134":"syndra", "136":"aurelionsol", "143":"zyra", "150":"gnar", "154":"zac", "157":"yasuo", "161":"velkoz", "201":"braum", "202":"jhin", "203":"kindred", "222":"jinx", "223":"tahmkench", "236":"lucian", "238":"zed", "245":"ekko", "254":"vi", "266":"aatrox", "267":"nami", "268":"azir", "412":"thresh", "420":"illaoi", "421":"reksai", "429":"kalista", "432":"bard"}
var CHAMPION_KEYS_BY_KEY_PROPER_CASING = map[string]string{ "35": "Shaco", "36": "DrMundo", "33": "Rammus", "34": "Anivia", "39": "Irelia", "157": "Yasuo", "37": "Sona", "38": "Kassadin", "154": "Zac", "150": "Gnar", "43": "Karma", "42": "Corki", "41": "Gangplank", "40": "Janna", "202": "Jhin", "203": "Kindred", "201": "Braum", "22": "Ashe", "23": "Tryndamere", "24": "Jax", "25": "Morgana", "26": "Zilean", "27": "Singed", "28": "Evelynn", "29": "Twitch", "3": "Galio", "161": "Velkoz", "2": "Olaf", "1": "Annie", "30": "Karthus", "7": "Leblanc", "6": "Urgot", "32": "Amumu", "5": "XinZhao", "31": "Chogath", "4": "TwistedFate", "9": "FiddleSticks", "8": "Vladimir", "19": "Warwick", "17": "Teemo", "18": "Tristana", "15": "Sivir", "16": "Soraka", "13": "Ryze", "14": "Sion", "11": "MasterYi", "12": "Alistar", "21": "MissFortune", "20": "Nunu", "107": "Rengar", "106": "Volibear", "105": "Fizz", "104": "Graves", "103": "Ahri", "102": "Shyvana", "99": "Lux", "101": "Xerath", "412": "Thresh", "98": "Shen", "96": "KogMaw", "222": "Jinx", "223": "TahmKench", "92": "Riven", "91": "Talon", "90": "Malzahar", "10": "Kayle", "429": "Kalista", "421": "RekSai", "420": "Illaoi", "89": "Leona", "117": "Lulu", "79": "Gragas", "78": "Poppy", "114": "Fiora", "115": "Ziggs", "77": "Udyr", "112": "Viktor", "113": "Sejuani", "110": "Varus", "111": "Nautilus", "119": "Draven", "432": "Bard", "82": "Mordekaiser", "245": "Ekko", "83": "Yorick", "80": "Pantheon", "81": "Ezreal", "86": "Garen", "84": "Akali", "85": "Kennen", "67": "Vayne", "126": "Jayce", "127": "Lissandra", "69": "Cassiopeia", "68": "Rumble", "121": "Khazix", "122": "Darius", "120": "Hecarim", "72": "Skarner", "236": "Lucian", "74": "Heimerdinger", "75": "Nasus", "238": "Zed", "76": "Nidalee", "134": "Syndra", "59": "JarvanIV", "133": "Quinn", "58": "Renekton", "57": "Maokai", "136": "AurelionSol", "56": "Nocturne", "55": "Katarina", "64": "LeeSin", "62": "MonkeyKing", "268": "Azir", "63": "Brand", "131": "Diana", "60": "Elise", "267": "Nami", "266": "Aatrox", "61": "Orianna", "143": "Zyra", "48": "Trundle", "45": "Veigar", "44": "Taric", "51": "Caitlyn", "53": "Blitzcrank", "54": "Malphite", "254": "Vi", "50": "Swain"}

var roleToSqlMapping = map [string]string{
    "Middle": "(lane = 'MID' OR lane = 'MIDDLE')",
    "Top": "lane = 'TOP'",
    "Jungle": "lane = 'JUNGLE'",
    "ADC": "((lane = 'BOTTOM' OR lane = 'BOT') AND role = 'DUO_CARRY')",
    "Support": "((lane = 'BOTTOM' OR lane = 'BOT') AND role = 'DUO_SUPPORT')",
}

var apiEndpointMap map[string]*lol.APIEndpoint

type Mastery struct {
    ChampionId                   int   `json:"championId"`
    ChampionLevel                int   `json:"championLevel"`
    ChampionPoints               int   `json:"championPoints"`
    ChampionPointsSinceLastLevel int64 `json:"championPointsSinceLastLevel"`
    ChampionPointsUntilNextLevel int64 `json:"championPointsUntilNextLevel"`
    LastPlayTime                 int64 `json:"lastPlayTime"`
    SummonerId                   int64 `json:"playerId"`
}

type RiotError struct {
    StatusCode int
}

func (err RiotError) Error() string {
    return fmt.Sprintf("Error: HTTP Status %d", err.StatusCode)
}

type MySummoner struct {
    SummonerId         uint64  `json:"id", sql:",pk"`
    Name               string  `json:"name"`
    ProfileIconID      int     `json:"profileIconId"`
    MasteriesUpdatedAt time.Time
    MatchlistUpdatedAt time.Time
    RevisionDate       uint64  `json:"revisionDate"`
    StartIndex         int
    SummonerLevel      uint32  `json:"summonerLevel"`
}

func (s *MySummoner) SetSummoner(name string) {
    s.Name = name
}

func (s *MySummoner) UpdateStartIndex(i int) {
    s.StartIndex = i
    s.MatchlistUpdatedAt = time.Now().UTC()
}

type ChampionMatchupWinrate struct {
    Role string `json:"role"`
    Matchups []Matchup
}

type Matchup struct {
    Games int `json:"games"`
    StatScore float32 `json:"statScore"`
    WinRate float32 `json:"winRate"`
    Enemy string `json:"key"`
}

type ChampionMatchup struct {
    Champion string `sql:",pk`
    Enemy string `sql:",pk`
    Games int
    Role string `sql:",pk`
    StatScore float32
    WinRate float32
    PersonalWinRate string
    PersonalGames int
}

func (c *ChampionMatchup) SetChampion(name string) {
    c.Champion = name
}

func (m ChampionMatchup) String() string {
    return fmt.Sprintf("[%s] %s vs %s. %g %d games\n", m.Role, CHAMPION_KEYS_BY_KEY[m.Champion], CHAMPION_KEYS_BY_KEY[m.Enemy], m.WinRate, m.Games)
}

func (c *ChampionMatchup) UpdatePersonalData(p PersonalMatchup) {
    c.PersonalWinRate = strings.TrimSuffix(fmt.Sprintf("%0.2f", p.WinRate), ".00")
    c.PersonalGames = p.Games
}

type PersonalMatch struct {
    Champion   string
    Enemy      string
    Lane       string
    MatchId    uint64
    Role       string
    PlatformId string
    QueueType  string
    Season     string
    SummonerId uint64
    Timestamp  uint64
    Winner     bool
}

type PersonalMatchup struct {
    Champion string
    Winner   bool
    WinRate  float32
    Games    int
}

func setupRiotApi() error {
    // keyStorer, err := lol.NewXdgAPIKeyStorer()
    // if err != nil {
    //     panic(err)
    // }
    // riotKey, ok := keyStorer.Get()
    // if ok == false {
    //     panic("No API key found")
    // }
    key := lol.APIKey(viper.GetString("riot.key"))
    apiEndpointMap = make(map[string]*lol.APIEndpoint)
    for r, _ := range RIOT_REGIONS {
        region, _ := lol.NewRegionByCode(r)
        api, err := lol.NewAPIEndpoint(region, key)
        if err != nil {
            return err
        }
        apiEndpointMap[r] = api
    }
    return nil
}

func createSchema(db *pg.DB) error {
    queries := []string{
        // `DROP TABLE champion_matchups`,
        // `DROP TABLE masteries`,
        // `DROP TABLE summoners`,

        `CREATE TABLE IF NOT EXISTS champion_matchups (
            champion text,
            enemy text,
            games int,
            role text,
            stat_score decimal,
            win_rate decimal,
            PRIMARY KEY(champion, enemy, role))`,

        `CREATE TABLE IF NOT EXISTS masteries (
            champion_id bigint,
            champion_level int,
            champion_points int,
            champion_points_since_last_level bigint,
            champion_points_until_next_level bigint,
            last_play_time bigint,
            summoner_id bigint,
            PRIMARY KEY (champion_id, summoner_id))`,

        `CREATE TABLE IF NOT EXISTS personal_matches (
            champion    text,
            enemy       text,
            lane        text,
            match_id    bigint,
            role        text,
            platform_id text,
            queue_type  text,
            season      text,
            summoner_id bigint,
            timestamp   bigint,
            winner      boolean)`,

        `CREATE TABLE IF NOT EXISTS my_summoners (
            name text,
            profile_icon_id int,
            masteries_updated_at timestamp with time zone DEFAULT (now() at time zone 'utc'),
            matchlist_updated_at timestamp with time zone DEFAULT (now() at time zone 'utc'),
            revision_date bigint,
            summoner_level int,
            start_index int DEFAULT 0,
            summoner_id bigint PRIMARY KEY)`,

        `CREATE OR REPLACE FUNCTION upsert_masteries(c_id bigint, c_lv int, c_pts int, c_pts_since_last_lv bigint, c_pts_until_next_lv bigint, last_play_t bigint, s_id bigint) RETURNS VOID AS $$
            DECLARE
            BEGIN
                UPDATE masteries SET champion_id = c_id, champion_level = c_lv, champion_points = c_pts, champion_points_since_last_level = c_pts_since_last_lv, champion_points_until_next_level = c_pts_until_next_lv, last_play_time = last_play_t, summoner_id = s_id WHERE champion_id = c_id AND summoner_id = s_id;
                IF NOT FOUND THEN
                INSERT INTO masteries VALUES (c_id, c_lv, c_pts, c_pts_since_last_lv, c_pts_until_next_lv, last_play_t, s_id);
                END IF;
            END;
            $$ LANGUAGE 'plpgsql'`,

    }
    // fmt.Println("hey", queries, createTableSql)
    queries = append(queries, createTableSql...)

    for _, q := range queries {
        _, err := db.Exec(q)
        if err != nil {
            fmt.Println(q)
            return err
        }
    }
    return nil
}

func getSummonerMasteriesAndSave(region string, summonerName string, db *pg.DB) (err error) {
    names := NormalizeSummonerName(summonerName)
    summoners, err := getSummonerIdByNameAndSave(region, names, db)
    if err != nil {
        return
    }
    _, err = getChampionMasteriesBySummonerIdAndSave(region, summoners[0].SummonerId, db)
    return
}

func getSummonerIdByNameAndSave(region string, names []string, db *pg.DB) (players []MySummoner, err error) {
    summoners, err := apiEndpointMap[region].GetSummonerByName(names)
    fmt.Println(summoners)
    if err != nil {
        fmt.Println("Failed to get summoners:", names, err)
        return
    }
    for _, summoner := range summoners {
        if summoner.ID == 0 {
            fmt.Println("Summoner not found in returning json from riot, summoner name:", names, err)
            err = errors.New("Failed to get all summoners")
            return
        }
        s := MySummoner{
            uint64(summoner.ID),
            NormalizeSummonerName(summoner.Name)[0],
            summoner.ProfileIconID,
            time.Now().UTC(),
            time.Now().UTC(),
            uint64(summoner.RevisionDate),
            0,
            summoner.Level}
        err = db.Create(&s)
        if err != nil {
            fmt.Println("Failed to save summoner:", summoner.Name, "in db", err)
            return
        }
        players = append(players, s)
    }
    return
}

// func getAllChampions(region string) (champions, err error) {
//     /api/lol/static-data/{region}/v1.2/champion
// }

func getChampionMasteriesBySummonerIdAndSave(region string, summonerId uint64, db *pg.DB) (masteries []lol.ChampionMastery, err error) {
    // args := "api_key=" + viper.GetString("riot.key")
    // platformId := PLATFORM_IDS[region]
    // // var masteries []Mastery
    // url := fmt.Sprintf(
    //         "https://%v.%v/championmastery/location/%v/player/%v/champions?%v",
    //         region,
    //         LOL_API_HOST_SUFFIX,
    //         platformId,
    //         summonerId,
    //         args)
    // fmt.Println(url)
    // err = requestAndUnmarshal(url, &masteries)
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    masteries, err = apiEndpointMap[region].GetChampionMasteries(lol.SummonerID(summonerId))
    if err != nil {
        fmt.Println("Failed to get champion masteries for summoner:", summonerId, err)
        return
    }

    for _, mastery := range masteries {
        sql := fmt.Sprintf(
            "SELECT upsert_masteries(%v, %v, %v, %v, %v, %v, %v)",
            mastery.Champion,
            mastery.Level,
            mastery.Points,
            mastery.PointsSinceLastLevel,
            mastery.PointsUntilNextLevel,
            mastery.LastPlayTime,
            mastery.Player)
        _, err = db.Exec(sql)
        // err = db.Create(&mastery)
        if err != nil {
            // fmt.Println(sql)
            fmt.Println("Failed to upsert masteries for summoner:", summonerId, err)
            return
        }
    }
    // var s Summoner
    // err = db.Model(&s).Where("summoner_id = ?", summonerId).Select()
    // s.MasteriesUpdatedAt = time.Now().UTC()
    // fmt.Println(s)
    // db.Model(&s).Set("masteries_updated_at = ?masteries_updated_at").Where("summoner_id = ?summoner_id").Update()

    // err = db.Update(&s)
    // if err != nil {
    //     fmt.Println("failed to upsert new masteries for summoner:", summonerId, err)
    // }

    // _, err = db.Model(&s).Set("masteries_updated_at = ?", time.Now().UTC()).Returning("*").Update()
    // fmt.Println("wtf", db.Model(&s))

    // sql = fmt.Sprintf(
    //     "UPDATE summoners SET masteries_updated_at = '%v' WHERE summoner_id = '%v'",
    //     summonerId,
    //     time.Now().UTC())
    // fmt.Println(sql)
    // _, err = db.Exec(sql)
    // err = db.Update(&Summoner{
    //     SummonerId: summonerId,
    //     MasteriesUpdatedAt: time.Now().UTC(),
    // })
    var s MySummoner
    _, err = db.Model(&s).Set("masteries_updated_at = ?", time.Now().UTC()).Where("summoner_id = ?", summonerId).Update()

    if err != nil {
        fmt.Println("Failed to update summoner:", summonerId, "masteries updated at time.", err)
        return
    }

    // fmt.Println(masteries)
    return
}

/* ======================== db methods ======================= */

func getOrCreateSummoner(region string, summonerName string, db *pg.DB) (summoner MySummoner, err error) {
    name := NormalizeSummonerName(summonerName)[0]
    err = db.Model(&summoner).Where("name = ?", name).Select()
    if err != nil {
        if err.Error() == "pg: no rows in result set" {
            err = getSummonerMasteriesAndSave(region, name, db)
            if err != nil {
                fmt.Println(err)
                return
            }
            err = db.Model(&summoner).Where("name = ?", name).Select()
            if err != nil {
                fmt.Println(err)
                return
            }
            getMatchlistBySummonerIdAndSave(region, summoner, db)
        } else {
            fmt.Println(err)
            return
        }
    }
    if summoner.MasteriesUpdatedAt.UTC().Add(time.Duration(60*60*24)*time.Second).Before(time.Now().UTC()) {
        fmt.Printf("Summoner %s masteries are out of date...lets update them!\n", summoner.Name)
        getChampionMasteriesBySummonerIdAndSave(region, summoner.SummonerId, db)
        getMatchlistBySummonerIdAndSave(region, summoner, db)
    }

    fmt.Println("getSummonerById", name, summoner)
    return
}

func getMatchups(summoner_id uint64, enemy_champion_id string, role string, db *pg.DB) (matchups []ChampionMatchup, finalPm []PersonalMatchup, err error) {
    //select * from champion_matchups where enemy = '57' and champion IN (select cast(champion_id as text)  from masteries where summoner_id = 26691960) order by win_rate desc;

    err = db.Model(&matchups).Where("role = ? AND enemy = ? AND champion IN (SELECT CAST(champion_id AS text) FROM masteries WHERE summoner_id = ?)", role, enemy_champion_id, summoner_id).Order("win_rate desc").Select()
    if err != nil {
        fmt.Println("Failed to get matchups from database", err)
        return
    }

    sql := fmt.Sprintf(
            `select champion, winner, (100.0*games_played) / games as win_rate, games
            from
              (select champion, winner, count(enemy) as games_played from personal_matches where enemy = '%s' and %s group by champion, winner) t1
            natural join
              (select count(champion) as games, champion from personal_matches where enemy = '%s' and %s group by champion) t2 order by win_rate desc;`,
            enemy_champion_id,
            roleToSqlMapping[role],
            enemy_champion_id,
            roleToSqlMapping[role])
    // fmt.Println(sql)

    var pm []PersonalMatchup
    _, err = db.Query(&pm, sql)
    // return users, err
    // res, err := db.Exec(sql, pm)
    if err != nil {
        fmt.Println("Failed to get personal winrates for summoner:", summoner_id)
        return
    }
    for _, p := range pm {
        found := false
        if !p.Winner {
            for _, p2 := range pm {
                if (p.Champion == p2.Champion && p.Winner != p2.Winner) {
                    // fmt.Println(p.Champion, p2.Champion, p.Winner, p2.Winner)
                    found = true
                    break
                }
            }
        }
        // fmt.Println(found, p)
        if !found {
            finalPm = append(finalPm, p)
        }
    }
    // fmt.Println("personal matchup:", pm, finalPm)

    // fmt.Println("matchups", matchups)
    return
}

func getMatchlistBySummonerIdAndSave(region string, summoner MySummoner, db *pg.DB) (err error) {
    matchList, err := apiEndpointMap[region].GetMatchlist(summoner.SummonerId, summoner.StartIndex)
    fmt.Println(matchList)
    if err != nil {
        fmt.Println("Failed to get matchlist for summoner:", summoner.SummonerId, err)
        return
    }
    // err = db.Model(&summoner).Update()
    fmt.Println("before:", summoner)
    summoner.UpdateStartIndex(matchList.EndIndex)
    fmt.Println("after:", summoner)
    _, err = db.Model(&summoner).Set("matchlist_updated_at = ?, start_index = ?", summoner.MatchlistUpdatedAt, summoner.StartIndex).Where("summoner_id = ?summoner_id").Update()
    if err != nil {
        fmt.Println("failed to update summoners matchlist start index:", summoner.SummonerId, err)
        return
    }
    // var pm []PersonalMatch
    q := `INSERT INTO personal_matches VALUES`
    for _, m := range matchList.Matches {
        q += fmt.Sprintf(
            " ('%v', '', '%v', %v, '%v', '%v', '%v', '%v', %v, %v, false),",
        // pm = append(pm, PersonalMatch{
            m.Champion,
            // "",
            m.Lane,
            m.MatchId,
            m.Role,
            m.PlatformId,
            m.Queue,
            m.Season,
            summoner.SummonerId,
            m.Timestamp,
            // false})
            )
    }
    q = strings.TrimSuffix(q, ",")
    // err = db.Create(&pm)
    fmt.Println("q:", q)
    _, err = db.Exec(q)
    if err != nil {
        fmt.Println("failed to bulk insert match list for summoner:", summoner.SummonerId, err)
        return
    }
    for _, m := range matchList.Matches {
        game, err := apiEndpointMap[region].GetMatch(m.MatchId, false)
        if err != nil {
            fmt.Println("Failed to get ranked game for summoner:", summoner.SummonerId, err)
        } else {
            var pm PersonalMatch
            enemy, winner := determineEnemyAndWinner(summoner.SummonerId, game)
            if enemy == "" {
                fmt.Println("failed to determine enemy or winner for summoner:", summoner.SummonerId, "and match:", m.MatchId)
                continue
            }
            fmt.Println("enemy and winner:", enemy, winner)
            //TODO do a get, then an update to save writes if need be
            _, err = db.Model(&pm).Set("enemy = ?, winner = ?", enemy, winner).Where("summoner_id = ? and match_id = ?", summoner.SummonerId, m.MatchId).Update()
            if err != nil {
                fmt.Println("failed to update personal match with summoner:", summoner.SummonerId, "and match:", m.MatchId, err)
            }
        }
    }
    return nil
}

/* ======================== helper ========================== */

func createSummonerIDString(summonerID []int64) (summonerIDstr string, err error) {
    if len(summonerID) > 40 {
        return summonerIDstr, errors.New("A Maximum of 40 SummonerIDs are allowed")
    }
    for k, v := range summonerID {
        summonerIDstr += strconv.FormatInt(v, 10)
        if k != len(summonerID)-1 {
            summonerIDstr += ","
        }
    }
    return
}

//NormalizeSummonerName takes an arbitrary number of strings and returns a string array containing the strings
//standardized to league of legends internal standard (lowecase and strings removed)
func NormalizeSummonerName(summonerNames ...string) []string {
    for i, v := range summonerNames {
        summonerName := strings.ToLower(v)
        summonerName = strings.Replace(summonerName, " ", "", -1)
        summonerNames[i] = summonerName
    }
    return summonerNames
}

func NormalizeChampion(name string) string {
    name = strings.ToLower(name)
    name = strings.Replace(name, " ", "", -1)
    name = strings.Replace(name, "'", "", -1)
    name = strings.Replace(name, ".", "", -1)
    return name
}

func determineEnemyAndWinner(summonerId uint64, game *lol.Match) (enemy string, winner bool) {
    // type Match struct {
    //     Participants []struct {
    //         Stat struct {
    //             Winner bool `json:"winner"`
    //         }   `json:"stats"`
    //         Timeline struct {
    //             Role string `json:"role"`
    //             Lane string `json:"lane"`
    //         }   `json:"timeline"`
    //         ParticipantId int `json:"participantId"`
    //         ChampionId    int `json:"championId"`
    //         TeamId        int `json:"teamId"`
    //     }   `json:"participants"`
    //     ParticipantIdentities []struct {
    //         Player struct {
    //             SummonerId uint64 `json:"summonerId"`
    //         }
    //         ParticipantId int `json:"participantId"`
    //     }   `json:"participantIdentities"`
    // }
    var participantId int
    for _, p := range game.ParticipantIdentities {
        if p.Player.SummonerId == summonerId {
            participantId = p.ParticipantId
            break
        }
    }
    var teamId int
    var role string
    var lane string
    for _, p := range game.Participants {
        if p.ParticipantId == participantId {
            teamId = p.TeamId
            role = p.Timeline.Role
            lane = p.Timeline.Lane
            winner = p.Stat.Winner
        }
    }
    for _, p := range game.Participants {
        if role == p.Timeline.Role && lane == p.Timeline.Lane && p.TeamId != teamId {
            enemy = fmt.Sprintf("%v", p.ChampionId)
        }
    }
    return
}
