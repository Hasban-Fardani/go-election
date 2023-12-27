CREATE TABLE IF NOT EXISTS elections
(
  id              BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  name            VARCHAR(255) NOT NULL,
  description     TEXT,
  number_of_votes BIGINT UNSIGNED NOT NULL DEFAULT 0,
  is_active       BOOLEAN NOT NULL DEFAULT FALSE,
  start_time      TIMESTAMP NOT NULL,
  end_time        TIMESTAMP NOT NULL
);