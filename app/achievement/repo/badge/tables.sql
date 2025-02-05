
CREATE TABLE IF NOT EXISTS `t_trpc_demo_user_badge` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
    `user_id` varchar(128) NOT NULL COMMENT '用户 ID',
    `badge_name` varchar(128) NOT NULL COMMENT '徽章名称',
    `create_msec` bigint NOT NULL COMMENT '创建时间戳, 毫秒',
    PRIMARY KEY (`id`),
    KEY `i_user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户徽章表';
