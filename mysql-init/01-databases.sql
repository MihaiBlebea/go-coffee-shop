CREATE DATABASE IF NOT EXISTS `account`;
CREATE DATABASE IF NOT EXISTS `order`;
CREATE DATABASE IF NOT EXISTS `bonus`;


CREATE TABLE IF NOT EXISTS account.users (
    id VARCHAR(255) PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    age INT(10) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order.payments (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    coffee_id VARCHAR(255) NOT NULL,
    amount INT(10) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS bonus.stamps (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE USER 'mihaiaccount' IDENTIFIED BY 'pass';
GRANT ALL privileges ON `account`.* TO 'mihaiaccount'@'%' IDENTIFIED BY 'pass';

CREATE USER 'mihaiorder' IDENTIFIED BY 'pass';
GRANT ALL privileges ON `order`.* TO 'mihaiorder'@'%' IDENTIFIED BY 'pass';

CREATE USER 'mihaibonus' IDENTIFIED BY 'pass';
GRANT ALL privileges ON `bonus`.* TO 'mihaibonus'@'%' IDENTIFIED BY 'pass';

CREATE USER 'read' IDENTIFIED BY 'read';
GRANT SELECT ON *.* TO 'read'@'%' IDENTIFIED BY 'read';