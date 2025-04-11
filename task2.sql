CREATE TABLE Teams (
    id INT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    directorId INT,
    FOREIGN KEY (directorId) REFERENCES Employee(id)
);

CREATE TABLE Positions (
    id INT PRIMARY KEY,
    title VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE Employee (
    id INT PRIMARY KEY,
    teamId INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    managerId INT,
    positionId INT,
    FOREIGN KEY (teamId) REFERENCES Teams(id),
    FOREIGN KEY (managerId) REFERENCES Employee(id),
    FOREIGN KEY (positionId) REFERENCES Positions(id)
);

WITH RECURSIVE subordinates AS (
    SELECT id, name, managerId
    FROM Employee
    WHERE managerId = 1

    UNION ALL

    SELECT e.id, e.name, e.managerId
    FROM Employee e
    JOIN subordinates s ON e.managerId = s.id
)
SELECT * FROM subordinates;

WITH RECURSIVE managers AS (
    SELECT id, name, managerId
    FROM Employee
    WHERE id = 1

    UNION ALL

    SELECT e.id, e.name, e.managerId
    FROM Employee e
    JOIN managers m ON e.id = m.managerId
)
SELECT * FROM managers;

