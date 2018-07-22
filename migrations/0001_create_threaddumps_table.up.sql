CREATE TABLE `threaddumps` (
    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, 
    `application` varchar(255), 
    `host` varchar(255), 
    `upload_time` timestamp
);