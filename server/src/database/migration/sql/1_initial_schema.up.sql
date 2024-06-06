CREATE TABLE controller
(
    id          SERIAL PRIMARY KEY,
    placea        VARCHAR(64) NOT NULL
);

CREATE TABLE plants
(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(64) NOT NULL,
    temperature     INTEGER,
    moisture        INTEGER,
    humidity        INTEGER,
    lighting        INTEGER,
    created_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    controller_id   INTEGER REFERENCES controller(id)
);

