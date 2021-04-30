CREATE TABLE IF NOT EXISTS teams (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `created_at` DATETIME default NOW(),
  `updated_at` DATETIME default NOW(),
  PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8;

CREATE INDEX teams_id_index USING BTREE ON teams (id);
