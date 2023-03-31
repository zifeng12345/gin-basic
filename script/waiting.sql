DROP TABLE IF EXISTS `waiting_lists`;

CREATE TABLE `waiting_lists` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `store_id` varchar(255) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  `day` varchar(255) DEFAULT NULL,
  `pax` int DEFAULT '0',
  `phone_number` varchar(255) DEFAULT NULL,
  `number` int DEFAULT '0',
  `user_name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_store_id` (`store_id`),
  KEY `idx_number` (`number`),
  KEY `idx_status` (`status`),
  KEY `idx_day` (`day`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8;