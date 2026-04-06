MATCH (n:LoreNode{id: $node_id})-[:%s]->(to: LoreNode)
RETURN to;
