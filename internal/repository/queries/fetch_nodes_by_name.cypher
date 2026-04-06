CALL db.index.fulltext.queryNodes('lorenode_name', $query)
YIELD node, score
WHERE node.world_id = $wid
RETURN node, score ORDER BY score DESC;
