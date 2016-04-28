package main

import (
    "gopkg.in/pg.v4"
)

func main() {
    db := pg.Connect(&pg.Options{
        User: "yawhide",
    })

    _, err := db.Exec(`DROP TABLE champion_matchups`)
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
