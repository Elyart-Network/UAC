package models

type Clients struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name     string `json:"name" gorm:"not null"`
	Secret   string `json:"secret" gorm:"not null"`
	Type     string `json:"type" gorm:"not null"`
	Provider int    `json:"provider"`
	Addition string `json:"addition"`
}

type Providers struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name string `json:"name" gorm:"not null"`
	Type string `json:"type" gorm:"not null"`
	Data string `json:"data" gorm:"not null"`
}

type Credentials struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	ClientID   int    `json:"client_id" gorm:"not null"`
	Credential string `json:"credential" gorm:"not null"`
}

type Users struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email" gorm:"not null"`
	Addition string `json:"addition"`
}

type Tokens struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserID   int    `json:"user_id" gorm:"not null"`
	ClientID int    `json:"client_id" gorm:"not null"`
	Token    string `json:"token" gorm:"not null"`
	Expired  int64  `json:"expired" gorm:"not null"` // unix timestamp
	Type     string `json:"type" gorm:"not null"`    // access or refresh
	Addition string `json:"addition"`
}
