package adapters

import (
	"time"
)

type Province struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"varchar(200)"`
}

type Store struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"varchar(200)"`
	UserID     uint
	ProvinceID uint
	Province   Province
	IsVerified bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Warehouse struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"varchar(250)"`
	StoreID    uint   `gorm:"index"`
	ProvinceID uint
	Province   Province
	Store      Store
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Media struct {
	ID        uint `grom:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductCategory struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"varchar(250)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}

type Product struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"varchar(250)"`
	Description   string
	StoreID       uint
	WarehouseID   uint
	SoldAvailable bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time `gorm:"index"`
}

type ProductVariation struct {
	ID        uint `gorm:"primaryKey"`
	Stock     uint
	Details   string
	Price     uint
	OldPrice  uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"index"`
}
