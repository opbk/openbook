CREATE OR REPLACE FUNCTION update_modified_column()	
RETURNS TRIGGER AS $$
BEGIN
    NEW.modified = now();
    RETURN NEW;	
END;
$$ language 'plpgsql';

-- book table --
CREATE SEQUENCE auto_id_books;
CREATE TABLE books (
	id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('auto_id_books'),
	title CHARACTER VARYING (500) NOT NULL,
	description TEXT,
	release DATE
);

-- category tables -- 
CREATE TABLE book_categories (
	book_id INTEGER NOT NULL,
	category_id INTEGER NOT NULL,
	PRIMARY KEY (book_id, category_id)
);
CREATE INDEX book_categories_cid_idx on book_categories (category_id);

CREATE SEQUENCE auto_id_categories;
CREATE TABLE categories (
	id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('auto_id_categories'),
	category_id INTEGER NOT NULL DEFAULT 0,
	name CHARACTER VARYING (250) NOT NULL
);
CREATE INDEX categories_cid_idx ON categories (category_id);

-- author tables --
CREATE TABLE author_books (
	book_id INTEGER NOT NULL,
	author_id INTEGER NOT NULL,
	PRIMARY KEY (book_id, author_id)
);
CREATE INDEX author_books_aid_idx on author_books (author_id);

CREATE SEQUENCE auto_id_authors;
CREATE TABLE authors (
	id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('auto_id_authors'),
	name CHARACTER VARYING (500) NOT NULL,
	description TEXT
);

-- user and address tables --
CREATE SEQUENCE auto_id_users;
CREATE TABLE users (
  id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('auto_id_users'),
  email CHARACTER VARYING(250) NOT NULL,
  password CHARACTER VARYING(250),
  name TEXT,
  created TIMESTAMP NOT NULL DEFAULT NOW(),
  modified TIMESTAMP NOT NULL DEFAULT NOW(),
  last_enter TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE INDEX users_email_idx ON users (email);
CREATE TRIGGER update_users_modtime BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE SEQUENCE auto_id_addresses;
CREATE TABLE addresses (
	id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('auto_id_addresses'),
	user_id INTEGER NOT NULL,
	address TEXT NOT NULL,
	comment TEXT
);
CREATE INDEX addresses_uid_idx ON addresses (user_id);

-- subscription tables --
CREATE SEQUENCE auto_id_subscriptions;
CREATE TABLE subscriptions (
  id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('auto_id_subscriptions'),
  name CHARACTER VARYING(250) NOT NULL,
  description TEXT,
  price NUMERIC NOT NULL
);

CREATE TABLE user_subscriptions (
	user_id INTEGER NOT NULL,
	subscription_id INTEGER NOT NULL,
	PRIMARY KEY (user_id, subscription_id)
);
CREATE INDEX user_subscriptions_sid_idx on user_subscriptions (subscription_id);

CREATE TABLE book_subscriptions (
	book_id INTEGER NOT NULL,
	subscription_id INTEGER NOT NULL,
	PRIMARY KEY (book_id, subscription_id)
);
CREATE INDEX book_subscriptions_sid_idx on book_subscriptions (subscription_id);

-- order tables --
CREATE SEQUENCE auto_id_orders;
CREATE TABLE orders (
  id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('auto_id_orders'),
  user_id INTEGER NOT NULL,
  address_id INTEGER NOT NULL,
  status CHARACTER VARYING(50) NOT NULL,
  comment TEXT,
  created TIMESTAMP NOT NULL DEFAULT NOW(),
  modified TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE INDEX orders_uid_idx on orders (user_id);
CREATE TRIGGER update_orders_modtime BEFORE UPDATE ON orders FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE TABLE book_orders (
	order_id INTEGER NOT NULL,
	book_id INTEGER NOT NULL,
	price_type_id INTEGER NOT NULL,
	PRIMARY KEY (order_id, book_id)
);
CREATE INDEX book_orders_bid_idx on book_orders (book_id);

-- price tables --
CREATE SEQUENCE auto_id_prices;
CREATE TABLE prices (
  id INTEGER NOT NULL PRIMARY KEY DEFAULT nextval('auto_id_prices'),
  name CHARACTER VARYING(250) NOT NULL
);

CREATE TABLE book_prices (
	book_id INTEGER NOT NULL,
	price_type_id INTEGER NOT NULL,
	price NUMERIC NOT NULL,
	PRIMARY KEY (book_id, price_type_id)
);
CREATE INDEX book_prices_ptid_idx on book_prices (price_type_id);