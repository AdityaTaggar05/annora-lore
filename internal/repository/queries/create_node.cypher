MATCH (w:WorldNode {id: $world_id})
CREATE (n:LoreNode {
	id: $id,
	type: $type,
	name: $name,
	world_id: $world_id,
	created_by: $created_by,
	canon_status: $canon_status,
	created_at: datetime($created_at),
	updated_at: datetime($updated_at),
	custom: "null"
})
CREATE (w)-[:CONTAINS]->(n)
RETURN n