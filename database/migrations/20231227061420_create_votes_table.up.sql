CREATE TABLE IF NOT EXISTS votes
(
  id            BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  candidate_id  BIGINT UNSIGNED NOT NULL REFERENCES candidates(id) ON DELETE CASCADE,
  user_right_id BIGINT UNSIGNED NOT NULL REFERENCES user_rights(id) ON DELETE CASCADE,
  created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);