create table buses {
	id int primary key,
	number int,
	driver_id int,
	route_id int,
	created_at datetime,
	updated_at datetime
}

create table drivers {
	id int primary key,
}