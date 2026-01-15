package models

type NodeType string
type CanonStatus string

const (
	NodeTypePlace     NodeType = "place"
	NodeTypeSpecies   NodeType = "species"
	NodeTypeCharacter NodeType = "character"
	NodeTypeArtifact  NodeType = "artifact"
	NodeTypeFaction   NodeType = "faction"
	NodeTypeEvent     NodeType = "event"
)

const (
	CanonStatusAccepted CanonStatus = "accepted"
	CanonStatusRejected CanonStatus = "rejected"
	CanonStatusPending  CanonStatus = "pending"
)
