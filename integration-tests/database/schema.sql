CREATE TABLE profiles (
    id INT NOT NULL AUTO_INCREMENT,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,

    uuid VARCHAR(50) UNIQUE NOT NULL,
    gender VARCHAR(50) NOT NULL,
    gender_preference VARCHAR(50) NOT NULL,
    age int NOT NULL,
    min_age_preference int NOT NULL,
    max_age_preference int NOT NULL,
    profile_lat double NOT NULL,
    profile_lng double NOT NULL,
    PRIMARY KEY (id)
) ENGINE=INNODB;

CREATE TABLE tracked_questions (
    id INT NOT NULL AUTO_INCREMENT,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,

    uuid VARCHAR(50) UNIQUE NOT NULL,
    questionText TEXT NOT NULL,
    questionIndex int NOT NULL
    category TEXT NOT NULL,
    user_uuid  VARCHAR(50) UNIQUE NOT NULL,
    question_uuid VARCHAR(50) UNIQUE NOT NULL,
    PRIMARY KEY (id)
) ENGINE=INNODB;

CREATE TABLE tracked_likes (
    id INT NOT NULL AUTO_INCREMENT,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,

    uuid VARCHAR(50) UNIQUE NOT NULL,
    user_uuid VARCHAR(50) UNIQUE NOT NULL,
    target_uuid VARCHAR(50) UNIQUE NOT NULL,
    liked boolean UNIQUE NOT NULL,
    PRIMARY KEY (id)
) ENGINE=INNODB;

CREATE TABLE matches (
    id INT NOT NULL AUTO_INCREMENT,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,

    uuid VARCHAR(50) UNIQUE NOT NULL,
    room_uuid VARCHAR(50) UNIQUE NOT NULL,
    PRIMARY KEY (id)
) ENGINE=INNODB;

