-- +migrate Up
ALTER TABLE submissions MODIFY status VARCHAR(8);
ALTER TABLE submissions MODIFY score INT;
ALTER TABLE submissions MODIFY execution_time BIGINT;
ALTER TABLE submissions MODIFY execution_memory BIGINT;
ALTER TABLE submissions MODIFY compile_log VARCHAR(1024);

-- +migrate Down
ALTER TABLE submissions MODIFY status VARCHAR(8) NOT NULL;
ALTER TABLE submissions MODIFY score INT NOT NULL;
ALTER TABLE submissions MODIFY execution_time BIGINT NOT NULL;
ALTER TABLE submissions MODIFY execution_memory BIGINT NOT NULL;
ALTER TABLE submissions MODIFY compile_log VARCHAR(1024) NOT NULL;
