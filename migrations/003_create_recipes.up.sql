CREATE TABLE recipes (
    id           UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id      UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    coffee_id    UUID NOT NULL REFERENCES coffees(id) ON DELETE RESTRICT,
    method       TEXT NOT NULL,
    water_temp   INT NOT NULL,
    dose         NUMERIC(5,2) NOT NULL,
    yield        NUMERIC(5,2) NOT NULL,
    brew_time    INT NOT NULL,
    description  TEXT,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now()
);
