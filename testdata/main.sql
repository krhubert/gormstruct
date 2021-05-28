CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS hstore;
CREATE EXTENSION IF NOT EXISTS ltree;
CREATE EXTENSION IF NOT EXISTS isn;
CREATE EXTENSION IF NOT EXISTS cube;

CREATE TABLE T (
  cn1 bigint,
  cnn1 bigint NOT NULL,
  ca1 bigint[],
  cann1 bigint[] NOT NULL,

  cn2 int8,
  cnn2 int8 NOT NULL,
  ca2 int8[],
  cann2 int8[] NOT NULL,

  cn3 bigserial,
  cnn3 bigserial NOT NULL,
  -- ca3 bigserial[], ERROR: array of serial is not implemented
  -- cann3 bigserial[] NOT NULL, ERROR: array of serial is not implemented

  cn4 serial8,
  cnn4 serial8 NOT NULL,
  -- ca4 serial8[], ERROR: array of serial is not implemented
  -- cann4 serial8[] NOT NULL, ERROR: array of serial is not implemented

  cn5 bit,
  cnn5 bit NOT NULL,
  ca5 bit[],
  cann5 bit[] NOT NULL,

  cn6 bit(2),
  cnn6 bit(2) NOT NULL,
  ca6 bit(2)[],
  cann6 bit(2)[] NOT NULL,

  cn7 bit varying,
  cnn7 bit varying NOT NULL,
  ca7 bit varying[],
  cann7 bit varying[] NOT NULL,

  cn8 bit varying(2),
  cnn8 bit varying(2) NOT NULL,
  ca8 bit varying(2)[],
  cann8 bit varying(2)[] NOT NULL,

  cn9 boolean,
  cnn9 boolean NOT NULL,
  ca9 boolean[],
  cann9 boolean[] NOT NULL,

  cn10 bool,
  cnn10 bool NOT NULL,
  ca10 bool[],
  cann10 bool[] NOT NULL,

  cn11 box,
  cnn11 box NOT NULL,
  ca11 box[],
  cann11 box[] NOT NULL,

  cn12 bytea,
  cnn12 bytea NOT NULL,
  ca12 bytea[],
  cann12 bytea[] NOT NULL,

  cn13 character,
  cnn13 character NOT NULL,
  ca13 character[],
  cann13 character[] NOT NULL,

  cn14 character(2),
  cnn14 character(2) NOT NULL,
  ca14 character(2)[],
  cann14 character(2)[] NOT NULL,

  cn15 char,
  cnn15 char NOT NULL,
  ca15 char[],
  cann15 char[] NOT NULL,

  cn16 char(2),
  cnn16 char(2) NOT NULL,
  ca16 char(2)[],
  cann16 char(2)[] NOT NULL,

  cn17 character varying,
  cnn17 character varying NOT NULL,
  ca17 character varying[],
  cann17 character varying[] NOT NULL,

  cn18 character varying(2),
  cnn18 character varying(2) NOT NULL,
  ca18 character varying(2)[],
  cann18 character varying(2)[] NOT NULL,

  cn19 varchar,
  cnn19 varchar NOT NULL,
  ca19 varchar[],
  cann19 varchar[] NOT NULL,

  cn20 varchar(2),
  cnn20 varchar(2) NOT NULL,
  ca20 varchar(2)[],
  cann20 varchar(2)[] NOT NULL,

  cn21 cidr,
  cnn21 cidr NOT NULL,
  ca21 cidr[],
  cann21 cidr[] NOT NULL,

  cn22 circle,
  cnn22 circle NOT NULL,
  ca22 circle[],
  cann22 circle[] NOT NULL,

  cn23 date,
  cnn23 date NOT NULL,
  ca23 date[],
  cann23 date[] NOT NULL,

  cn24 double precision,
  cnn24 double precision NOT NULL,
  ca24 double precision[],
  cann24 double precision[] NOT NULL,

  cn25 float8,
  cnn25 float8 NOT NULL,
  ca25 float8[],
  cann25 float8[] NOT NULL,

  cn26 inet,
  cnn26 inet NOT NULL,
  ca26 inet[],
  cann26 inet[] NOT NULL,

  cn27 integer,
  cnn27 integer NOT NULL,
  ca27 integer[],
  cann27 integer[] NOT NULL,

  cn28 int,
  cnn28 int NOT NULL,
  ca28 int[],
  cann28 int[] NOT NULL,

  cn29 int4,
  cnn29 int4 NOT NULL,
  ca29 int4[],
  cann29 int4[] NOT NULL,

  cn30 interval,
  cnn30 interval NOT NULL,
  ca30 interval[],
  cann30 interval[] NOT NULL,

  cn31 json,
  cnn31 json NOT NULL,
  ca31 json[],
  cann31 json[] NOT NULL,

  cn32 jsonb,
  cnn32 jsonb NOT NULL,
  ca32 jsonb[],
  cann32 jsonb[] NOT NULL,

  cn33 line,
  cnn33 line NOT NULL,
  ca33 line[],
  cann33 line[] NOT NULL,

  cn34 lseg,
  cnn34 lseg NOT NULL,
  ca34 lseg[],
  cann34 lseg[] NOT NULL,

  cn35 macaddr,
  cnn35 macaddr NOT NULL,
  ca35 macaddr[],
  cann35 macaddr[] NOT NULL,

  cn36 macaddr8,
  cnn36 macaddr8 NOT NULL,
  ca36 macaddr8[],
  cann36 macaddr8[] NOT NULL,

  cn37 money,
  cnn37 money NOT NULL,
  ca37 money[],
  cann37 money[] NOT NULL,

  cn38 numeric,
  cnn38 numeric NOT NULL,
  ca38 numeric[],
  cann38 numeric[] NOT NULL,

  cn39 numeric(1, 1),
  cnn39 numeric(1, 1) NOT NULL,
  ca39 numeric(1, 1)[],
  cann39 numeric(1, 1)[] NOT NULL,

  cn40 decimal,
  cnn40 decimal NOT NULL,
  ca40 decimal[],
  cann40 decimal[] NOT NULL,

  cn41 decimal(1, 1),
  cnn41 decimal(1, 1) NOT NULL,
  ca41 decimal(1, 1)[],
  cann41 decimal(1, 1)[] NOT NULL,

  cn42 path,
  cnn42 path NOT NULL,
  ca42 path[],
  cann42 path[] NOT NULL,

  cn43 pg_lsn,
  cnn43 pg_lsn NOT NULL,
  ca43 pg_lsn[],
  cann43 pg_lsn[] NOT NULL,

  cn44 point,
  cnn44 point NOT NULL,
  ca44 point[],
  cann44 point[] NOT NULL,

  cn45 polygon,
  cnn45 polygon NOT NULL,
  ca45 polygon[],
  cann45 polygon[] NOT NULL,

  cn46 real,
  cnn46 real NOT NULL,
  ca46 real[],
  cann46 real[] NOT NULL,

  cn47 float4,
  cnn47 float4 NOT NULL,
  ca47 float4[],
  cann47 float4[] NOT NULL,

  cn48 smallint,
  cnn48 smallint NOT NULL,
  ca48 smallint[],
  cann48 smallint[] NOT NULL,

  cn49 int2,
  cnn49 int2 NOT NULL,
  ca49 int2[],
  cann49 int2[] NOT NULL,

  cn50 smallserial,
  cnn50 smallserial NOT NULL,
  -- ca50 smallserial[], ERROR: array of serial is not implemented
  -- cann50 smallserial[] NOT NULL, ERROR: array of serial is not implemented

  cn51 serial2,
  cnn51 serial2 NOT NULL,
  -- ca51 serial2[], ERROR: array of serial is not implemented
  -- cann51 serial2[] NOT NULL, ERROR: array of serial is not implemented

  cn52 serial,
  cnn52 serial NOT NULL,
  -- ca52 serial[], ERROR: array of serial is not implemented
  -- cann52 serial[] NOT NULL, ERROR: array of serial is not implemented

  cn53 serial4,
  cnn53 serial4 NOT NULL,
  -- ca53 serial4[], ERROR: array of serial is not implemented
  -- cann53 serial4[] NOT NULL, ERROR: array of serial is not implemented

  cn54 text,
  cnn54 text NOT NULL,
  ca54 text[],
  cann54 text[] NOT NULL,

  cn55 time,
  cnn55 time NOT NULL,
  ca55 time[],
  cann55 time[] NOT NULL,

  cn56 time without time zone,
  cnn56 time without time zone NOT NULL,
  ca56 time without time zone[],
  cann56 time without time zone[] NOT NULL,

  cn57 time with time zone,
  cnn57 time with time zone NOT NULL,
  ca57 time with time zone[],
  cann57 time with time zone[] NOT NULL,

  cn58 timetz,
  cnn58 timetz NOT NULL,
  ca58 timetz[],
  cann58 timetz[] NOT NULL,

  cn59 timestamp,
  cnn59 timestamp NOT NULL,
  ca59 timestamp[],
  cann59 timestamp[] NOT NULL,

  cn60 timestamp without time zone,
  cnn60 timestamp without time zone NOT NULL,
  ca60 timestamp without time zone[],
  cann60 timestamp without time zone[] NOT NULL,

  cn61 timestamp with time zone,
  cnn61 timestamp with time zone NOT NULL,
  ca61 timestamp with time zone[],
  cann61 timestamp with time zone[] NOT NULL,

  cn62 timestamptz,
  cnn62 timestamptz NOT NULL,
  ca62 timestamptz[],
  cann62 timestamptz[] NOT NULL,

  cn63 tsquery,
  cnn63 tsquery NOT NULL,
  ca63 tsquery[],
  cann63 tsquery[] NOT NULL,

  cn64 tsvector,
  cnn64 tsvector NOT NULL,
  ca64 tsvector[],
  cann64 tsvector[] NOT NULL,

  cn65 txid_snapshot,
  cnn65 txid_snapshot NOT NULL,
  ca65 txid_snapshot[],
  cann65 txid_snapshot[] NOT NULL,

  cn66 uuid,
  cnn66 uuid NOT NULL,
  ca66 uuid[],
  cann66 uuid[] NOT NULL,

  cn67 xml,
  cnn67 xml NOT NULL,
  ca67 xml[],
  cann67 xml[] NOT NULL,

  cn68 oid,
  cnn68 oid NOT NULL,
  ca68 oid[],
  cann68 oid[] NOT NULL,

  cn69 oidvector,
  cnn69 oidvector NOT NULL,
  ca69 oidvector[],
  cann69 oidvector[] NOT NULL,

  cn70 int4,
  cnn70 int4 NOT NULL,
  ca70 int4[],
  cann70 int4[] NOT NULL,

  cn71 bpchar,
  cnn71 bpchar NOT NULL,
  ca71 bpchar[],
  cann71 bpchar[] NOT NULL,

  cn72 cid,
  cnn72 cid NOT NULL,
  ca72 cid[],
  cann72 cid[] NOT NULL,

  cn74 int2vector,
  cnn74 int2vector NOT NULL,
  ca74 int2vector[],
  cann74 int2vector[] NOT NULL,

  cn75 int4range,
  cnn75 int4range NOT NULL,
  ca75 int4range[],
  cann75 int4range[] NOT NULL,

  cn76 int8range,
  cnn76 int8range NOT NULL,
  ca76 int8range[],
  cann76 int8range[] NOT NULL,

  cn77 tid,
  cnn77 tid NOT NULL,
  ca77 tid[],
  cann77 tid[] NOT NULL,

  cn78 tsrange,
  cnn78 tsrange NOT NULL,
  ca78 tsrange[],
  cann78 tsrange[] NOT NULL,

  cn79 tstzrange,
  cnn79 tstzrange NOT NULL,
  ca79 tstzrange[],
  cann79 tstzrange[] NOT NULL,

  cn82 xid,
  cnn82 xid NOT NULL,
  ca82 xid[],
  cann82 xid[] NOT NULL,

  cn83 xid8,
  cnn83 xid8 NOT NULL,
  ca83 xid8[],
  cann83 xid8[] NOT NULL,

  cn84 daterange,
  cnn84 daterange NOT NULL,
  ca84 daterange[],
  cann84 daterange[] NOT NULL,

  cn85 gtsvector,
  cnn85 gtsvector NOT NULL,
  ca85 gtsvector[],
  cann85 gtsvector[] NOT NULL,

  cn86 jsonpath,
  cnn86 jsonpath NOT NULL,
  ca86 jsonpath[],
  cann86 jsonpath[] NOT NULL,

  cn87 name,
  cnn87 name NOT NULL,
  ca87 name[],
  cann87 name[] NOT NULL,

  cn88 numrange,
  cnn88 numrange NOT NULL,
  ca88 numrange[],
  cann88 numrange[] NOT NULL,

  cn89 citext,
  cnn89 citext NOT NULL,
  ca89 citext[],
  cann89 citext[] NOT NULL,

  cn90 hstore,
  cnn90 hstore NOT NULL,
  ca90 hstore[],
  cann90 hstore[] NOT NULL,

  cn91 ltree,
  cnn91 ltree NOT NULL,
  ca91 ltree[],
  cann91 ltree[] NOT NULL,

  cn92 cube,
  cnn92 cube NOT NULL,
  ca92 cube[],
  cann92 cube[] NOT NULL,

  cn93 cube,
  cnn93 cube NOT NULL,
  ca93 cube[],
  cann93 cube[] NOT NULL,

  cn94 ean13,
  cnn94 ean13 NOT NULL,
  ca94 ean13[],
  cann94 ean13[] NOT NULL,

  cn95 isbn13,
  cnn95 isbn13 NOT NULL,
  ca95 isbn13[],
  cann95 isbn13[] NOT NULL,

  cn96 ismn13,
  cnn96 ismn13 NOT NULL,
  ca96 ismn13[],
  cann96 ismn13[] NOT NULL,

  cn97 issn13,
  cnn97 issn13 NOT NULL,
  ca97 issn13[],
  cann97 issn13[] NOT NULL,

  cn98 isbn13,
  cnn98 isbn13 NOT NULL,
  ca98 isbn13[],
  cann98 isbn13[] NOT NULL,

  cn99 isbn,
  cnn99 isbn NOT NULL,
  ca99 isbn[],
  cann99 isbn[] NOT NULL,

  cn100 ismn,
  cnn100 ismn NOT NULL,
  ca100 ismn[],
  cann100 ismn[] NOT NULL,

  cn101 issn,
  cnn101 issn NOT NULL,
  ca101 issn[],
  cann101 issn[] NOT NULL,

  cn102 upc,
  cnn102 upc NOT NULL,
  ca102 upc[],
  cann102 upc[] NOT NULL
);
