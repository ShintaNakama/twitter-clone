
-- -----------------------------------------------------
-- Table `posts`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `posts`;

CREATE TABLE IF NOT EXISTS `posts` (
  `id` VARCHAR(26) NOT NULL,
  `user_id` VARCHAR(26) NOT NULL COMMENT 'ユーザーID',
  `body` TEXT NOT NULL COMMENT '投稿本文',
  `posted_at` DATETIME NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `index_posts_on_user_id` (`user_id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COMMENT="投稿文情報";

-- -----------------------------------------------------
-- Table `users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `users`;

CREATE TABLE IF NOT EXISTS `users` (
  `id` VARCHAR(26) NOT NULL,
  `email` VARCHAR(255) UNIQUE NOT NULL COMMENT 'メールアドレス',
  `name` VARCHAR(255) NOT NULl COMMENT 'ユーザー名',
  `image` VARCHAR(255) COMMENT '画像URL',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COMMENT = 'ユーザー';

-- -----------------------------------------------------
-- Table `profiles`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `profiles`;

CREATE TABLE IF NOT EXISTS `profiles` (
  `user_id` VARCHAR(26) NOT NULL COMMENT 'ユーザーID',
  `description` TEXT COMMENT '自己紹介',
  `birthday` DATE COMMENT '生年月日',
  `location` VARCHAR(255) COMMENT '居住地',
  `gender` TINYINT COMMENT '性別',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  CONSTRAINT `fk_profile_user`
    FOREIGN KEY (`user_id`)
    REFERENCES `users` (`id`)
    ON DELETE CASCADE
    ON UPDATE NO ACTION
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COMMENT = 'プロフィール';

-- -----------------------------------------------------
-- Table `images`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `images` (
  `id` VARCHAR(26) NOT NULL,
  `name` VARCHAR(255) NOT NULL COMMENT '投稿画像ファイル名',
  `display_order` TINYINT(4) NOT NULL DEFAULT '-1' COMMENT '投稿画像表示順',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COMMENT="投稿画像情報";

-- -----------------------------------------------------
-- Table `post_images`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `post_images` (
  `post_id` VARCHAR(26) NOT NULL COMMENT '投稿ID 外部キー',
  `image_id` VARCHAR(26) NOT NULL COMMENT '投稿画像ID 外部キー',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`post_id`, `image_id`),
  CONSTRAINT `post_images_ibfk_1`
      FOREIGN KEY(`post_id`) REFERENCES `posts` (`id`)
      ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT `post_images_ibfk_2`
      FOREIGN KEY(`image_id`) REFERENCES `images` (`id`)
      ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE = InnoDB DEFAULT CHARACTER SET = utf8mb4 COMMENT="投稿と投稿画像の中間テーブル";
