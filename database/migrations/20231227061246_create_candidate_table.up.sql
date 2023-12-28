CREATE TABLE IF NOT EXISTS candidates
(
  id          BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  leader      VARCHAR(255) NOT NULL,
  vice        VARCHAR(255) NOT NULL,
  image       VARCHAR(255) NOT NULL,  -- image url
  election_id BIGINT UNSIGNED NOT NULL REFERENCES elections(id),
  description TEXT NOT NULL
)