USE billing_system;

-- Billing Plans Table
CREATE TABLE IF NOT EXISTS billing_plan (
    plan_id INT UNSIGNED NOT NULL,
    plan_name VARCHAR(100) NOT NULL,
    plan_description TEXT,
    price BIGINT UNSIGNED NOT NULL, -- in the smallest currency unit
    currency SMALLINT UNSIGNED, -- ISO-4217 (3 digits)
    billing_cycle ENUM('monthly', 'yearly', 'lifetime') NOT,
    features JSON,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (plan_id),

    INDEX idx_plan_active (is_active),
    INDEX idx_plan_cycle (billing_cycle)
) ENGINE=InnoDB ROW_FORMAT=COMPRESSED KEY_BLOCK_SIZE=8;