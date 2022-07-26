// Created by EldersJavas(EldersJavas&gmail.com)
// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    gameConfig, err := UnmarshalGameConfig(bytes)
//    bytes, err = gameConfig.Marshal()

package config

import "encoding/json"

func UnmarshalGameConfig(data []byte) (GameConfig, error) {
	var r GameConfig
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GameConfig) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GameConfig struct {
	Version    *string      `json:"Version,omitempty"`
	GameName   *string      `json:"GameName,omitempty"`
	GameIcon   *string      `json:"GameIcon,omitempty"`
	RecordList []RecordList `json:"RecordList,omitempty"`
	TimingOpt  *string      `json:"TimingOpt,omitempty"`
	TimeFormat *string      `json:"TimeFormat,omitempty"`
	Site       *string      `json:"Site,omitempty"`
}

type RecordList struct {
	Player *string `json:"Player,omitempty"`
	Time   *string `json:"Time,omitempty"`
	Note   *string `json:"note,omitempty"`
}
