// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    gameSetting, err := UnmarshalGameSetting(bytes)
//    bytes, err = gameSetting.Marshal()

package est

import "encoding/json"

func UnmarshalGameSetting(data []byte) (GameSetting, error) {
	var r GameSetting
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GameSetting) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type GameSetting struct {
	Info     *Info     `json:"info,omitempty"`
	Settings *Settings `json:"settings,omitempty"`
}

type Info struct {
	GameName          string   `json:"GameName,omitempty"`
	Version           string   `json:"Version,omitempty"`
	Platform          []string `json:"Platform,omitempty"`
	GoVersion         string   `json:"GoVersion,omitempty"`
	EbitengineVersion string   `json:"EbitengineVersion,omitempty"`
}

type Settings struct {
	ScreenSize              [2]int `json:"ScreenSize,omitempty"`
	FullScreen              bool   `json:"FullScreen,omitempty"`
	Tps                     int    `json:"TPS,omitempty"`
	FPS                     int    `json:"FPS,omitempty"`
	RunnableOnUnfocused     bool   `json:"RunnableOnUnfocused,omitempty"`
	ScreenClearedEveryFrame bool   `json:"ScreenClearedEveryFrame,omitempty"`
	VsyncEnabled            bool   `json:"VsyncEnabled,omitempty"`
	WindowClosingHandled    bool   `json:"WindowClosingHandled,omitempty"`
	WindowDecorated         bool   `json:"WindowDecorated,omitempty"`
	WindowFloating          bool   `json:"WindowFloating,omitempty"`
	WindowIcon              string `json:"WindowIcon,omitempty"`
	WindowPosition          [2]int `json:"WindowPosition,omitempty"`
	WindowResizingMode      int    `json:"WindowResizingMode,omitempty"`
	WindowSize              [2]int `json:"WindowSize,omitempty"`
	WindowSizeLimits        [2]int `json:"WindowSizeLimits,omitempty"`
	WindowTitle             string `json:"WindowTitle,omitempty"`
}

/*json
{

{

    "info": {
        "GameName": "value",
        "Version": "value",
        "Platform": "value",
        "GoVersion": "value",
        "EbitengineVersion": "value"
    },
    "settings": {
     "ScreenSize":[
         1,1],
    "FullScreen": true ,
    "TPS": 1,
    "FPS": 1,
    "RunnableOnUnfocused": true,
    "ScreenClearedEveryFrame": true,
    "VsyncEnabled": true,
    "WindowClosingHandled": true,
    "WindowDecorated": true,
    "WindowFloating": true,
    "WindowIcon": "value",
    "WindowPosition": [
        1,
        1
    ],
    "WindowResizingMode": 1,
    "WindowSize": [
        1,
        1
    ],
    "WindowSizeLimits": [
        1,
        1
    ],
    "WindowTitle": "value"
    }
}
*/
