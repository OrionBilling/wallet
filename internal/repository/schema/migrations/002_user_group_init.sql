USE billing_system;

-- Users Table (Partitioned by user_id range)
CREATE TABLE IF NOT EXISTS users (
    user_id BIGINT UNSIGNED NOT NULL,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    country_code SMALLINT UNSIGNED,  -- ISO 3166-1 number(3 digits)
    status ENUM('active', 'inactive', 'suspended', 'blocked', 'deleted'),
    current_plan_id INT UNSIGNED,
    plan_start_date DATE,
    plan_end_date DATE,
    auto_renew BOOLEAN DEFAULT TRUE,
    timezone VARCHAR(50) DEFAULT 'UTC',
    metadata JSON,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (user_id),

    UNIQUE KEY uk_email (email),
    UNIQUE KEY uk_username (username),

    INDEX idx_user_status (status),
    INDEX idx_user_plan (current_plan_id),
    INDEX idx_plan_dates (plan_start_date, plan_end_date),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB

PARTITION BY RANGE (user_id) (
    PARTITION p0 VALUES LESS THAN (1000000000),
    PARTITION p1 VALUES LESS THAN (2000000000),
    PARTITION p2 VALUES LESS THAN (3000000000),
    PARTITION p3 VALUES LESS THAN (4000000000),
    PARTITION p_max VALUES LESS THAN MAXVALUE
);