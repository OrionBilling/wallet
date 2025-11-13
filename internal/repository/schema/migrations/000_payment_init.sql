-- Create database
CREATE DATABASE billing_system 

CHARACTER SET utf8mb4 

COLLATE utf8mb4_unicode_ci;

USE billing_system;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


-- Create payment orders table with sharding by UserID + time bucket
CREATE TABLE payment_order(
    user_id UUID NOT NULL,
    request_id UUID NOT NULL, -- idempotency ID
    time_bucket TIMESTAMP NOT NULL,  -- truncated to month level
    created_at TIMESTAMP NOT NULL,,
    updated_at TIMESTAMP,
    
    amount BIGINT UNSIGNED,
    currency SMALLINT UNSIGNED, -- ISO-4217 (3 digits)
    order_data JSONB,

    order_status ENUM('init', 'processing', 'success', 'fail'),

    PRIMARY KEY (user_id, time_bucket DESC, request_id),
    UNIQUE (request_id, user_id)
) PARTITION BY RANGE (time_bucket);

-- Create index for efficient time-based queries
CREATE INDEX  idx_user_requests_time_bucket 
ON user_requests (time_bucket DESC);

-- Create index for status-based queries
CREATE INDEX idx_user_requests_status 
ON user_requests (order_status);

-- Each  new mots the new partition should be created like:
--
-- CREATE TABLE payment_order_2025_11 PARTITION OF payment_order
--     FOR VALUES FROM ('2025-11-01') TO ('2025-12-01');
