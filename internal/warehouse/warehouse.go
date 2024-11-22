package warehouse

import "fmt"

// Warehouse represents a delivery warehouse
type Warehouse struct {
	ID       int
	Name     string
	Location string
}

// GetWarehouseDetails returns details about a warehouse
func GetWarehouseDetails(id int) string {
	// Placeholder logic
	return fmt.Sprintf("Warehouse %d: Central Warehouse at Location XYZ", id)
}
