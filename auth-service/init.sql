CREATE TABLE ADMINS(
    ID SERIAL PRIMARY KEY,
    LOGIN VARCHAR(255) UNIQUE NOT NULL,
    PWD VARCHAR(255) NOT NULL
);

CREATE TABLE MANAGERS(
    ID SERIAL PRIMARY KEY,
    LOGIN VARCHAR(255) UNIQUE NOT NULL,
    PWD VARCHAR(255) NOT NULL
);

CREATE TABLE CLIENTS(
    ID SERIAL PRIMARY KEY,
    LOGIN VARCHAR(255) UNIQUE NOT NULL,
    PWD VARCHAR(255) NOT NULL
);

INSERT INTO ADMINS VALUES(default, 'uid', 'admin', 'admin');
INSERT INTO ADMINS VALUES(default, 'uid1', 'admin1', 'admin1');
INSERT INTO ADMINS VALUES(default, 'uid2', 'admin2', 'admin2');