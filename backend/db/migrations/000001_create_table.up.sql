CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(20) PRIMARY KEY,
    pw VARCHAR(20) NOT NULL,
    private_key VARCHAR(2048) NOT NULL,
    public_key VARCHAR(2048) NOT NULL
);
CREATE TABLE IF NOT EXISTS subjects(
    subject_code VARCHAR(20) PRIMARY KEY,
    subject_name VARCHAR(20) NOT NULL
)