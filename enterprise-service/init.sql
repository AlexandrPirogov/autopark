CREATE TABLE ENTERPRISES(
    ID SERIAL PRIMARY KEY,
    TITLE VARCHAR(500) UNIQUE NOT NULL CHECK (replace(TITLE, ' ', '') <> '')
);


CREATE TABLE MANAGERS(
    ID SERIAL PRIMARY KEY,
    E_ID INT NOT NULL,
    NAME VARCHAR(500) NOT NULL CHECK (replace(NAME, ' ', '') <> ''),
    SURNAME VARCHAR(500) NOT NULL CHECK (replace(SURNAME, ' ', '') <> ''),
    CONSTRAINT FK_ENTERPRISES FOREIGN KEY (E_ID) REFERENCES ENTERPRISES(ID)
);