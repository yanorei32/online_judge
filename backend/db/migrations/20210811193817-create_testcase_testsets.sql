-- +migrate Up
CREATE TABLE testcase_testsets
(
    id          BIGINT   NOT NULL AUTO_INCREMENT,
    problem_id  BIGINT   NOT NULL,
    testcase_id BIGINT   NOT NULL,
    testset_id  BIGINT   NOT NULL,
    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (problem_id, testcase_id, testset_id)
);

-- +migrate Down
DROP TABLE testcase_testsets;
