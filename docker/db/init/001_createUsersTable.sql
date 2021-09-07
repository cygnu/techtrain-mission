DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
    `id`         INT(12) NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(50) NOT NULL,
    `token`      VARCHAR(50) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
