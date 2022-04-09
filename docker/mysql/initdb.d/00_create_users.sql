DROP TABLE IF EXISTS users;

CREATE TABLE users (
  `id` INT unsigned NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `month` int NOT NULL,
  `day` int NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO users (id, name, month, day) VALUES (1, "Yamada", "1", "1");
INSERT INTO users (id, name, month, day) VALUES (2, "Tanaka", "12", "31");