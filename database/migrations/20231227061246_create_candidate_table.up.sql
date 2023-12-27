CREATE TABLE IF NOT EXISTS candidates
(
  id          BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  election_id BIGINT UNSIGNED NOT NULL REFERENCES elections(id),
  name        VARCHAR(255) NOT NULL,
  description TEXT NOT NULL
)