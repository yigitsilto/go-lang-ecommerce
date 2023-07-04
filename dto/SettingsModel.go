package dto

type SettingsModel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GeneralSettingsModel struct {
	Settings map[string]string `json:"settings"`
	Menu     []MenuModel       `json:"menu"`
	Footer1  []MenuModel       `json:"footer1"`
	Footer2  []MenuModel       `json:"footer2"`
}

type MenuModel struct {
	Id       uint        `json:"id"`
	Type     string      `json:"type"`
	Url      string      `json:"url"`
	Name     string      `json:"name"`
	IsRoot   string      `json:"is_root"`
	ParentId uint        `json:"parentId"`
	Items    []MenuModel `json:"items" gorm:"-"`
}
