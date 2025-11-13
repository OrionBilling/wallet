-- Create database
CREATE DATABASE IF NOT EXISTS billing_system 

CHARACTER SET utf8mb4 

COLLATE utf8mb4_unicode_ci;

USE billing_system;

-- Create the user_requests table with sharding by UserID + time bucket
CREATE TABLE IF NOT EXISTS user_requests (
    user_id UUID NOT NULL,
    request_id UUID NOT NULL,
    time_bucket TIMESTAMP NOT NULL,  -- truncated to month level
    created_at TIMESTAMP DEFAULT NOW(),
    payload JSONB,
    status VARCHAR(50) DEFAULT 'pending',
    PRIMARY KEY (user_id, time_bucket DESC, request_id),
    UNIQUE (request_id, user_id)
);

-- Create index for efficient time-based queries
CREATE INDEX IF NOT EXISTS idx_user_requests_time_bucket 
ON user_requests (time_bucket DESC);

-- Create index for status-based queries
CREATE INDEX IF NOT EXISTS idx_user_requests_status 
ON user_requests (status);
