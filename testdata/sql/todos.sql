-- Create todos table
CREATE TABLE IF NOT EXISTS `todos` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(50) NOT NULL,
  `done` BOOLEAN NOT NULL DEFAULT false,
  `priority` INT NOT NULL DEFAULT 0,
  `due_date` DATETIME,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=INNODB;

-- Insert test data
INSERT INTO `todos` (
  `title`,
  `done`,
  `priority`,
  `created_at`,
  `updated_at`
) VALUES (
  "Test Task",
  FALSE,
  3,
  "2025-05-01 10:00:00",
  "2025-05-01 10:00:00"
);
