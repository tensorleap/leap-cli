package entity

// EntityDescriptor describes an entity and provides a standardized way to handle generic inputs and messaging related to the entity.
type EntityDescriptor[TEntity any] struct {
	Name           string
	PluralName     string
	GetDisplayName func(e *TEntity) string
	GetID          func(e *TEntity) string
}

func NewEntityDescriptor[TEntity any](
	name, pluralName string,
	getDisplayName func(*TEntity) string,
	getId func(*TEntity) string,
) *EntityDescriptor[TEntity] {
	return &EntityDescriptor[TEntity]{
		Name:           name,
		PluralName:     pluralName,
		GetDisplayName: getDisplayName,
		GetID:          getId,
	}
}
