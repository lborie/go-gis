alter table "regions-20180101" add column geom_simplified geometry(MultiPolygon,4326);
create index r_geom_simpl_idx on "regions-20180101" using gist(geom_simplified);
update "regions-20180101" set geom_simplified = st_makevalid(st_simplify(geom, 0.001));

alter table "departements-20180101" add column geom_simplified geometry(MultiPolygon,4326);
create index d_geom_simpl_idx on "departements-20180101" using gist(geom_simplified);
update "departements-20180101" set geom_simplified = st_makevalid(st_simplify(geom, 0.001));