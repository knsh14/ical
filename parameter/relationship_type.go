package parameter

type RelationshipTypeKind string

const (
	RelationshipTypeKindParent  RelationshipTypeKind = "PARENT"
	RelationshipTypeKindChild   RelationshipTypeKind = "CHILD"
	RelationshipTypeKindSibling RelationshipTypeKind = "SIBLING"
	RelationshipTypeKindXToken  RelationshipTypeKind = "XToken"
)
