package model

// Identifier is a generic id for different tupes of objects.
type Identifier string

// String return string value of an identifier.
func (id Identifier) String() string {
	return string(id)
}
