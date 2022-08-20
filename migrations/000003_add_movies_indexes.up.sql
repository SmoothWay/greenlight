CREATE INDEX IF NOT EXISTs movies_title_idx ON movies USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTs movies_genres_idx ON movies USING GIN (genres);