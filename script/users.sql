DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(50) DEFAULT '',
    `password` varchar(50) DEFAULT '',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `users` (
    `id`,
    `username`,
    `password`
) VALUES (
    '1',
    'jack',
    'jack'
);