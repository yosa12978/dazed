CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR PRIMARY KEY,
    username VARCHAR UNIQUE NOT NULL,
    password_hash VARCHAR NOT NULL,
    salt VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    role VARCHAR NOT NULL DEFAULT 'USER'
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_accounts_username 
ON accounts(username);

CREATE TABLE IF NOT EXISTS categories(
    id VARCHAR PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    parent_id VARCHAR NULL,
    FOREIGN KEY (parent_id) REFERENCES categories(id)
);

CREATE INDEX IF NOT EXISTS idx_categories_parent_id
ON categories(parent_id);

CREATE TABLE IF NOT EXISTS articles (
    id VARCHAR PRIMARY KEY,
    title VARCHAR NOT NULL,
    body VARCHAR NOT NULL,
    description VARCHAR NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    category_id VARCHAR NULL,
    author_id VARCHAR NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL,
    FOREIGN KEY (author_id) REFERENCES accounts(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_articles_category_id
ON articles(category_id);

CREATE INDEX IF NOT EXISTS idx_articles_author_id
ON articles(author_id);