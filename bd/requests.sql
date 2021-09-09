CREATE DATABASE IF NOT EXISTS requestsdb;
USE requestsdb;
DROP TABLE IF EXISTS requests;
CREATE TABLE requests(id INT PRIMARY KEY AUTO_INCREMENT, city VARCHAR(255), time_requested VARCHAR(255), temperaute VARCHAR(255));
INSERT INTO requests(city, time_requested, temperaute) VALUES('Abu Dhabi','21:08:08-26:08:2021', '35');