/*
Navicat MySQL Data Transfer

Source Server         : conexiones
Source Server Version : 50505
Source Host           : 127.0.0.1:3306
Source Database       : go-api

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2018-09-30 14:13:16
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for recipes
-- ----------------------------
DROP TABLE IF EXISTS `recipes`;
CREATE TABLE `recipes` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `description` text,
  `ingredients` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of recipes
-- ----------------------------
INSERT INTO `recipes` VALUES ('2', 'Modificará', 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsum eius repellat cupiditate odio similique incidunt dolorem rem nesciunt adipisci provident eum, nulla impedit nisi facere! Nisi, natus. Id, officiis ipsam!, Lorem ipsum dolor sit amet, consectetur adipisicing elit. Recusandae aliquam hic asperiores, quo, dolorem molestias dolores similique officia, est accusantium reprehenderit repudiandae culpa expedita ratione rem at quidem facilis veritatis!', 'otro ingredientes');
INSERT INTO `recipes` VALUES ('3', 'Pollo asado', 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsum eius repellat cupiditate odio similique incidunt dolorem rem nesciunt adipisci provident eum, nulla impedit nisi facere! Nisi, natus. Id, officiis ipsam!, Lorem ipsum dolor sit amet, consectetur adipisicing elit. Recusandae aliquam hic asperiores, quo, dolorem molestias dolores similique officia, est accusantiu', 'otro pollo');
INSERT INTO `recipes` VALUES ('4', 'ddd', 'Lorem ipsum dolor sit amet consectetur adipisicing elit. Ipsum eius repellat cupiditate odio similique incidunt dolorem rem nesciunt adipisci provident eum, nulla impedit nisi facere! Nisi, natus. Id, officiis ipsam!, Lorem ipsum dolor sit amet, consectetur adipisicing elit. Recusandae aliquam hic asperiores, quo, dolorem molestias dolores similique officia, est accusantium reprehenderit repudiandae culpa expedita ratione rem at quidem facilis veritatis!', 'ddd');
INSERT INTO `recipes` VALUES ('5', 'Torta de chocolate', 'Otra descripción', '');
INSERT INTO `recipes` VALUES ('6', 'Ja jajaja', 'Otra descripción', '');
INSERT INTO `recipes` VALUES ('7', 'I am', 'Otra descripción', '');
INSERT INTO `recipes` VALUES ('8', 'You are', 'Otra descripción', '');
INSERT INTO `recipes` VALUES ('13', 'ddd', 'ddddddddddddd', '');
INSERT INTO `recipes` VALUES ('14', 'jjjjjjjjjjjjjjj', 'jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj', '');
INSERT INTO `recipes` VALUES ('20', 'Sol y Meme', 'Mis sobrinos yrrrrrrrrrrrrrrrrrr', '');
INSERT INTO `recipes` VALUES ('21', 'Hola', 'Mundo', '');
