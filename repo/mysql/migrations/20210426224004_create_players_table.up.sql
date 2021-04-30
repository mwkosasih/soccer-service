CREATE TABLE IF NOT EXISTS `players` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `team_id` int(10) unsigned NOT NULL,
  `name` varchar(255) NOT NULL,
  `position` varchar(255) NOT NULL,
  `height` int(3) NOT NULL,
  `weight` int(3) NOT NULL,
  `created_at` DATETIME default NOW(),
  `updated_at` DATETIME default NOW(),
  PRIMARY KEY (`id`),
  KEY `fk_players_teams1_idx` (`team_id`),
  CONSTRAINT `fk_players_teams1` FOREIGN KEY (`team_id`) REFERENCES `teams` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8;

CREATE INDEX players_id_index USING BTREE ON players (id);
CREATE INDEX players_team_id_index USING BTREE ON players (team_id);
