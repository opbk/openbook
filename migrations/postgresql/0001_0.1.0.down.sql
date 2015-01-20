DROP TABLE IF EXISTS books CASCADE;
DROP TABLE IF EXISTS book_categories CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS author_books CASCADE;
DROP TABLE IF EXISTS authors CASCADE;
DROP TABLE IF EXISTS publishers CASCADE;
DROP TABLE IF EXISTS series CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS addresses CASCADE;
DROP TABLE IF EXISTS subscriptions CASCADE;
DROP TABLE IF EXISTS user_subscriptions CASCADE;
DROP TABLE IF EXISTS book_subscriptions CASCADE;
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS book_orders CASCADE;
DROP TABLE IF EXISTS prices CASCADE;
DROP TABLE IF EXISTS book_prices CASCADE;

DROP SEQUENCE IF EXISTS auto_id_books;
DROP SEQUENCE IF EXISTS auto_id_categories;
DROP SEQUENCE IF EXISTS auto_id_authors;
DROP SEQUENCE IF EXISTS auto_id_publishers;
DROP SEQUENCE IF EXISTS auto_id_series;
DROP SEQUENCE IF EXISTS auto_id_users;
DROP SEQUENCE IF EXISTS auto_id_addresses;
DROP SEQUENCE IF EXISTS auto_id_subscriptions;
DROP SEQUENCE IF EXISTS auto_id_orders;
DROP SEQUENCE IF EXISTS auto_id_prices;
