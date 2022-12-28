CREATE DATABASE IF NOT EXISTS `someservice`;
CREATE DATABASE IF NOT EXISTS `secret`;
CREATE USER kid IDENTIFIED BY 'iamkid';

GRANT ALL PRIVILEGES ON *.* TO 'kid'@'%';

USE `someservice`;
CREATE TABLE `users` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `is_visible` BOOLEAN NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4;

INSERT INTO `users` (`name`, `is_visible`) VALUES 
    ('Alice', true),
    ('Bob', true),
    ('Charly', true),
    ('David', true),
    ('Mallory', false),
    ('Victor', false),
    ('FLAG{7H15_15_1NV151BL3_US3R}', false);

USE `secret`;
CREATE TABLE `flags` (
    `flag` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`flag`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4;

INSERT INTO `flags` (`flag`) VALUES 
    ('FLAG{SQL1_15_5UCC33D3D}');
