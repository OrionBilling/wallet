USE billing_system;

-- Table for the current wallet state
CREATE TABLE wallet (
    user_id UUID,
    balance decimal static, -- Static column, stored once per partition
    reserved decimal static, -- Static column, stored once per partition
    currency text static,
    last_updated timestamp,
    PRIMARY KEY (user_id)
) WITH compaction = { 'class' : 'LeveledCompactionStrategy' }; -- Good for read-heavy data

-- Table for the immutable ledger of all operations
CREATE TABLE wallet_ledger (
    user_id UUID,
    operation_id timeuuid, -- Ensures time ordering
    type decimal, -- 0-'CHARGE', 1- 'RESERVE', 2- 'RESERVE_END', 3 -'FIX', 4- 'REFUND'
    amount decimal,
    description text,
    reference_id text, -- e.g., ID of the resource that was charged
    new_balance decimal,
    new_reserved decimal,
    created_at timestamp,
    PRIMARY KEY (user_id, operation_id)
) WITH CLUSTERING ORDER BY (operation_id DESC) -- Newest operations first
  AND compaction = { 'class' : 'SizeTieredCompactionStrategy' }; -- Good for write-heavy data