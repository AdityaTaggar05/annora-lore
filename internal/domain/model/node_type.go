package model

type NodeType string

const (
	NodeTypePlace     NodeType = "place"
	NodeTypeSpecies   NodeType = "species"
	NodeTypeCharacter NodeType = "character"
	NodeTypeArtifact  NodeType = "artifact"
	NodeTypeFaction   NodeType = "faction"
	NodeTypeEvent     NodeType = "event"
)
