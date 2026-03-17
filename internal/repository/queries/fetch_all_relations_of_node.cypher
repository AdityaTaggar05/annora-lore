MATCH (n:LoreNode {id: $node_id})-[r]->(to:LoreNode)
RETURN r, to;
