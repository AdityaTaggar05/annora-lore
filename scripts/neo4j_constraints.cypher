CREATE CONSTRAINT world_id_unique IF NOT EXISTS
FOR (w:WorldNode)
REQUIRE w.id IS UNIQUE;

CREATE CONSTRAINT lore_node_id_unique IF NOT EXISTS
FOR (n:LoreNode)
REQUIRE n.id IS UNIQUE;

CREATE CONSTRAINT lore_node_type_required IF NOT EXISTS
FOR (n:LoreNode)
REQUIRE n.type IS NOT NULL;

CREATE CONSTRAINT lore_node_world_required IF NOT EXISTS
FOR (n:LoreNode)
REQUIRE n.world_id IS NOT NULL;

CREATE CONSTRAINT relation_creator_required IF NOT EXISTS
FOR ()-[r]-()
REQUIRE r.created_by IS NOT NULL;

CREATE INDEX lore_node_world_idx IF NOT EXISTS
FOR (n:LoreNode)
ON (n.world_id);

CREATE INDEX lore_node_type_idx IF NOT EXISTS
FOR (n:LoreNode)
ON (n.type);

CREATE INDEX relation_type_idx IF NOT EXISTS
FOR ()-[r]-()
ON (type(r));
