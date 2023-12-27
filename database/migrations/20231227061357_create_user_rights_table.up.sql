CREATE TABLE IF NOT EXISTS user_rights
(
  id          BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  user_id     BIGINT UNSIGNED NOT NULL REFERENCES users(id),
  election_id BIGINT UNSIGNED NOT NULL REFERENCES elections(id),
  is_used     BOOLEAN NOT NULL DEFAULT FALSE
);