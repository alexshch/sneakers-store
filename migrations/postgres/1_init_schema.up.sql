create table if not exists users
(
    id uuid primary key,
    full_name varchar not null,
    email varchar not null,
    pwd_hash varchar not null,
    created_at timestamptz not null,
    updated_at timestamptz not null
);

create table if not exists sneaker
(
    id uuid primary key,
    title varchar(256) not null,
    price int not null,
    image_url varchar not null,
    is_available bool not null,
    created_at timestamptz not null,
    updated_at timestamptz not null
);

create table if not exists orders
(
    id uuid primary key,
    user_id uuid not null,
    status int not null,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    finished_at timestamptz not null,
    constraint orders_user_id foreign key (user_id) references users(id) on delete cascade
);

create table if not exists orders_item
(
    order_id uuid not null,
    sneaker_id uuid not null,
    constraint orders_item_orders_id foreign key (order_id) references orders(id) on delete cascade,
    constraint orders_item_sneaker_id foreign key (sneaker_id) references sneaker(id) on delete cascade
);