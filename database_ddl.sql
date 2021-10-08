CREATE DATABASE `mtga` /*!40100 DEFAULT CHARACTER SET latin1 */;

-- mtga.decks definition

CREATE TABLE `decks` (
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
  `numenchant` int(11) DEFAULT 0,
  `numartifacts` int(11) DEFAULT 0,
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
  `game_type` varchar(100) DEFAULT NULL,
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
  `numenchant` int(11) DEFAULT 0,
  `numartifacts` int(11) DEFAULT 0,
  PRIMARY KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- mtga.cards definition

CREATE TABLE `mtga`.`cards` (
  `deck` varchar(100) DEFAULT NULL,
  `numcopy` int(11) DEFAULT NULL,
  `cardname` varchar(100) DEFAULT NULL,
  `set` varchar(100) DEFAULT NULL,
  `setnum` int(11) DEFAULT NULL,
  `side_board` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- mtga.set_abbreviations definition

CREATE TABLE `mtga`.`set_abbreviations` (
  `set_name` varchar(100) DEFAULT NULL,
  `set_abbrev` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- mtga.`sets` definition

CREATE TABLE `sets` (
  `set_name` varchar(100) NOT NULL,
  `card_name` varchar(1000) DEFAULT NULL,
  `colors` varchar(100) DEFAULT NULL,
  `mana_cost` decimal(10,0) DEFAULT NULL,
  `mana_colors` varchar(100) DEFAULT NULL,
  `converted_mana_cost` decimal(10,0) DEFAULT NULL,
  `set_number` varchar(100) DEFAULT NULL,
  `card_text` mediumtext DEFAULT NULL,
  `type` varchar(100) DEFAULT NULL,
  `sub_type` varchar(100) DEFAULT NULL,
  `super_type` varchar(100) DEFAULT NULL,
  `types` varchar(100) DEFAULT NULL,
  `rarity` varchar(100) DEFAULT NULL,
  `set_code` varchar(100) DEFAULT NULL,
  `card_side` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

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
    (`r`.`wins` + 1) / (2 + sum(`r`.`wins` + `r`.`loses`)) as `ranking`,
    `r`.`wins` as `wins`,
    `r`.`loses` as `loses`
from
    `mtga`.`record` `r`
group by
    `r`.`deck`
order by
    (`r`.`wins` + 1) / (2 + sum(`r`.`wins` + `r`.`loses`)) desc,
    `r`.`wins` desc,
    `r`.`loses`
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

-- mtga.loses_by_day source

create or replace
algorithm = UNDEFINED view `mtga`.`loses_by_day` as
select
    `g`.`deck` as `deck`,
    dayname(`g`.`Timestamp`) as `day_of_week`,
    count(`g`.`results`) as `lose_count`
from
    `mtga`.`games` `g`
where
    `g`.`results` = 1
group by
    `g`.`deck`,
    dayname(`g`.`Timestamp`);
	
-- mtga.wins_by_day source

create or replace
algorithm = UNDEFINED view `mtga`.`wins_by_day` as
select
    `g`.`deck` as `deck`,
    dayname(`g`.`Timestamp`) as `day_of_week`,
    count(`g`.`results`) as `win_count`
from
    `mtga`.`games` `g`
where
    `g`.`results` = 0
group by
    `g`.`deck`,
    dayname(`g`.`Timestamp`);

-- mtga.most_wbd source

create or replace
algorithm = UNDEFINED view `mtga`.`most_wbd` as with added_row_number as (
select
    `wins_by_day`.`deck` as `deck`,
    `wins_by_day`.`day_of_week` as `day_of_week`,
    `wins_by_day`.`win_count` as `win_count`,
    row_number() over ( partition by `wins_by_day`.`deck`
order by
    `wins_by_day`.`win_count` desc) as `row_number`
from
    `mtga`.`wins_by_day`
)select
    `added_row_number`.`deck` as `deck`,
    `added_row_number`.`day_of_week` as `day_of_week`,
    `added_row_number`.`win_count` as `win_count`
from
    `added_row_number`
where
    `added_row_number`.`row_number` = 1;
	
-- mtga.most_lbd source

create or replace
algorithm = UNDEFINED view `mtga`.`most_lbd` as with added_row_number as (
select
    `loses_by_day`.`deck` as `deck`,
    `loses_by_day`.`day_of_week` as `day_of_week`,
    `loses_by_day`.`lose_count` as `lose_count`,
    row_number() over ( partition by `loses_by_day`.`deck`
order by
    `loses_by_day`.`lose_count` desc) as `row_number`
from
    `mtga`.`loses_by_day`
)select
    `added_row_number`.`deck` as `deck`,
    `added_row_number`.`day_of_week` as `day_of_week`,
    `added_row_number`.`lose_count` as `lose_count`
from
    `added_row_number`
where
    `added_row_number`.`row_number` = 1;