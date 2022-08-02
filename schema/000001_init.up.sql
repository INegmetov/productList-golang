CREATE TABLE product_lists
(
    id          serial       not null unique,
    title       varchar(255) not null
);

CREATE TABLE product_items
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    kcal int
);


CREATE TABLE lists_items
(
    id      serial not null unique,
    item_id int references product_items (id) on delete cascade not null,
    list_id int references product_lists (id) on delete cascade not null
);