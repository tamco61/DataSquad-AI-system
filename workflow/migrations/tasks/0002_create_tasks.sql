-- +goose Up
CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Исходное письмо
    title TEXT NOT NULL,
    raw_text TEXT NOT NULL,

    -- AI аналитика
    classification TEXT,
    summary TEXT,
    risk_level TEXT CHECK (risk_level IN ('low', 'medium', 'high')),

    -- Workflow
    assignee_id UUID REFERENCES users(id) ON DELETE SET NULL,
    addressed_id UUID REFERENCES users(id) ON DELETE SET NULL,
    status TEXT NOT NULL DEFAULT 'new'
        CHECK (status IN ('new', 'in_progress', 'wait_approval', 'done')),

    sla_deadline TIMESTAMP,
    resolved_at TIMESTAMP,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS tasks;