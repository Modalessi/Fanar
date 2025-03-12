-- +goose Up
-- +goose StatementBegin
CREATE TABLE resources (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL REFERENCES courses(id),
    title VARCHAR(255) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    file_ext VARCHAR(255) NOT NULL,
    s3_url VARCHAR(255) NOT NULL,
    tags TEXT[] NOT NULL,
    created_by UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE resources;
-- +goose StatementEnd
