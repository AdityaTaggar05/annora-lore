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

var AllowedTypes map[NodeType]bool = map[NodeType]bool{
	NodeTypePlace:     true,
	NodeTypeSpecies:   true,
	NodeTypeCharacter: true,
	NodeTypeArtifact:  true,
	NodeTypeFaction:   true,
	NodeTypeEvent:     true,
}

