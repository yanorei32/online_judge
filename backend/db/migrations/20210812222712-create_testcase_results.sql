-- +migrate Up
CREATE TABLE testcase_results
(
    id               BIGINT     NOT NULL AUTO_INCREMENT,
    submission_id    BIGINT     NOT NULL,
    testcase_id      BIGINT     NOT NULL,
    status           VARCHAR(8) NOT NULL,
    execution_time   BIGINT     NOT NULL,
    execution_memory BIGINT     NOT NULL,
    created_at       DATETIME   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       DATETIME   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (submission_id, testcase_id)
);

-- +migrate Down
DROP TABLE testcase_results;
