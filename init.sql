CREATE TABLE IF NOT EXISTS file_content (
    id VARCHAR(255) PRIMARY KEY,
    hash VARCHAR(64) NOT NULL,
    content JSONB NOT NULL
    );

CREATE INDEX IF NOT EXISTS idx_file_content_hash ON file_content (hash);