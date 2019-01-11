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