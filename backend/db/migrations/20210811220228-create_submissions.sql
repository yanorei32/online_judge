-- +migrate Up
CREATE TABLE submissions
(
    id               BIGINT        NOT NULL AUTO_INCREMENT,
    problem_id       BIGINT        NOT NULL,
    user_id          BIGINT        NOT NULL,
    language         VARCHAR(32)   NOT NULL,
    status           VARCHAR(8)    NOT NULL,
    score            INT           NOT NULL,
    execution_time   BIGINT        NOT NULL,
    execution_memory BIGINT        NOT NULL,
    compile_log      VARCHAR(1024) NOT NULL,
    created_at       DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE submissions;
