CREATE TABLE songs (
                       id SERIAL PRIMARY KEY,
                       group VARCHAR(255) NOT NULL,
                       song VARCHAR(255) NOT NULL,
                       text TEXT,
                       link VARCHAR(255),
                       releaseDate DATE,
                       api_fetched BOOLEAN DEFAULT FALSE
);