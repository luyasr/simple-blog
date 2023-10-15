CREATE TABLE `users`
(
    `id`        int                                                           NOT NULL AUTO_INCREMENT,
    `username`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
    `password`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
    `create_at` int                                                           NOT NULL COMMENT '创建时间',
    `update_at` int                                                           NOT NULL COMMENT '更新时间',
    `delete_at` int                                                           NOT NULL COMMENT '删除时间',
    `role`      tinyint                                                       NOT NULL COMMENT '角色',
    PRIMARY KEY (`id`),
    KEY         `idx_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci

CREATE TABLE `tokens`
(
    `created_at`               int                                                           NOT NULL COMMENT '创建时间',
    `updated_at`               int                                                           NOT NULL COMMENT '更新时间',
    `user_id`                  int                                                           NOT NULL COMMENT '用户的Id',
    `username`                 varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名, 用户名不允许重复的',
    `access_token`             varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户的访问令牌',
    `access_token_expired_at`  int                                                           NOT NULL COMMENT '令牌过期时间',
    `refresh_token`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '刷新令牌',
    `refresh_token_expired_at` int                                                           NOT NULL COMMENT '刷新令牌过期时间',
    PRIMARY KEY (`access_token`) USING BTREE,
    UNIQUE KEY `idx_token` (`access_token`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci

CREATE TABLE `blogs`
(
    `id`           int                                                           NOT NULL AUTO_INCREMENT,
    `create_at`    int                                                           NOT NULL COMMENT '创建时间',
    `update_at`    int                                                           NOT NULL COMMENT '更新时间',
    `status`       tinyint                                                       NOT NULL COMMENT '状态',
    `title`        varchar(255)                                                  NOT NULL COMMENT '标题',
    `author`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '作者',
    `create_by`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '创建的用户',
    `summary`      varchar(255)                                                  NOT NULL COMMENT '概要',
    `content`      varchar(255)                                                  NOT NULL COMMENT '内容',
    `audit_at`     int                                                           NOT NULL COMMENT '审核时间',
    `audit_status` tinyint                                                       NOT NULL COMMENT '审核状态',
    `tags`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci