-- +goose Up
-- +goose StatementBegin
CREATE TABLE courses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    course_code VARCHAR(255) NOT NULL UNIQUE,
    credit_hours INT NOT NULL,
    contact_hours INT NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE courses_pre_requisites (
    course_id UUID NOT NULL REFERENCES courses(id),
    pre_requisite_id UUID NOT NULL REFERENCES courses(id),
    PRIMARY KEY (course_id, pre_requisite_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE courses_pre_requisites;
DROP TABLE courses;
-- +goose StatementEnd
