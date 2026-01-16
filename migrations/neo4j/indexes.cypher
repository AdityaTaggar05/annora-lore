CREATE INDEX lore_node_world_idx IF NOT EXISTS
FOR (n:LoreNode)
ON (n.world_id);

CREATE INDEX lore_node_type_idx IF NOT EXISTS
FOR (n:LoreNode)
ON (n.type);

CREATE INDEX relation_type_idx IF NOT EXISTS
FOR ()-[r]-()
ON (type(r));