package models

import "time"

// Ad 定義廣告模型
type Ad struct {
    ID         uint      `json:"id" gorm:"primary_key"`
    Title      string    `json:"title"`
    StartAt    time.Time `json:"startAt"`
    EndAt      time.Time `json:"endAt"`
    Conditions Condition `json:"conditions"`
}

// Condition 定義廣告條件
type Condition struct {
    AgeStart  int      `json:"ageStart"`
    AgeEnd    int      `json:"ageEnd"`
    Country   []string `json:"country"`
    Platform  []string `json:"platform"`
}