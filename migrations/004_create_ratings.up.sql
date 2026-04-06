CREATE TABLE ratings (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    recipe_id  UUID NOT NULL REFERENCES recipes(id) ON DELETE CASCADE,
    user_id    UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    score      INT NOT NULL CHECK (score >= 1 AND score <= 5),
    comment    TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (recipe_id, user_id)
);
