package model

import "time"

type RelationType string

type Relation struct {
	FromID      string       `json:"from_id"`
	ToID        string       `json:"to_id"`
	Rel         RelationType `json:"rel"`
	Description string       `json:"description"`
	CreatedBy   string       `json:"created_by"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

var AllowedRelations = map[NodeType]map[NodeType]map[RelationType]bool{
	NodeTypePlace: {
		NodeTypePlace: {
			"LOCATED_IN":   true,
			"CONNECTED_TO": true,
		},
		NodeTypeSpecies: {
			"HABITAT_OF": true,
		},
		NodeTypeArtifact: {
			"HOUSES":        true,
			"DISCOVERED_AT": true,
		},
		NodeTypeFaction: {
			"CONTROLLED_BY": true,
			"CONTESTED_BY":  true,
		},
		NodeTypeEvent: {
			"SITE_OF":      true,
			"DESTROYED_IN": true,
		},
	},
	NodeTypeSpecies: {
		NodeTypePlace: {
			"NATIVE_TO":   true,
			"MIGRATED_TO": true,
		},
		NodeTypeFaction: {
			"OPPRESSED_BY":        true,
			"DOMINANT_SPECIES_OF": true,
		},
		NodeTypeEvent: {
			"EMERGED_IN":   true,
			"DECIMATED_IN": true,
		},
	},
	NodeTypeCharacter: {
		NodeTypePlace: {
			"BORN_IN":    true,
			"RESIDES_IN": true,
			"DIED_IN":    true,
		},
		NodeTypeCharacter: {
			"ALLY_OF":  true,
			"ENEMY_OF": true,
			"KIN_OF":   true,
		},
		NodeTypeArtifact: {
			"WIELDS":    true,
			"CREATED":   true,
			"DESTROYED": true,
		},
		NodeTypeFaction: {
			"MEMBER_OF":   true,
			"LEADER_OF":   true,
			"FOUNDER_OF":  true,
			"BETRAYER_OF": true,
		},
		NodeTypeEvent: {
			"INVOLVED_IN": true,
			"CAUSED":      true,
			"DIED_IN":     true,
		},
	},
	NodeTypeArtifact: {
		NodeTypePlace: {
			"DISCOVERED_AT": true,
			"SEALED_IN":     true,
			"DESTROYED_AT":  true,
			"FORGED_IN":     true,
		},
		NodeTypeCharacter: {
			"OWNED_BY":   true,
			"WIELDED_BY": true,
			"CREATED_BY": true,
		},
		NodeTypeFaction: {
			"SYMBOL_OF":  true,
			"GUARDED_BY": true,
		},
		NodeTypeEvent: {
			"USED_IN":      true,
			"CREATED_FOR":  true,
			"DESTROYED_IN": true,
		},
	},
	NodeTypeFaction: {
		NodeTypePlace: {
			"CONTROLS":   true,
			"FOUNDED_AT": true,
		},
		NodeTypeArtifact: {
			"OWNS":  true,
			"SEEKS": true,
		},
		NodeTypeFaction: {
			"ALLIED_WITH":  true,
			"ENEMIES_WITH": true,
			"VASSAL_OF":    true,
			"LIEGE_OF":     true,
		},
		NodeTypeEvent: {
			"INVOLVED_IN":      true,
			"INITIATED":        true,
			"COLLAPSED_DURING": true,
		},
	},
	NodeTypeEvent: {
		NodeTypePlace: {
			"OCCURRED_AT": true,
			"ALTERED":     true,
			"DESTROYED":   true,
		},
		NodeTypeCharacter: {
			"INVOLVED": true,
			"KILLED":   true,
		},
		NodeTypeSpecies: {
			"AFFECTED":  true,
			"CREATED":   true,
			"DECIMATED": true,
		},
		NodeTypeFaction: {
			"DISSOLVED": true,
			"CREATED":   true,
		},
		NodeTypeEvent: {
			"TRIGGERED":    true,
			"SUCCEEDED_BY": true,
			"PRECEDED_BY":  true,
		},
	},
}
