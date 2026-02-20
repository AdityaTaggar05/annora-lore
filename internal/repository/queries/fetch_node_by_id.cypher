MATCH (w: WorldNode{id: $world_id})-[:CONTAINS]->(node: LoreNode{id: $node_id})
RETURN node
