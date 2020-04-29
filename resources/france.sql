create table "france" as select st_makevalid(st_simplify(st_union(geom), 0.001)) as geom from "regions-20180101";
create index f_geom_simpl_idx on "france" using gist(geom);
