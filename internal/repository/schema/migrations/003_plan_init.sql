-- FixMe: add correct partitioning to place users by companies and avoid hot partitions!

USE billing_system;

-- Account types enumeration
CREATE TYPE plan_type AS ENUM ('payg', 'subscription', 'hybrid');
CREATE TYPE billing_cycle AS ENUM ('monthly', 'quarterly', 'annual', 'custom');
CREATE TYPE plan_status AS ENUM ('active', 'inactive', 'suspended', 'pending');

-- Billing Plans Table
CREATE TABLE IF NOT EXISTS plan (
    plan_id UUID PRIMARY KEY,
    plan_name VARCHAR(100) NOT NULL,
    plan_description TEXT,
    plan_status ENUM('active', 'inactive', 'closed'),
    price BIGINT UNSIGNED NOT NULL, -- in the smallest currency unit
    currency SMALLINT UNSIGNED, -- ISO-4217 (3 digits)
    billing_cycle billing_cycle NOT NULL,
   -- Resource limits
    shared_resources JSONB, -- For company/group shared resources
    user_resources JSONB,   -- Per-user resource limits
    
    -- PAYG specific fields
    payg_rates JSONB, -- {"cpu_per_hour": 0.05, "storage_gb_per_month": 0.10}

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    INDEX idx_plan_active (is_active),
    INDEX idx_plan_cycle (billing_cycle)
) ENGINE=InnoDB ROW_FORMAT=COMPRESSED KEY_BLOCK_SIZE=8;


-- Account plan assignments
CREATE TABLE account_plan(
    account_plan_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_id UUID NOT NULL,
    plan_id UUID NOT NULL,
    status plan_status NOT NULL DEFAULT 'active',
    starts_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ends_at TIMESTAMP WITH TIME ZONE NULL, -- NULL means no end date
    auto_renew BOOLEAN NOT NULL DEFAULT true,
    custom_price_amount DECIMAL(15,4) NULL, -- override base price if needed
    custom_price_currency VARCHAR(3) NULL,
    assigned_by UUID NULL, -- who assigned this plan (for audit)
    metadata JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    -- Constraints
    CONSTRAINT fk_account_plan_account 
        FOREIGN KEY (account_id) 
        REFERENCES accounts(account_id) 
        ON DELETE CASCADE,
    CONSTRAINT fk_account_plan_plan 
        FOREIGN KEY (plan_id) 
        REFERENCES billing_plans(plan_id) 
        ON DELETE RESTRICT,
    CONSTRAINT fk_account_plan_assigned_by 
        FOREIGN KEY (assigned_by) 
        REFERENCES accounts(account_id) 
        ON DELETE SET NULL,

);


-- Account plans junction table (partitioned by company_id)
CREATE TABLE account_plans (
    account_plan_id UUID DEFAULT uuid_generate_v4(),
    account_id UUID NOT NULL,
    company_id UUID NOT NULL, -- For partitioning with accounts
    plan_id UUID NOT NULL,
    
    -- Plan assignment context
    assigned_by UUID, -- Who assigned this plan
    assigned_at TIMESTAMPTZ DEFAULT NOW(),
    
    -- Plan status and dates
    plan_status plan_status DEFAULT 'active',
    activated_at TIMESTAMPTZ,
    expires_at TIMESTAMPTZ,
    
    -- Billing information
    current_billing_cycle_start TIMESTAMPTZ,
    current_billing_cycle_end TIMESTAMPTZ,
     
    -- Custom overrides
    custom_limits JSONB, -- Override default plan limits
    custom_rates JSONB,  -- Override default plan rates
    
    PRIMARY KEY (account_plan_id, company_id),
    FOREIGN KEY (account_id, company_id) REFERENCES accounts(account_id, company_id),
    FOREIGN KEY (plan_id) REFERENCES billing_plans(plan_id)
)