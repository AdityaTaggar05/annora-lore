// Composite Indexes
CREATE INDEX lorenode_world_type FOR (n:LoreNode) ON (n.world_id, n.type);
CREATE INDEX lorenode_world_status FOR (n:LoreNode) ON (n.world_id, n.canon_status);
CREATE INDEX lorenode_world_type_status FOR (n:LoreNode) ON (n.world_id, n.type, n.canon_status);
CREATE INDEX lorenode_world_created FOR (n:LoreNode) ON (n.world_id, n.created_at);

// Full Text Search Index
CREATE FULLTEXT INDEX lorenode_name FOR (n:LoreNode) ON EACH [n.name];
