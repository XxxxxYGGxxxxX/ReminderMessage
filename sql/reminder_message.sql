/*
SQLyog Community v13.1.6 (64 bit)
MySQL - 8.0.26 
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `reminder_message` (
	`id` int (11),
	`creator_id` varchar (765),
	`content` varchar (765),
	`remind_time` datetime 
); 
insert into `reminder_message` (`id`, `creator_id`, `content`, `remind_time`) values('22','123','test111','2024-05-23 23:47:18');
insert into `reminder_message` (`id`, `creator_id`, `content`, `remind_time`) values('23','123','测试更新数据','2024-05-23 23:47:14');
insert into `reminder_message` (`id`, `creator_id`, `content`, `remind_time`) values('24','123','222333','2024-05-23 23:47:51');
insert into `reminder_message` (`id`, `creator_id`, `content`, `remind_time`) values('25','','','0000-00-00 00:00:00');
