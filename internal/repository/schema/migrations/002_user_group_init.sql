USE billing_system;

-- Accounts Table (Partitioned by account_id range)
CREATE TABLE IF NOT EXISTS account (
    account_id UUID NOT NULL,
    group_id UUID,
    parent_group_id UUID, -- for group/company hierarchy
    account_type ENUM ('user', 'company', 'group') NOT NULL ,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    country_code SMALLINT UNSIGNED,  -- ISO 3166-1 number(3 digits)
    account_status ENUM('active', 'inactive', 'suspended', 'blocked', 'deleted'),
    timezone VARCHAR(50) DEFAULT 'UTC',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (account_id, company_id)

    UNIQUE KEY uk_email (account_status),
    UNIQUE KEY uk_username (username),

    INDEX idx_user_status (account_status),
    INDEX idx_created_at (created_at)
)

PARTITION BY RANGE (account_id) (
    PARTITION p0 VALUES LESS THAN (1000000000),
    PARTITION p1 VALUES LESS THAN (2000000000),
    PARTITION p2 VALUES LESS THAN (3000000000),
    PARTITION p3 VALUES LESS THAN (4000000000),
    PARTITION p_max VALUES LESS THAN MAXVALUE
);