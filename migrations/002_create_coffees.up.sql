CREATE TABLE coffees (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name          TEXT NOT NULL,
    brand         TEXT NOT NULL,
    type          TEXT NOT NULL,
    flavor_notes  TEXT,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT now()
);
