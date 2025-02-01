
CREATE TABLE IF NOT EXISTS `t_trpc_demo_user_account` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
    `username` varchar(128) NOT NULL COMMENT '用户名称',
    `password_hash` varchar(64) NOT NULL COMMENT '用户密码哈希值',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_at_ms` bigint(11) NOT NULL DEFAULT 0 COMMENT '删除时间戳, 毫秒',
    `nickname` varchar(128) NOT NULL DEFAULT '' COMMENT '用户昵称',
    PRIMARY KEY (`id`),
    UNIQUE KEY `u_username` (`username`, `delete_at_ms`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户账户表';
