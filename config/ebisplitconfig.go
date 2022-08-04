// Created by EldersJavas(EldersJavas&gmail.com)

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    config, err := UnmarshalConfig(bytes)
//    bytes, err = config.Marshal()

package config

import "encoding/json"

func UnmarshalConfig(data []byte) (Config, error) {
	var r Config
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Config) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Config struct {
	Version    string      `json:"Version,omitempty"`
	WindowSize *WindowSize `json:"WindowSize,omitempty"`
	Background *Background `json:"Background,omitempty"`
	Foreground *Background `json:"Foreground,omitempty"`
	Font       *Font       `json:"Font,omitempty"`
}

type Background struct {
	Color *Color `json:"Color,omitempty"`
}

type Color struct {
	R int `json:"R,omitempty"`
	G int `json:"G,omitempty"`
	B int `json:"B,omitempty"`
	A int `json:"A,omitempty"`
}

type Font struct {
	Normal *Background `json:"Normal,omitempty"`
	Time   *Background `json:"Time,omitempty"`
}

type WindowSize struct {
	Width  int `json:"Width,omitempty"`
	Height int `json:"Height,omitempty"`
}
