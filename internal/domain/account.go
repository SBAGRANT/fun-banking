package domain

type Account struct {
	Audit
	Name           string   `gorm:"not null; default:Checking"`
	CurrentBalance float64  `gorm:"decimal(10,2); default:0.00"`
	CustomerID     uint     `gorm:"not null"`
	Customer       Customer `gorm:"foreignKey:CustomerID; constraint:OnDelete:CASCADE"`
}
