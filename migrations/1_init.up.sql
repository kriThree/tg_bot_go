CREATE TABLE IF NOT EXISTS definitions (
    id TEXT PRIMARY KEY,
    word TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE IF NOT EXISTS meaning (
    id TEXT PRIMARY KEY,
    definition_id TEXT NOT NULL,
    part_of_speach TEXT,
    value TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (definition_id) REFERENCES definitions(id)
)