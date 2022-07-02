CREATE TABLE `Plans` (
  `id` bigint PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `content` varchar(255) NOT NULL,
  `state` varchar(255) NOT NULL,
  `plan_start_date_and_hour` timestamp NOT NULL,
  `plan_finish_date_and_hour` timestamp NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT "now()"
);

CREATE TABLE `Users` (
  `id` bigint PRIMARY KEY,
  `name` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL
);
