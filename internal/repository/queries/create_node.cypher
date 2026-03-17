MATCH (w:WorldNode {id: $world_id})
CREATE (n:LoreNode {
	id: $id,
	type: $type,
	name: $name,
	world_id: $world_id,
	created_by: $created_by,
  creator_username: $creator_username,
	canon_status: $canon_status,
	created_at: datetime($created_at),
	updated_at: datetime($updated_at),
})
CREATE (w)-[:CONTAINS]->(n)
RETURN n
