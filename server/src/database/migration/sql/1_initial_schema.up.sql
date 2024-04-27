CREATE TABLE controller
(
    id       SERIAL PRIMARY KEY,
    place    VARCHAR(64) NOT NULL,
    moisture INTEGER
);

CREATE TABLE plants
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(64) NOT NULL,
    moisture    INTEGER,
    humidity    INTEGER,
    lighting    INTEGER,
    created_at  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    controller_id INTEGER REFERENCES controller(id)
);



