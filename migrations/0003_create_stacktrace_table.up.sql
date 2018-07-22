CREATE TABLE `stacktrace_lines` (
    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    `kind` varchar(255) NOT NULL,
    `line_number` integer NOT NULL,
    `lock_address` varchar(255),
    `locked_class` varchar(255),
    `lock_class` varchar(255),
    `java_class` varchar(255), 
    `java_method` varchar(255), 
    `source_line` varchar(255), 
    `source_file` varchar(255), 
    `java_thread_id` integer REFERENCES `java_threads`
);