CREATE TABLE IF NOT EXISTS `t_user_account` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
    `username` varchar(128) NOT NULL COMMENT '用户名称',
    `password_hash` varchar(64) NOT NULL COMMENT '用户密码哈希值',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_at_ms` bigint(11) NOT NULL DEFAULT 0 COMMENT '删除时间戳, 毫秒',
    PRIMARY KEY (`id`),
    KEY `i_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户账户表';
