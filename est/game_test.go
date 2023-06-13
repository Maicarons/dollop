package est

import "testing"

func TestGame_SceneSort(t *testing.T) {
	type fields struct {
		Scene       []*Scene
		RootObject  *Object
		GameSetting *GameSetting
	}
	tests := []struct {
		name   string
		fields fields
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Game{
				Scene:       tt.fields.Scene,
				RootObject:  tt.fields.RootObject,
				GameSetting: tt.fields.GameSetting,
			}
			g.SceneSort()
		})
	}
}
