MATCH (w: WorldNode{id: $world_id})-[:CONTAINS]->(n: LoreNode)
WHERE n.type IN $node_types
RETURN n
