// ID Uniquness Constraint (implicitly creates indexes)
CREATE CONSTRAINT world_id_unique IF NOT EXISTS FOR (w:WorldNode) REQUIRE w.id IS UNIQUE;
CREATE CONSTRAINT lorenode_id_unique IF NOT EXISTS FOR (n:LoreNode) REQUIRE n.id IS UNIQUE;
