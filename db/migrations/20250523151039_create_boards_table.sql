-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `boards` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `room_id` INT NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  `priority` INT NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`room_id`) REFERENCES rooms(`id`) ON DELETE CASCADE
) ENGINE=INNODB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `boards`;
-- +goose StatementEnd
