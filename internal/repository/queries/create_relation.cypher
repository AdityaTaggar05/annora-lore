CREATE (from:LoreNode{id:$from_id})-[rel:%s{
	description: $description,
	created_by: $created_by,
	created_at: $created_at,
	updated_at: $updated_at
}]->(to:LoreNode{id:$to_id})
RETURN rel
