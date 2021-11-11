-- +migrate Up
CREATE TABLE users
(
    id              BIGINT       NOT NULL AUTO_INCREMENT,
    email           VARCHAR(255),
    name            VARCHAR(64)  NOT NULL,
    hashed_password VARCHAR(255) NOT NULL,
    role            VARCHAR(8)   NOT NULL DEFAULT 'member',
    created_at      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE (name)
);

-- +migrate Down
DROP TABLE users;
