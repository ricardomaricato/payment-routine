CREATE DATABASE IF NOT EXISTS accounts;
USE accounts;

CREATE TABLE accounts(
    account_id int auto_increment primary key,
    available_credit_limit decimal(10, 2) not null,
    available_with_drawal_limit decimal(10, 2) not null
) ENGINE=INNODB;