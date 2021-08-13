-- +migrate Up
CREATE TABLE problems
(
    id           BIGINT        NOT NULL AUTO_INCREMENT,
    name         VARCHAR(64)   NOT NULL,
    public       BOOLEAN       NOT NULL DEFAULT false,
    writer_id    BIGINT        NOT NULL,
    time_limit   BIGINT        NOT NULL DEFAULT 2000,
    memory_limit BIGINT        NOT NULL DEFAULT 256,
    statement    VARCHAR(4096) NOT NULL,
    limitation   VARCHAR(2048) NOT NULL,
    input        VARCHAR(1024) NOT NULL,
    output       VARCHAR(1024) NOT NULL,
    created_at   DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   DATETIME      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE problems;
