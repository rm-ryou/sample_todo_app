-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `todos` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `board_id` INT NOT NULL,
  `title` VARCHAR(50) NOT NULL,
  `done` BOOLEAN NOT NULL DEFAULT false,
  `priority` INT NOT NULL DEFAULT 0,
  `due_date` DATETIME,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `idx_board_id` (`board_id`),
  FOREIGN KEY (`board_id`) REFERENCES boards(`id`) ON DELETE CASCADE
) ENGINE=INNODB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `todos`;
-- +goose StatementEnd
