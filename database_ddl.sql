CREATE DATABASE `mtga` /*!40100 DEFAULT CHARACTER SET latin1 */;

-- mtga.decks definition

CREATE TABLE `mtga`.`decks` (
  `name` varchar(100) NOT NULL,
  `colors` varchar(100) DEFAULT NULL,
  `date_entered` date NOT NULL DEFAULT curdate(),
  `favorite` tinyint(1) NOT NULL DEFAULT 1,
  `max_streak` int(11) DEFAULT 0,
  `cur_streak` int(11) DEFAULT 0,
  `numcards` int(11) DEFAULT 0,
  `numlands` int(11) DEFAULT 0,
  `numspells` int(11) DEFAULT 0,
  `numcreatures` int(11) DEFAULT 0,
  `disable` binary(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- mtga.games definition

CREATE TABLE `mtga`.`games` (
  `UID` bigint(20) NOT NULL DEFAULT uuid_short(),
  `Timestamp` timestamp NOT NULL DEFAULT current_timestamp(),
  `results` binary(1) DEFAULT '0',
  `cause` varchar(100) DEFAULT 'Unknown',
  `deck` varchar(100) NOT NULL,
  `opponent` varchar(100) DEFAULT 'Unknown',
  `level` varchar(100) DEFAULT 'Unknown',
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- mtga.decks_deleted definition

CREATE TABLE `mtga`.`decks_deleted` (
  `name` varchar(100) NOT NULL,
  `colors` varchar(100) DEFAULT NULL,
  `date_entered` date NOT NULL DEFAULT curdate(),
  `favorite` tinyint(1) NOT NULL DEFAULT 1,
  `max_streak` int(11) DEFAULT 0,
  `cur_streak` int(11) DEFAULT 0,
  `numcards` int(11) DEFAULT 0,
  `numlands` int(11) DEFAULT 0,
  `numspells` int(11) DEFAULT 0,
  `numcreatures` int(11) DEFAULT 0,
  `disable` binary(1) NOT NULL DEFAULT '1',
  `UID` bigint(20) NOT NULL DEFAULT uuid_short(),
  PRIMARY KEY (`UID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- mtga.game_count source

create or replace
algorithm = UNDEFINED view `mtga`.`game_count` as
select
    count(`g`.`results`) as `results`,
    `g`.`deck` as `deck`
from
    `mtga`.`games` `g`
group by
    `g`.`deck`;
	
-- mtga.record source

create or replace
algorithm = UNDEFINED view `mtga`.`record` as
select
   -- `g`.`UID` as `UID`,
    count(case when `g`.`results` = 0 then 1 end) as `wins`,
    count(case when `g`.`results` = 1 then 1 end) as `loses`,
    `g`.`deck` as `deck`
from
    `mtga`.`games` `g`
group by
    `g`.`deck`;
	
-- mtga.topten source

create or replace
algorithm = UNDEFINED view `mtga`.`topten` as
select
    `r`.`deck` as `deck`,
    `r`.`wins` as `wins`,
    `r`.`loses` as `loses`
from
    `mtga`.`record` `r`
order by
    `r`.`wins` desc,
    `r`.`loses` desc
limit 10;

-- mtga.lose_percentage source

create or replace
algorithm = UNDEFINED view `mtga`.`lose_percentage` as
select
    `g`.`lose_count` / `gc`.`results` as `lose_pct`,
    `gc`.`deck` as `deck`,
    `g`.`lose_count` as `lose_count`,
    `gc`.`results` as `games`
from
    (`mtga`.`game_count` `gc`
join (
    select
        count(`mtga`.`games`.`results`) as `lose_count`,
        `mtga`.`games`.`deck` as `deck`
    from
        `mtga`.`games`
    where
        `mtga`.`games`.`results` = 1
    group by
        `mtga`.`games`.`deck`) `g` on
    (`gc`.`deck` = `g`.`deck`));

-- mtga.win_percentage source

create or replace
algorithm = UNDEFINED view `mtga`.`win_percentage` as
select
    `g`.`win_count` / `gc`.`results` as `win_pct`,
    `gc`.`deck` as `deck`,
    `g`.`win_count` as `win_count`,
    `gc`.`results` as `games`
from
    (`mtga`.`game_count` `gc`
join (
    select
        count(`mtga`.`games`.`results`) as `win_count`,
        `mtga`.`games`.`deck` as `deck`
    from
        `mtga`.`games`
    where
        `mtga`.`games`.`results` = 0
    group by
        `mtga`.`games`.`deck`) `g` on
    (`gc`.`deck` = `g`.`deck`));