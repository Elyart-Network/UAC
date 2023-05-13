package model

type Clients struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	UUID     string `json:"uuid" gorm:"not null;unique"`
	Name     string `json:"name" gorm:"not null;unique"`
	Secret   string `json:"secret" gorm:"not null;unique"`
	Type     string `json:"type" gorm:"not null"`
	Provider string `json:"provider"`
	Addition string `json:"addition"`
}

type Providers struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	UUID string `json:"uuid" gorm:"not null;unique"`
	Name string `json:"name" gorm:"not null;unique"`
	Type string `json:"type" gorm:"not null"`
	Data string `json:"data" gorm:"not null"`
}

type Credentials struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	UUID       string `json:"uuid" gorm:"not null;unique"`
	ClientID   string `json:"client_id" gorm:"not null;unique"`
	Credential string `json:"credential" gorm:"not null;unique"`
}

type Users struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	UUID     string `json:"uuid" gorm:"not null;unique"`
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null;unique"`
	Identity string `json:"identity" gorm:"not null;unique"`
	IsAdmin  bool   `json:"is_admin" gorm:"not null;default:false"`
	Addition string `json:"addition"`
}

type Tokens struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	UUID     string `json:"uuid" gorm:"not null;unique"`
	UserID   string `json:"user_id" gorm:"not null"`
	ClientID string `json:"client_id" gorm:"not null"`
	Token    string `json:"token" gorm:"not null;unique"`
	Expired  int64  `json:"expired" gorm:"not null"` // unix timestamp
	Type     string `json:"type" gorm:"not null"`    // access or refresh
	Addition string `json:"addition"`
}
