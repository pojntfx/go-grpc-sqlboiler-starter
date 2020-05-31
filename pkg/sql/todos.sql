drop table if exists todos;
create table todos (
  id serial not null primary key,
  title text not null,
  body text not null,
  index bigint not null,
  namespace text not null
)