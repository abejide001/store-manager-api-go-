package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Product struct
type Product struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"size:255;not null;unique" json:"name"`
	Description string    `gorm:"size:255;not null;" json:"description"`
	Price       uint32    `json:"price"`
	Attendant   User      `gorm:"not null" json:"attendant"`
	AttendantID uint32    `gorm:"not null" json:"attendant_id"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Prepare Method
func (p *Product) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.Attendant = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

// Validate Method
func (p *Product) Validate() error {

	if p.Name == "" {
		return errors.New("Required Name")
	}
	if p.Description == "" {
		return errors.New("Required Description")
	}
	if p.Price < 1 {
		return errors.New("Required Price")
	}
	return nil
}

// SaveProduct Method
func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AttendantID).Take(&p.Attendant).Error
		if err != nil {
			return &Product{}, err
		}
	}
	return p, nil
}

// FindAllProducts Method
func (p *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	if len(products) > 0 {
		for i := range products {
			err := db.Debug().Model(&User{}).Where("id = ?", products[i].AttendantID).Take(&products[i].Price).Error
			if err != nil {
				return &[]Product{}, err
			}
		}
	}
	return &products, nil
}

// FindProductsByID method
func (p *Product) FindProductsByID(db *gorm.DB, pid uint64) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AttendantID).Take(&p.Attendant).Error
		if err != nil {
			return &Product{}, err
		}
	}
	return p, nil
}

// UpdateAProduct Method
func (p *Product) UpdateAProduct(db *gorm.DB, pid uint64) (*Product, error) {

	var err error
	db = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&Product{}).UpdateColumns(
		map[string]interface{}{
			"name":        p.Name,
			"description": p.Description,
			"updated_at":  time.Now(),
		},
	)
	err = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AttendantID).Take(&p.Attendant).Error
		if err != nil {
			return &Product{}, err
		}
	}
	return p, nil
}

// DeleteAProduct Method
func (p *Product) DeleteAProduct(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Product{}).Where("id = ? and attendant_id = ?", pid, uid).Take(&Product{}).Delete(&Product{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Product not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
