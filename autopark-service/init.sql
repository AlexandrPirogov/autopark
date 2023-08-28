CREATE TABLE BRANDS(
    ID SERIAL PRIMARY KEY,
    BRAND VARCHAR(500)  UNIQUE NOT NULL CHECK (replace(BRAND, ' ', '') <> '')
);

CREATE TYPE CAR_TYPE AS ENUM('sedan', 'hatchback', 'sport');

CREATE TABLE CARS(
    ID SERIAL PRIMARY KEY,
    BRAND_ID INT NOT NULL,
    UID VARCHAR(200) UNIQUE NOT NULL CHECK (replace(UID, ' ', '') <> ''),
    TYPE CAR_TYPE NOT NULL,
    CONSTRAINT FK_CAR_BRAND FOREIGN KEY(BRAND_ID) REFERENCES BRANDS(ID)
);

CREATE TYPE CAR_STATUS AS ENUM('set', 'booked', 'unset');


CREATE TABLE CARS_FOR_BOOKING (
    ID SERIAL PRIMARY KEY,
    CAR_ID INT NOT NULL,
    STATUS CAR_STATUS default 'unset',
    WATCHING INT,
    CONSTRAINT FK_CAR_FOR_BOOKING FOREIGN KEY (CAR_ID) REFERENCES CARS(ID) 
);

INSERT INTO BRANDS values(default, 'bmw');
INSERT INTO BRANDS values(default, 'mets');
INSERT INTO BRANDS values(default, 'gaz');

INSERT INTO CARS values(default, 1, 'bmw1', 'sport');
INSERT INTO CARS values(default, 2, 'mers1', 'sport');
INSERT INTO CARS values(default, 3, 'gaz1', 'sport');
INSERT INTO CARS values(default, 1, 'bmw2', 'sport');
INSERT INTO CARS values(default, 2, 'mers2', 'sport');
INSERT INTO CARS values(default, 3, 'gaz3', 'sport');

INSERT INTO CARS_FOR_BOOKING values(default, 1, 'set', 0);
INSERT INTO CARS_FOR_BOOKING values(default, 2, 'unset', 0);
INSERT INTO CARS_FOR_BOOKING values(default, 3, 'set', 0);