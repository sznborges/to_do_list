``` sql
-- +goose Up 
CREATE TABLE tasks (
  id varchar PRIMARY KEY,
  title varchar NOT NULL, 
  description_task varchar, 
  completed boolean,
  created_at timestamp with time zone NOT NULL,
);

-- +goose Down 
DROP TABLE tasks;
```
