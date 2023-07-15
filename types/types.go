package types

type Config struct {
	IP         string      `toml:"ip" json:"ip"`
	Port       int         `toml:"port" json:"port"`
	Secret     string      `toml:"secret" json:"secret"`
	SiteGroups []SiteGroup `toml:"groups" json:"groups"`
	Users      []User      `toml:"users" json:"users"`
}

type SiteGroup struct {
	Name  string `toml:"name" json:"name"`
	Icon  string `toml:"icon" json:"icon"`
	Sites []Site `toml:"sites" json:"sites"`
}

type Site struct {
	Name string `toml:"name" json:"name"`
	Icon string `toml:"icon" json:"icon"`
	Desc string `toml:"desc" json:"desc"`
	Path string `toml:"path" json:"path"`
	Port int    `toml:"port" json:"-"`
}

type User struct {
	Name     string `json:"name" gorm:"unique" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}
