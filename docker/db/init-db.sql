create table users
(
    id         serial not null
        constraint users_pkey
            primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    login      text,
    password   text,
    role       text
);

alter table users
    owner to app;

create index idx_users_deleted_at
    on users (deleted_at);

create index idx_users_login
    on users (login);

INSERT INTO public.users (id, created_at, updated_at, deleted_at, login, password, role) VALUES (1, '2021-03-20 12:27:35.910000', '2021-03-20 12:27:37.016000', null, 'admin', '$2a$14$z.GRSWrKfsg6xsbtC4fMveXmTZKar53a.PJ1StRQjHmHn6qFf4.G.', 'admin');