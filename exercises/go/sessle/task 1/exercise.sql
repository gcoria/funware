create table buses (
	id int primary key,
	origin varchar(255),
	destination varchar(255),
	time varchar,
	unique(origin, destination, time)
)

create table passengers (
	id int primary key,
	origin varchar(255),
	destination varchar(255),
	time varchar,
)

-- Query to find buses and their passenger counts
WITH boarded_passengers AS (
    SELECT 
        p.id AS passenger_id,
        MIN(b.id) AS bus_id
    FROM 
        passengers p
    LEFT JOIN 
        buses b ON p.origin = b.origin 
        AND p.destination = b.destination 
        AND b.time >= p.time
        AND p.time <= '23:59'
    GROUP BY 
        p.id
)
SELECT 
    b.id AS bus_id,
    COUNT(bp.passenger_id) AS passengers_on_board
FROM 
    buses b
LEFT JOIN 
    boarded_passengers bp ON b.id = bp.bus_id
GROUP BY 
    b.id
ORDER BY 
    b.id ASC;



