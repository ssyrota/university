-- +goose Up
-- +goose StatementBegin

CREATE TABLE users (
  login VARCHAR(255) PRIMARY KEY,
  password VARCHAR(255) NOT NULL
);
CREATE TABLE cities (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);
CREATE TABLE hobbies (
  id VARCHAR(255) PRIMARY KEY,
  name VARCHAR(255) NOT NULL
);

CREATE TABLE cvs (
  id VARCHAR(255) PRIMARY KEY,
  "user" VARCHAR(255) NOT NULL REFERENCES users(login)
);

CREATE TABLE cvs_hobbies (
  cv_id VARCHAR(255) NOT NULL REFERENCES cvs(id),
  hobby_id VARCHAR(255) NOT NULL REFERENCES hobbies(id),
  PRIMARY KEY (cv_id, hobby_id)
);

CREATE TABLE jobs (
  id VARCHAR(255) PRIMARY KEY,
  cv_id VARCHAR(255) NOT NULL REFERENCES cvs(id),
  city_id VARCHAR(255) NOT NULL REFERENCES cities(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE jobs;
DROP TABLE cvs_hobbies;
DROP TABLE cvs;
DROP TABLE hobbies;
DROP TABLE cities;
DROP TABLE users;
-- +goose StatementEnd
