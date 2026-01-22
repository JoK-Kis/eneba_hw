package models

type Platform struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"unique;not null" json:"name"`
	IconURL   string `gorm:"not null" json:"icon"`
}

type Product struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Name            string    `gorm:"not null" json:"name"`
	ImageURL        string    `gorm:"not null" json:"image"`
	PlatformID      uint      `gorm:"not null;constraint:OnDelete:RESTRICT" json:"platform_id"`
	Platform        *Platform `gorm:"foreignKey:PlatformID" json:"platform,omitempty"`
	Region          string    `gorm:"not null" json:"location"`
	OriginalPrice   float64   `gorm:"type:decimal(10,2)" json:"originalPrice"`
	CurrentPrice    float64   `gorm:"type:decimal(10,2)" json:"currentPrice"`
	DiscountPercent int       `gorm:"default:0" json:"discount"`
	CashbackAmount  float64   `gorm:"type:decimal(10,2)" json:"cashbackAmount"`
	Likes           int       `gorm:"default:0" json:"likes"`
	InStock         bool      `gorm:"default:true" json:"inStock"`
}