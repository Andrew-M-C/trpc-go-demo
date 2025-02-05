
CREATE TABLE IF NOT EXISTS `t_trpc_demo_user_reputation` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
    `user_id` varchar(128) NOT NULL COMMENT '用户 ID',
    `exp_points` int unsigned NOT NULL COMMENT '声望点数',
    `rank_level` smallint unsigned NOT NULL COMMENT '等级级别',
    PRIMARY KEY (`id`),
    UNIQUE KEY `u_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户声望表';
