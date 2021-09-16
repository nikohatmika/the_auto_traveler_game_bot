CREATE DATABASE the_auto_traveler

USE the_auto_traveler

CREATE TABLE equipments(
    id serial PRIMARY KEY,
    created_at timestamp,
    updated_at timestamp,
    type varchar(8),
    name varchar(32),
    description varchar(255),
    atk int,
    def int,
    hp int
)

CREATE TABLE events(
    id serial PRIMARY KEY,
    created_at timestamp,
    updated_at timestamp,
    type varchar(8),
    name varchar(32),
    description varchar(255),
    gold_reward int,
    xp_reward int
)

CREATE TABLE users(
    id serial PRIMARY KEY,
    created_at timestamp,
    updated_at timestamp,
    name varchar(32),
    level int,
    xp int,
    gold int,
    fame_points int,
    spare_stat_points int,
    atk int,
    def int,
    hp int,
    weapon_equipment_id bigint UNSIGNED,
    armor_equipment_id bigint UNSIGNED,
	FOREIGN KEY (weapon_equipment_id) REFERENCES equipments(id),
    FOREIGN KEY (armor_equipment_id) REFERENCES equipments(id)
)

CREATE TABLE user_friends(
    user_id int NOT NULL,
    friend_id int NOT NULL,
    created_at timestamp,
    PRIMARY KEY(user_id, friend_id)
)

CREATE TABLE user_equipments(
    user_id int NOT NULL,
    equipment_id int NOT NULL,
    created_at timestamp,
    PRIMARY KEY(user_id, equipment_id)
)