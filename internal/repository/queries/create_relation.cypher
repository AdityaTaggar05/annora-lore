MATCH (w:WorldNode{id:$world_id})
MATCH (w)-[:CONTAINS]->(from:LoreNode{id:$from_id})
MATCH (w)-[:CONTAINS]->(to:LoreNode{id:$to_id})
CREATE p=(from)-[:$relation{
	description: $description,
	created_at: $created_at,
	updated_at: $updated_at
}]->(to)
RETURN p