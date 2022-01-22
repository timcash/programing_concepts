-- A schema is a collection of tables.
-- A table is a collection of columns.
-- A column is a named field in a table.

-- A table of commumities.

CREATE TABLE communities (
  id INTEGER PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- A table of members with foreign keys to the community table.

CREATE TABLE members (
    id INTEGER PRIMARY KEY,
    community_id INTEGER NOT NULL,
    member_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (community_id) REFERENCES communities(id),
);

-- Insert 3 example communities.
INSERT INTO communities (name, description) VALUES ('Example Community 1', 'This is an example community.');
INSERT INTO communities (name, description) VALUES ('Example Community 2', 'This is an example community.');
INSERT INTO communities (name, description) VALUES ('Example Community 3', 'This is an example community.');

-- Insert 3 example members.
INSERT INTO members (community_id, member_name) VALUES (1, 'Example Member 1');
INSERT INTO members (community_id, member_name) VALUES (1, 'Example Member 2');
INSERT INTO members (community_id, member_name) VALUES (1, 'Example Member 3');


