-- +goose Up
ALTER TABLE feeds
ADD COLUMN id UUID PRIMARY KEY DEFAULT gen_random_uuid();

-- +goose Down
ALTER TABLE feeds
DROP COLUMN id;
