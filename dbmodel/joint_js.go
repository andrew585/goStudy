package dbmodel

type JoinJs struct {
	ID        uint `gorm:"primaryKey"`
	JoinStr   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
