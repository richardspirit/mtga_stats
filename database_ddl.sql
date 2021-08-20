CREATE DATABASE `mgta` /*!40100 DEFAULT CHARACTER SET latin1 */;

-- mgta.decks definition

CREATE TABLE `decks` (
  `name` varchar(100) NOT NULL,
  `colors` varchar(100) DEFAULT NULL,
  `date_entered` date NOT NULL DEFAULT curdate(),
  `favorite` tinyint(1) NOT NULL DEFAULT 1,
  `max_streak` int(11) DEFAULT NULL,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- mgta.games definition

CREATE TABLE `games` (
  `UID` bigint(20) NOT NULL DEFAULT uuid_short(),
  `Timestamp` timestamp NOT NULL DEFAULT current_timestamp(),
  `results` binary(1) DEFAULT NULL,
  `cause` varchar(100) DEFAULT NULL,
  `deck` varchar(100) NOT NULL,
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- mgta.game_count source

create or replace
algorithm = UNDEFINED view `mgta`.`game_count` as
select
    `g`.`UID` as `UID`,
    count(`g`.`results`) as `count(results)`,
    `g`.`deck` as `deck`
from
    `mgta`.`games` `g`
group by
    `g`.`deck`;
	
-- mgta.record source

create or replace
algorithm = UNDEFINED view `mgta`.`record` as
select
    `g`.`UID` as `UID`,
    count(case when `g`.`results` = 0 then 1 end) as `wins`,
    count(case when `g`.`results` = 1 then 1 end) as `loses`,
    `g`.`deck` as `deck`
from
    `mgta`.`games` `g`
group by
    `g`.`deck`;