create database tticket;

CREATE TABLE `user` (
                        `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                        `name` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                        `mail` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                        `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `deleted_time` timestamp NULL DEFAULT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `ball` (
                        `id` int unsigned NOT NULL AUTO_INCREMENT,
                        `lottery_drawing_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
                        `num1` int DEFAULT NULL,
                        `num2` int DEFAULT NULL,
                        `num3` int DEFAULT NULL,
                        `num4` int unsigned DEFAULT NULL,
                        `num5` int unsigned DEFAULT NULL,
                        `num6` int DEFAULT NULL,
                        `num7` int DEFAULT NULL,
                        `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `deleted_time` timestamp NULL DEFAULT NULL,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_drawing_time` (`lottery_drawing_time`)
) ENGINE=InnoDB AUTO_INCREMENT=411 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `predict_ball` (
                                `id` int unsigned NOT NULL AUTO_INCREMENT,
                                `lottery_drawing_time` varchar(255) DEFAULT NULL,
                                `num1` int DEFAULT NULL,
                                `num2` int DEFAULT NULL,
                                `num3` int DEFAULT NULL,
                                `num4` int DEFAULT NULL,
                                `num5` int DEFAULT NULL,
                                `num6` int DEFAULT NULL,
                                `num7` int DEFAULT NULL,
                                `predict_lottery_drawing_time` varchar(255) DEFAULT NULL,
                                `strategy` varchar(255) DEFAULT NULL,
                                `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                `deleted_time` timestamp NULL DEFAULT NULL,
                                PRIMARY KEY (`id`),
                                UNIQUE KEY `idx_drawing_time` (`lottery_drawing_time`),
                                UNIQUE KEY `idx_predict_drawding_time` (`predict_lottery_drawing_time`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `task` (
                        `id` int unsigned NOT NULL AUTO_INCREMENT,
                        `name` varchar(255) DEFAULT NULL,
                        `interval_second` int DEFAULT NULL,
                        `type` int DEFAULT NULL,
                        `executor` varchar(255) DEFAULT NULL,
                        `execute_time` timestamp NULL DEFAULT NULL,
                        `cron` varchar(255) DEFAULT NULL,
                        `created_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `updated_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `deleted_time` timestamp NULL DEFAULT NULL,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_name_type` (`name`,`type`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

insert into task (`name`, `interval_second`, `type`, `cron`) values ('cache_user', 1200, 2, '');

insert into task (`name`, `interval_second`, `type`, `cron`) values ('spider_lottery', 0, 3, '0 22 1,3,5');

insert into task (`name`, `interval_second`, `type`, `cron`) values ('predict_ball', 0, 3, '0 11 2,4,7');

insert into task (`name`, `interval_second`, `type`, `cron`) values ('send_mail', 0, 3, '0 12 2,4,7');
