CREATE TABLE `java_threads` (
    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    `name` varchar(255),
    `java_id` varchar(255),
    `is_daemon` Boolean,
    `prio` integer,
    `os_prio` integer,
    `tid` varchar(255),
    `nid` varchar(255),
    `native_thread_state` varchar(255),
    `condition_address` varchar(255),
    `java_thread_state` varchar(255),
    `java_state_clarification` varchar(255),
    `threaddump_id` integer REFERENCES `threaddumps`
);