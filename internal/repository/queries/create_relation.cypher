MATCH (from:LoreNode{id:$from_id})
MATCH (to:LoreNode{id:$to_id})
CREATE (from)-[rel:%s{
	description: $description,
	created_by: $created_by,
	created_at: $created_at,
	updated_at: $updated_at
}]->(to)
RETURN from, rel, to
