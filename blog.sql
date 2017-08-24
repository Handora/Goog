CREATE DATABASE blog;
USE blog;

-- create the Article table corresponding to Article structure
CREATE TABLE `Article` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT NULL,
  `intro` text,
  `content` text,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;

-- create the Tag table corresponding to Tag structure
CREATE TABLE `Tag` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

-- create the relation table between Article and Tag,
-- with two foreign key with each one point to Article
-- and Tag table's id
CREATE TABLE `ATrelation` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `aid` int(20) NOT NULL,
  `tid` int(20) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_aid` (`aid`),
  KEY `fk_tid` (`tid`),
  CONSTRAINT `fk_aid` FOREIGN KEY (`aid`) REFERENCES `Article` (`id`),
  CONSTRAINT `fk_tid` FOREIGN KEY (`tid`) REFERENCES `Tag` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;