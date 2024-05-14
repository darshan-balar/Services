CREATE TABLE Services (                                                                                   
    service_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    version_count INT DEFAULT 0,
    created_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Versions (                                                                                   
    version_id SERIAL PRIMARY KEY,
    service_id INT NOT NULL,
    version_name VARCHAR(50) NOT NULL,
    description TEXT,
    FOREIGN KEY (service_id) REFERENCES Services(service_id)
);

INSERT INTO Services (name, description,version_count) VALUES ('Service A', 'Description for Service A',2);
INSERT INTO Services (name, description,version_count) VALUES ('Service B', 'Description for Service B',1);
INSERT INTO Services (name, description,version_count) VALUES ('Service C', 'Description for Service C',3);

INSERT INTO Versions (service_id, version_name, description) VALUES
  (1, 'v1', 'Description for Service A Version 1'),
  (1, 'v2', 'Description for Service A Version 2'),
  (2, 'v1', 'Description for Service B Version 1'),
  (3, 'v1', 'Description for Service C Version 1'),
  (3, 'v2', 'Description for Service C Version 2'),
  (3, 'v3', 'Description for Service C Version 3');