package main

import (

)

type AnalyticsIndex struct {
    SummonerId int64
    Enemy      string
    Role       string
    RememberMe bool
    Referral   string // url

}

type AnalyticsMatchup struct {
    SummonerId int64
    Enemy      string
    Role       string
    Click      string // url

}

type AnalyticsExternalLink struct {
    Url  string
    Page string // which page user was on
}

var createTableSql = []string{
    `CREATE TABLE IF NOT EXISTS analytics_index (
        summoner_id bigint,
        enemy text,
        role text,
        remember_me boolean,
        referral text)`,

    `CREATE TABLE IF NOT EXISTS analytics_matchup (
        summoner_id bigint,
        enemy text,
        role text,
        click text)`,

    `CREATE TABLE IF NOT EXISTS analytics_external_link (
        url text,
        page text)`,
}


