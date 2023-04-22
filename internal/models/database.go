package models

type Clients struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Name     string `json:"name" gorm:"not null;unique"`
	Secret   string `json:"secret" gorm:"not null;unique"`
	Type     string `json:"type" gorm:"not null"`
	Provider int    `json:"provider"`
	Addition string `json:"addition"`
}

type Providers struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Name string `json:"name" gorm:"not null;unique"`
	Type string `json:"type" gorm:"not null"`
	Data string `json:"data" gorm:"not null"`
}

type Credentials struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	ClientID   int    `json:"client_id" gorm:"not null;unique"`
	Credential string `json:"credential" gorm:"not null;unique"`
}

type Users struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Verified int64  `json:"verified" gorm:"not null;default:0"`
	Addition string `json:"addition"`
}

type Tokens struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null;unique"`
	UserID   int    `json:"user_id" gorm:"not null"`
	ClientID int    `json:"client_id" gorm:"not null"`
	Token    string `json:"token" gorm:"not null;unique"`
	Expired  int64  `json:"expired" gorm:"not null"` // unix timestamp
	Type     string `json:"type" gorm:"not null"`    // access or refresh
	Addition string `json:"addition"`
}
