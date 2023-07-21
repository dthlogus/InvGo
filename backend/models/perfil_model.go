package models

type Perfil struct {
	Pv       float32 `json:"pv"`
	Pvp      float32 `json:"pvp"`
	Dy       float32 `json:"dy"`
	Lpa      float32 `json:"lpa"`
	Vpa      float32 `json:"vpa"`
	Simbolos string  `json:"simbolos"`
}
