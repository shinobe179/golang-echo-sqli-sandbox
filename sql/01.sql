CREATE DATABASE IF NOT EXISTS `someservice`;
CREATE DATABASE IF NOT EXISTS `secret`;
CREATE USER kid IDENTIFIED BY 'iamkid';

GRANT ALL PRIVILEGES ON *.* TO 'kid'@'%';

USE `someservice`;
CREATE TABLE `users` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4;

INSERT INTO `users` (`name`) VALUES 
    ('Alice'),
    ('Bob'),
    ('Charly'),
    ('David');

USE `secret`;
CREATE TABLE `flags` (
    `flag` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`flag`)
) ENGINE=InnoDB DEFAULT CHARACTER SET=utf8mb4;

INSERT INTO `flags` (`flag`) VALUES 
    ('FLAG{SQLi_is_succeeded}');
