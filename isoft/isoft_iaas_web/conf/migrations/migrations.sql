CREATE TABLE `blog` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `author` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `blog_title` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `key_words` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `catalog_id` bigint(20) NOT NULL,
  `content` longtext COLLATE utf8_bin NOT NULL,
  `blog_type` tinyint(4) NOT NULL DEFAULT '0',
  `blog_status` tinyint(4) NOT NULL DEFAULT '0',
  `views` bigint(20) NOT NULL DEFAULT '0',
  `edits` bigint(20) NOT NULL DEFAULT '0',
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  `short_desc` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `catalog` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `author` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `catalog_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `catalog_desc` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `comment_reply` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `comment_theme_id` int(11) NOT NULL,
  `depth` int(11) NOT NULL DEFAULT '0',
  `reply_content` varchar(4000) COLLATE utf8_bin NOT NULL DEFAULT '',
  `refer_user_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `sub_reply_amount` int(11) NOT NULL DEFAULT '0',
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  `reply_theme_type` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `reply_comment_type` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `comment_theme` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `comment_id` int(11) NOT NULL DEFAULT '0',
  `comment_content` varchar(4000) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  `theme_type` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `common_link` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `link_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `link_addr` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  `link_type` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `configuration` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) NOT NULL DEFAULT '0',
  `configuration_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `configuration_value` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  `status` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `course` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `course_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `course_author` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `course_type` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `course_sub_type` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `course_short_des` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `small_image` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `score` double NOT NULL DEFAULT '0',
  `course_number` int(11) NOT NULL DEFAULT '0',
  `course_status` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `media_type` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `watch_number` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `course_vedio` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `course_id` int(11) NOT NULL,
  `vedio_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `vedio_number` int(11) NOT NULL DEFAULT '0',
  `first_play` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `second_play` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `course_video` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `course_id` int(11) NOT NULL,
  `video_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `video_number` int(11) NOT NULL DEFAULT '0',
  `first_play` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `second_play` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `favorite` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `favorite_id` int(11) NOT NULL DEFAULT '0',
  `favorite_type` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `user_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `heart_beat2` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `addr` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `status_code` bigint(20) NOT NULL DEFAULT '0',
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `heart_beat_detail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `addr` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `status_code` int(11) NOT NULL DEFAULT '0',
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13926 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `history` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `history_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `history_value` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  `history_desc` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `history_link` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `i_file` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `fid` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `file_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `file_size` bigint(20) NOT NULL DEFAULT '0',
  `url` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `note` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `note_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `note_owner` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `note_key_words` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `note_content` longtext COLLATE utf8_bin NOT NULL,
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  `edit_time` int(11) NOT NULL DEFAULT '0',
  `view_time` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `share` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `share_type` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `author` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `link_href` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  `share_desc` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `content` longtext COLLATE utf8_bin NOT NULL,
  `views` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

CREATE TABLE `topic_reply` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `topic_theme_id` int(11) NOT NULL,
  `reply_theme_type` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `reply_content` varchar(4000) COLLATE utf8_bin NOT NULL DEFAULT '',
  `refer_user_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `sub_reply_amount` int(11) NOT NULL DEFAULT '0',
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


CREATE TABLE `topic_theme` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `topic_id` int(11) NOT NULL DEFAULT '0',
  `topic_type` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `topic_content` varchar(4000) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT '',
  `last_updated_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

