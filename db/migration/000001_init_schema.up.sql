CREATE TABLE IF NOT EXISTS `users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NULL,
  `last_name` VARCHAR(45) NULL,
  `email_id` VARCHAR(45) NOT NULL,
  `password` VARCHAR(45) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_id_UNIQUE` (`email_id` ASC) VISIBLE
);

CREATE TABLE IF NOT EXISTS `timings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `end_time` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `regions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) DEFAULT NULL,
  `region_type` int DEFAULT NULL,
  `parent_id` int DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `parent_id_idx` (`parent_id`)
);

CREATE TABLE IF NOT EXISTS `theatres` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `address` varchar(45) NOT NULL,
  `region_id` int NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `region_id_idx` (`region_id`)
);

CREATE TABLE IF NOT EXISTS `halls` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `seats` int NOT NULL,
  `theatre_id` int NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_theatre_id_idx` (`theatre_id`)
);

CREATE TABLE IF NOT EXISTS `movies` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `director_name` VARCHAR(45) NOT NULL,
  `release_date` DATE NOT NULL,
  `is_active` TINYINT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS `shows` (
  `id` int NOT NULL AUTO_INCREMENT,
  `movie_id` int NOT NULL,
  `hall_id` int NOT NULL,
  `show_date` datetime NOT NULL,
  `timing_id` int NOT NULL,
  `seat_price` float NOT NULL,
  `available_seats` int DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_movie_id_idx` (`movie_id`),
  KEY `fk_hall_id_idx` (`hall_id`),
  KEY `fk_timing_id_idx` (`timing_id`)
);

CREATE TABLE IF NOT EXISTS `bookings` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `show_id` int NOT NULL,
  `seats` int DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_id_idx` (`user_id`),
  KEY `fk_show_id_idx` (`show_id`)
);

ALTER TABLE `regions` ADD CONSTRAINT `parent_id` FOREIGN KEY (`parent_id`) REFERENCES `regions` (`id`);

ALTER TABLE `theatres` ADD CONSTRAINT `fk_region_id` FOREIGN KEY (`region_id`) REFERENCES `regions` (`id`);

ALTER TABLE `halls` ADD  CONSTRAINT `fk_theatre_id` FOREIGN KEY (`theatre_id`) REFERENCES `theatres` (`id`);

ALTER TABLE `shows` ADD  CONSTRAINT `fk_hall_id` FOREIGN KEY (`hall_id`) REFERENCES `halls` (`id`);
ALTER TABLE `shows` ADD  CONSTRAINT `fk_movie_id` FOREIGN KEY (`movie_id`) REFERENCES `movies` (`id`);
ALTER TABLE `shows` ADD  CONSTRAINT `fk_timing_id` FOREIGN KEY (`timing_id`) REFERENCES `timings` (`id`);

ALTER TABLE `bookings` ADD CONSTRAINT `fk_show_id` FOREIGN KEY (`show_id`) REFERENCES `shows` (`id`);
ALTER TABLE `bookings` ADD CONSTRAINT `fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);