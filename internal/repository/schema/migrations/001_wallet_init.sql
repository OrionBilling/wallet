USE billing_system;

-- Table for the current wallet state
CREATE TABLE wallet (
    user_id UUID PRIMARY KEY,
    balance BIGINT static, -- Static column, stored once per partition
    reserved BIGINT UNSIGNED static, -- Static column, stored once per partition
    currency SMALLINT UNSIGNED static, -- ISO-4217 (3 digits)
    last_updated timestamp,

) WITH compaction = { 'class' : 'LeveledCompactionStrategy' }; -- Good for read-heavy data

-- Table for the immutable ledger of all operations
CREATE TABLE wallet_ledger (
    user_id UUID NOT NULL,
    operation_id UUID NOT NULL, 
    operation_type ENUM('charge', 'reserve', 'reserve_end', 'reserve_part_end', 'fix', 'refund'),
    time_bucket TIMESTAMP NOT NULL,  -- truncated to month level
    amount BIGINT,
    descr text,
    reference_id UUID, -- e.g., ID of the resource that was charged
    new_balance decimal,
    new_reserved decimal,
    created_at timestamp,

    PRIMARY KEY (user_id, operation_id)
) PARTITION BY RANGE (time_bucket);

-- Each  new months the new partition should be created like:
--
-- CREATE TABLE wallet_ledger_2025_11 PARTITION OF wallet_ledger
--     FOR VALUES FROM ('2025-11-01') TO ('2025-12-01');
