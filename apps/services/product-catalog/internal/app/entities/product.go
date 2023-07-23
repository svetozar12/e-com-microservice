package entities

type ProductEntity struct {
	Model
	Name        string `json:"name" binding:"required"`
	Price       int32  `json:"price" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Description string `json:"description" binding:"required"`
	Available   bool   `json:"available" binding:"required"`
	Weight      int32  `json:"weight" binding:"required"`
	// Inventory entity
	Currency string `json:"currency" binding:"required"`
}
