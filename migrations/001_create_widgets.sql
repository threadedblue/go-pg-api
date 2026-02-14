create table if not exists widgets (
  id bigserial primary key,
  name text not null,
  created_at timestamptz not null default now()
);

insert into widgets (name) values ('alpha'), ('beta')
on conflict do nothing;
