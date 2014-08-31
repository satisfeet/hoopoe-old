-- MySQL dump 10.15  Distrib 10.0.13-MariaDB, for osx10.10 (x86_64)
--
-- Host: localhost    Database: satisfeet
-- ------------------------------------------------------
-- Server version	10.0.13-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `address`
--

DROP TABLE IF EXISTS `address`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `address` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `street` varchar(45) DEFAULT '',
  `code` int(5) DEFAULT NULL,
  `city_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `address-city` (`city_id`),
  CONSTRAINT `address-city` FOREIGN KEY (`city_id`) REFERENCES `city` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `address`
--

LOCK TABLES `address` WRITE;
/*!40000 ALTER TABLE `address` DISABLE KEYS */;
INSERT INTO `address` VALUES (1,'Geiserichstraße 3',12105,1),(2,'Hohlstraße 8',35781,2),(3,'Alboinplatz 1',12105,1),(4,'Hedemannstraße 21',10969,1),(5,'Homeyerstraße 24',13156,1),(6,'Wartburgstraße 18',10825,1),(7,'Gieshüglerstraße 46',97218,3),(8,'Hammerstreinstraße 5',14199,1),(9,'Tschaikowskistraße 13',13156,1),(10,'Forster Straße 3',10999,1),(11,'Goethestraße 3a',76547,4),(12,'Krähwinkeler Weg 38',42799,5),(13,'Kirchheimerstraße 46',97271,6),(14,'Tempelhofer Damm 230',12099,1),(15,'Elbestraße 1',12045,1),(16,NULL,NULL,1);
/*!40000 ALTER TABLE `address` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (2,'postage'),(1,'textile');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `city`
--

DROP TABLE IF EXISTS `city`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `city` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `city`
--

LOCK TABLES `city` WRITE;
/*!40000 ALTER TABLE `city` DISABLE KEYS */;
INSERT INTO `city` VALUES (1,'Berlin'),(3,'Gerbrunn'),(6,'Kleinrinderfeld'),(5,'Leichlingen'),(4,'Sinzheim'),(2,'Weilburg');
/*!40000 ALTER TABLE `city` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `color`
--

DROP TABLE IF EXISTS `color`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `color` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `color`
--

LOCK TABLES `color` WRITE;
/*!40000 ALTER TABLE `color` DISABLE KEYS */;
INSERT INTO `color` VALUES (6,'beige'),(2,'dunkelblau'),(3,'grau'),(4,'olive'),(1,'schwarz'),(5,'schwarz-grau');
/*!40000 ALTER TABLE `color` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `customer`
--

DROP TABLE IF EXISTS `customer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `customer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(45) NOT NULL,
  `email` varchar(50) DEFAULT NULL,
  `address_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`),
  KEY `address_id` (`address_id`),
  CONSTRAINT `customer-address` FOREIGN KEY (`address_id`) REFERENCES `address` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `customer`
--

LOCK TABLES `customer` WRITE;
/*!40000 ALTER TABLE `customer` DISABLE KEYS */;
INSERT INTO `customer` VALUES (1,'Bodo Kaiser','i@bodokaiser.io',1),(2,'Wolfgang Kaiser','wmkaiser2@t-online.de',2),(3,'Maximillian Krautwurm','max.krautwurm@gmx.de',3),(4,'Haci Erdal','haci-59@hotmail.de',4),(5,'Urte Cassens',NULL,5),(6,'Torsten Schlingelhof','info@schlingelhof.com',6),(7,'Sabine Kaiser','longvitykaiser@gmail.com',1),(8,'Beatrice Kaiser',NULL,1),(9,'Pascal Schlunek','st.pauliftw@web.de',7),(10,'Burkhard Schneider','ba-schneider@t-online.de',8),(11,'Andreas Schmidt','andreas.schmidt@private-asset.eu',9),(12,'Darius Hajiani','darius@hangload.com',10),(13,'Till Schiewer','till@schiewer.de',16),(14,'Veselina Petkova','vp@zweieinsdrei.de',16),(15,'Pascal Oser','strikedraven@yahoo.de',11),(16,'Gerald Schmidthaus','gerald@das-schmidthaus.de',12),(17,'Fabian Lang','fabi.lang@live.de',16),(18,'Lydia Henneberger',NULL,13),(19,'Jasmin Lewandowski','jasmin.lew@gmail.com',16),(20,'Christopher Schmidt','omm3@gmx.de',14),(21,'Christian Nitschke','hallo@urbix-berlin.de',15);
/*!40000 ALTER TABLE `customer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary table structure for view `customer_address_city`
--

DROP TABLE IF EXISTS `customer_address_city`;
/*!50001 DROP VIEW IF EXISTS `customer_address_city`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE TABLE `customer_address_city` (
  `id` tinyint NOT NULL,
  `name` tinyint NOT NULL,
  `email` tinyint NOT NULL,
  `street` tinyint NOT NULL,
  `code` tinyint NOT NULL,
  `city` tinyint NOT NULL
) ENGINE=MyISAM */;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `discount`
--

DROP TABLE IF EXISTS `discount`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `discount` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL DEFAULT '',
  `description` varchar(120) NOT NULL DEFAULT '',
  `rate` int(2) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `discount`
--

LOCK TABLES `discount` WRITE;
/*!40000 ALTER TABLE `discount` DISABLE KEYS */;
INSERT INTO `discount` VALUES (1,'Geschenk','Geschenk für besondere Kunden und Anlässe.',100);
/*!40000 ALTER TABLE `discount` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment`
--

DROP TABLE IF EXISTS `payment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `customer_id` int(11) NOT NULL,
  `created` datetime DEFAULT NULL,
  `shipped` datetime DEFAULT NULL,
  `cleared` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `customer_id` (`customer_id`),
  CONSTRAINT `payment-customer` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment`
--

LOCK TABLES `payment` WRITE;
/*!40000 ALTER TABLE `payment` DISABLE KEYS */;
INSERT INTO `payment` VALUES (1,1,'2013-12-11 00:00:00','2013-12-11 00:00:00','2013-12-11 00:00:00'),(2,3,'2013-12-11 00:00:00','2013-12-11 00:00:00','2013-12-11 00:00:00'),(3,2,'2013-12-12 00:00:00','2013-12-12 00:00:00','2013-12-12 00:00:00'),(4,4,'2013-12-20 00:00:00','2013-12-12 00:00:00','2013-12-12 00:00:00'),(5,5,'2013-12-20 00:00:00','2013-12-20 00:00:00','2013-12-20 00:00:00'),(6,6,'2013-12-20 00:00:00','2013-12-20 00:00:00','2013-12-20 00:00:00'),(7,7,'2013-12-24 00:00:00','2013-12-24 00:00:00','2013-12-24 00:00:00'),(8,8,'2013-12-24 00:00:00','2013-12-24 00:00:00','2013-12-24 00:00:00'),(9,9,'2014-01-09 00:00:00','2014-01-09 00:00:00','2014-01-09 00:00:00'),(10,7,'2014-02-01 00:00:00','2014-02-01 00:00:00','2014-02-01 00:00:00'),(11,3,'2014-02-08 00:00:00','2014-02-08 00:00:00','2014-02-08 00:00:00'),(12,10,'2014-02-10 00:00:00','2014-02-10 00:00:00','2014-02-10 00:00:00'),(13,11,'2014-02-10 00:00:00','2014-02-10 00:00:00','2014-02-10 00:00:00'),(14,12,'2014-02-10 00:00:00','2014-02-10 00:00:00','2014-02-10 00:00:00'),(15,4,'2014-02-13 00:00:00','2014-02-13 00:00:00','2014-02-13 00:00:00'),(16,13,'2014-02-24 00:00:00','2014-02-24 00:00:00','2014-02-24 00:00:00'),(17,14,'2014-02-27 00:00:00','2014-02-27 00:00:00','2014-02-27 00:00:00'),(18,1,'2014-02-28 00:00:00','2014-02-28 00:00:00','2014-02-28 00:00:00'),(19,15,'2014-03-01 00:00:00','2014-03-01 00:00:00','2014-03-01 00:00:00'),(20,16,'2014-03-04 00:00:00','2014-03-04 00:00:00','2014-03-04 00:00:00'),(21,17,'2014-03-11 00:00:00','2014-03-11 00:00:00','2014-03-11 00:00:00'),(22,18,'2014-03-11 00:00:00','2014-03-11 00:00:00','2014-03-11 00:00:00'),(23,19,'2014-03-18 00:00:00','2014-03-18 00:00:00','2014-03-18 00:00:00'),(24,20,'2014-04-05 00:00:00','2014-04-05 00:00:00','2014-04-05 00:00:00'),(25,12,'2014-04-20 00:00:00','2014-04-20 00:00:00','2014-04-20 00:00:00'),(26,21,'2014-04-20 00:00:00','2014-04-20 00:00:00','2014-04-20 00:00:00'),(27,1,'2014-05-30 00:00:00','2014-05-30 00:00:00','2014-05-30 00:00:00');
/*!40000 ALTER TABLE `payment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_variation`
--

DROP TABLE IF EXISTS `payment_variation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment_variation` (
  `payment_id` int(11) NOT NULL,
  `variation_id` int(11) NOT NULL,
  `discount_id` int(11) DEFAULT NULL,
  `quantity` int(3) NOT NULL,
  PRIMARY KEY (`payment_id`,`variation_id`),
  KEY `discount_id` (`discount_id`),
  KEY `variation_id` (`variation_id`),
  CONSTRAINT `payment_product-discount` FOREIGN KEY (`discount_id`) REFERENCES `discount` (`id`),
  CONSTRAINT `payment_product-payment` FOREIGN KEY (`payment_id`) REFERENCES `payment` (`id`),
  CONSTRAINT `payment_product-variation` FOREIGN KEY (`variation_id`) REFERENCES `variation` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payment_variation`
--

LOCK TABLES `payment_variation` WRITE;
/*!40000 ALTER TABLE `payment_variation` DISABLE KEYS */;
INSERT INTO `payment_variation` VALUES (1,1,1,1),(1,2,1,1),(1,5,1,1),(2,6,1,1),(3,2,1,1),(3,7,1,1),(3,8,1,1),(4,3,1,1),(5,2,NULL,1),(6,1,NULL,1),(7,1,1,1),(8,1,1,1),(9,2,1,1),(10,1,1,1),(11,2,NULL,1),(11,3,NULL,1),(11,4,NULL,1),(12,2,1,1),(12,4,1,1),(13,1,1,1),(13,3,1,1),(14,3,1,1),(15,1,NULL,1),(15,2,NULL,2),(16,1,1,1),(17,1,NULL,1),(18,4,1,1),(19,1,NULL,1),(19,2,NULL,1),(19,4,NULL,1),(20,1,NULL,1),(20,2,NULL,1),(20,3,NULL,1),(20,4,NULL,1),(21,3,NULL,1),(22,2,NULL,10),(23,3,NULL,1),(24,1,1,1),(24,2,1,1),(24,3,1,1),(24,4,1,1),(25,1,1,1),(25,2,NULL,4),(25,3,NULL,3),(25,4,NULL,3),(26,3,NULL,10),(27,4,1,1);
/*!40000 ALTER TABLE `payment_variation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product`
--

DROP TABLE IF EXISTS `product`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(30) NOT NULL DEFAULT '',
  `subtitle` varchar(90) NOT NULL DEFAULT '',
  `description` text NOT NULL,
  `price` decimal(19,4) NOT NULL DEFAULT '0.0000',
  `active` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

LOCK TABLES `product` WRITE;
/*!40000 ALTER TABLE `product` DISABLE KEYS */;
INSERT INTO `product` VALUES (1,'Überlebenssocken','Die perfekte Socken für das Überleben in der Wildnis','Bla Bla\n',5.9900,1),(2,'Alltagssocken','Die perfekte Socken für den normalen Alltag und darüber hinaus','Blab Bla\n',2.9900,1),(3,'Tennissocken','','',0.0000,0),(4,'Wintersocken','','',0.0000,0),(5,'Businesssocken','','',0.0000,0),(6,'Wollsocke','','',0.0000,0);
/*!40000 ALTER TABLE `product` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `product_category`
--

DROP TABLE IF EXISTS `product_category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `product_category` (
  `product_id` int(11) NOT NULL,
  `category_id` int(11) NOT NULL,
  PRIMARY KEY (`product_id`,`category_id`),
  KEY `category_id` (`category_id`),
  KEY `product_id` (`product_id`),
  CONSTRAINT `product_category-product` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`),
  CONSTRAINT `product_category-variation` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product_category`
--

LOCK TABLES `product_category` WRITE;
/*!40000 ALTER TABLE `product_category` DISABLE KEYS */;
INSERT INTO `product_category` VALUES (1,1),(2,1),(3,1),(4,1),(5,1),(6,1);
/*!40000 ALTER TABLE `product_category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary table structure for view `product_variation_category`
--

DROP TABLE IF EXISTS `product_variation_category`;
/*!50001 DROP VIEW IF EXISTS `product_variation_category`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE TABLE `product_variation_category` (
  `id` tinyint NOT NULL,
  `title` tinyint NOT NULL,
  `subtitle` tinyint NOT NULL,
  `description` tinyint NOT NULL,
  `price` tinyint NOT NULL,
  `variations` tinyint NOT NULL,
  `categories` tinyint NOT NULL
) ENGINE=MyISAM */;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `size`
--

DROP TABLE IF EXISTS `size`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `size` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(5) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `size`
--

LOCK TABLES `size` WRITE;
/*!40000 ALTER TABLE `size` DISABLE KEYS */;
INSERT INTO `size` VALUES (1,'42-44');
/*!40000 ALTER TABLE `size` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `variation`
--

DROP TABLE IF EXISTS `variation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `variation` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) NOT NULL,
  `color_id` int(11) NOT NULL,
  `size_id` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `color_id` (`color_id`),
  KEY `size_id` (`size_id`),
  KEY `product_id` (`product_id`),
  CONSTRAINT `variation-color` FOREIGN KEY (`color_id`) REFERENCES `color` (`id`),
  CONSTRAINT `variation-product` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`),
  CONSTRAINT `variation-size` FOREIGN KEY (`size_id`) REFERENCES `size` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `variation`
--

LOCK TABLES `variation` WRITE;
/*!40000 ALTER TABLE `variation` DISABLE KEYS */;
INSERT INTO `variation` VALUES (1,1,4,1),(2,2,1,1),(3,2,2,1),(4,2,3,1),(5,3,1,1),(6,4,5,1),(7,5,1,1),(8,6,6,1);
/*!40000 ALTER TABLE `variation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Final view structure for view `customer_address_city`
--

/*!50001 DROP TABLE IF EXISTS `customer_address_city`*/;
/*!50001 DROP VIEW IF EXISTS `customer_address_city`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `customer_address_city` AS select `cu`.`id` AS `id`,`cu`.`name` AS `name`,`cu`.`email` AS `email`,`ad`.`street` AS `street`,`ad`.`code` AS `code`,`ci`.`name` AS `city` from ((`customer` `cu` left join `address` `ad` on((`cu`.`address_id` = `ad`.`id`))) left join `city` `ci` on((`ad`.`city_id` = `ci`.`id`))) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `product_variation_category`
--

/*!50001 DROP TABLE IF EXISTS `product_variation_category`*/;
/*!50001 DROP VIEW IF EXISTS `product_variation_category`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `product_variation_category` AS select `pr`.`id` AS `id`,`pr`.`title` AS `title`,`pr`.`subtitle` AS `subtitle`,`pr`.`description` AS `description`,`pr`.`price` AS `price`,group_concat(concat(`co`.`name`,':',`si`.`name`) separator ',') AS `variations`,group_concat(distinct `ca`.`name` separator ',') AS `categories` from (((((`product` `pr` left join `product_category` `pc` on((`pc`.`product_id` = `pr`.`id`))) left join `category` `ca` on((`pc`.`category_id` = `ca`.`id`))) left join `variation` `va` on((`va`.`product_id` = `pr`.`id`))) left join `color` `co` on((`va`.`color_id` = `co`.`id`))) left join `size` `si` on((`va`.`size_id` = `si`.`id`))) group by `pr`.`id` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2014-08-31  8:36:14
