-- +migrate Up
CREATE TABLE testsets
(
    id         BIGINT      NOT NULL AUTO_INCREMENT,
    problem_id BIGINT      NOT NULL,
    name       VARCHAR(64) NOT NULL,
    sample     BOOLEAN     NOT NULL DEFAULT false,
    score      INT         NOT NULL DEFAULT 0,
    created_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (problem_id, name)
);

-- +migrate Down
DROP TABLE testsets;
