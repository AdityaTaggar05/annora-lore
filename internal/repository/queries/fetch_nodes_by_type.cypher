MATCH (n:LoreNode{world_id:$world_id})
WHERE n.type IN $node_types
RETURN n
