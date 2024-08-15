package model

// Quantity counts a number of objects.
type Quantity int

// Warehouse is used to store different products.
type Warehouse struct {

	// Id of a warehouse.
	Id Identifier

	// MaximumStorageCapacity is the max number of objects a warehouse can store.
	// Hpw much space a good requires is defined separately.
	MaximumStorageCapacity int

	// StorageUnits, list of good and stored quantity.
	StorageUnits map[Goods]Quantity
}
