package main

import (
    "encoding/json"
    "errors"
    "fmt"
    // "gopkg.in/pg.v4"
    "net/http"
)

type AnalyticsPage struct {
    Page     string
    Referrer string
}

type AnalyticsIndex struct {
    SummonerName string
    Enemy        string
    Role         string
    RememberMe   bool
    Referral     string // url
}

func (a *AnalyticsIndex) SetSummonerName(name string) {
    a.SummonerName = NormalizeSummonerName(name)[0]
}

func (a *AnalyticsIndex) SetEnemy(name string) {
    a.Enemy = NormalizeChampion(name)
}

type AnalyticsMatchup struct {
    SummonerName string
    Enemy        string
    Role         string
    Click        string // url
}

func (a *AnalyticsMatchup) SetSummonerName(name string) {
    a.SummonerName = NormalizeSummonerName(name)[0]
}

func (a *AnalyticsMatchup) SetEnemy(name string) {
    a.Enemy = NormalizeChampion(name)
}

type AnalyticsExternalLink struct {
    Url  string
    Page string // which page user was on
}

var createTableSql = []string{
    `CREATE TABLE IF NOT EXISTS analytics_pages (
        page text,
        referrer text)`,

    `CREATE TABLE IF NOT EXISTS analytics_indices (
        summoner_name text,
        enemy text,
        role text,
        remember_me boolean,
        referral text)`,

    `CREATE TABLE IF NOT EXISTS analytics_matchups (
        summoner_name text,
        enemy text,
        role text,
        click text)`,

    `CREATE TABLE IF NOT EXISTS analytics_external_links (
        url text,
        page text)`,
}

func AnalyzeIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        decoder := json.NewDecoder(r.Body)
        var a AnalyticsIndex
        err := decoder.Decode(&a)
        if err != nil {
            fmt.Println("Failed to decode index analyics data", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        a.SetSummonerName(a.SummonerName)
        a.SetEnemy(a.Enemy)
        fmt.Println(a)
        err = db.Create(&a)
        if err != nil {
            fmt.Println("Failed to save index analyics", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
    } else {
        http.Error(w, errors.New("not found").Error(), http.StatusNotFound)
    }
}

func AnalyzeMatchup(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        decoder := json.NewDecoder(r.Body)
        var a AnalyticsMatchup
        err := decoder.Decode(&a)
        if err != nil {
            fmt.Println("Failed to decode matchup analyics", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        a.SetSummonerName(a.SummonerName)
        a.SetEnemy(a.Enemy)
        fmt.Println(a)
        err = db.Create(&a)
        if err != nil {
            fmt.Println("Failed to save matchup analyics", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
    } else {
        http.Error(w, errors.New("not found").Error(), http.StatusNotFound)
    }
}

func AnalyzeExternalLink(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        decoder := json.NewDecoder(r.Body)
        var a AnalyticsExternalLink
        err := decoder.Decode(&a)
        if err != nil {
            fmt.Println("Failed to decode external link analyics", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        fmt.Println(a)
        err = db.Create(&a)
        if err != nil {
            fmt.Println("Failed to save external link analyics", err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusOK)
    } else {
        http.Error(w, errors.New("not found").Error(), http.StatusNotFound)
    }
}
