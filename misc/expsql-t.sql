prompt PL/SQL Developer import file
prompt Created on 2015-1-6 by windows
set feedback off
set define off
prompt Dropping CKTS_CX_GYXX...
drop table CKTS_CX_GYXX cascade constraints;
prompt Dropping CKTS_CX_XXXX...
drop table CKTS_CX_XXXX cascade constraints;
prompt Creating CKTS_CX_GYXX...
create table CKTS_CX_GYXX
(
  cx_bh   VARCHAR2(50) not null,
  cx_bm   VARCHAR2(50) not null,
  cx_zwm  VARCHAR2(50) not null,
  mxlj    VARCHAR2(200),
  wy_bz   VARCHAR2(1) default 'N',
  cx_sm   VARCHAR2(200),
  cxdc_bz VARCHAR2(1) default 'C' not null
)
;
comment on table CKTS_CX_GYXX
  is '出口退税-查询概要信息表';
comment on column CKTS_CX_GYXX.cx_bh
  is '查询编号';
comment on column CKTS_CX_GYXX.cx_bm
  is '查询表名';
comment on column CKTS_CX_GYXX.cx_zwm
  is '查询中文名称';
comment on column CKTS_CX_GYXX.mxlj
  is '明细链接';
comment on column CKTS_CX_GYXX.wy_bz
  is '结果唯一标志';
comment on column CKTS_CX_GYXX.cx_sm
  is '查询说明';
comment on column CKTS_CX_GYXX.cxdc_bz
  is '查询导出标识：查询或导出的标识字段，C为查询，D为导出';
alter table CKTS_CX_GYXX
  add constraint PK_CKTS_CX_GYXX primary key (CX_BH);

prompt Creating CKTS_CX_XXXX...
create table CKTS_CX_XXXX
(
  cx_bh   VARCHAR2(21) not null,
  zd_mc   VARCHAR2(200) not null,
  zd_zwmc VARCHAR2(200) not null,
  zdsj_lx VARCHAR2(20) not null,
  dm_mc   VARCHAR2(200),
  jsdml   VARCHAR2(200),
  tj_bz   VARCHAR2(1) not null,
  mr_tj   VARCHAR2(1) default 'N',
  bx_tj   VARCHAR2(1) default 'N',
  tj_xh   NUMBER,
  cx_xh   NUMBER default 1,
  srzd_cd NUMBER,
  fktj_bz VARCHAR2(1) default 'N',
  tjmrz   VARCHAR2(20),
  gltj_zd VARCHAR2(200),
  zdjshql VARCHAR2(200),
  jg_bz   VARCHAR2(1) not null,
  mr_jg   VARCHAR2(1) default 'N',
  bx_jg   VARCHAR2(1) default 'N',
  jg_xh   NUMBER,
  mxlj    VARCHAR2(200),
  gljg_zd VARCHAR2(200),
  zdxskd  VARCHAR2(200) default '120',
  dq_fs   VARCHAR2(20) default 'center',
  yczd_bz VARCHAR2(1) default 'N',
  px_bz   VARCHAR2(1) not null,
  mr_px   VARCHAR2(1) default 'N',
  bx_px   VARCHAR2(1) default 'N',
  px_fs   VARCHAR2(5) default 'ASC',
  px_xh   NUMBER,
  yx_bz   VARCHAR2(1) default 'Y'
)
;
comment on table CKTS_CX_XXXX
  is '出口退税-查询详细信息表';
comment on column CKTS_CX_XXXX.cx_bh
  is '查询编号';
comment on column CKTS_CX_XXXX.zd_mc
  is '字段名称';
comment on column CKTS_CX_XXXX.zd_zwmc
  is '字段中文名称';
comment on column CKTS_CX_XXXX.zdsj_lx
  is '字段数据类型';
comment on column CKTS_CX_XXXX.dm_mc
  is '代码名称';
comment on column CKTS_CX_XXXX.jsdml
  is '解释代码类';
comment on column CKTS_CX_XXXX.tj_bz
  is '做为条件标志';
comment on column CKTS_CX_XXXX.mr_tj
  is '默认条件';
comment on column CKTS_CX_XXXX.bx_tj
  is '必选条件';
comment on column CKTS_CX_XXXX.tj_xh
  is '条件序号';
comment on column CKTS_CX_XXXX.cx_xh
  is '查询序号';
comment on column CKTS_CX_XXXX.srzd_cd
  is '输入最大长度';
comment on column CKTS_CX_XXXX.fktj_bz
  is '非空条件标识';
comment on column CKTS_CX_XXXX.tjmrz
  is '条件默认值';
comment on column CKTS_CX_XXXX.gltj_zd
  is '关联条件字段';
comment on column CKTS_CX_XXXX.zdjshql
  is '自动检索获取类';
comment on column CKTS_CX_XXXX.jg_bz
  is '作为结果标识';
comment on column CKTS_CX_XXXX.mr_jg
  is '默认结果字段';
comment on column CKTS_CX_XXXX.bx_jg
  is '必选结果字段';
comment on column CKTS_CX_XXXX.jg_xh
  is '结果序号';
comment on column CKTS_CX_XXXX.mxlj
  is '明细链接';
comment on column CKTS_CX_XXXX.gljg_zd
  is '关联结果字段';
comment on column CKTS_CX_XXXX.zdxskd
  is '字段显示宽度';
comment on column CKTS_CX_XXXX.dq_fs
  is '字段对齐方式';
comment on column CKTS_CX_XXXX.yczd_bz
  is '隐藏字段标识';
comment on column CKTS_CX_XXXX.px_bz
  is '作为排序标识';
comment on column CKTS_CX_XXXX.mr_px
  is '默认排序标志';
comment on column CKTS_CX_XXXX.bx_px
  is '必选排序标志';
comment on column CKTS_CX_XXXX.px_fs
  is '默认排序方式';
comment on column CKTS_CX_XXXX.px_xh
  is '排序序号';
comment on column CKTS_CX_XXXX.yx_bz
  is '是否有效';
alter table CKTS_CX_XXXX
  add constraint PK_CKTS_CX_XXXX primary key (CX_BH, ZD_MC);

prompt Disabling triggers for CKTS_CX_GYXX...
alter table CKTS_CX_GYXX disable all triggers;
prompt Disabling triggers for CKTS_CX_XXXX...
alter table CKTS_CX_XXXX disable all triggers;
prompt Loading CKTS_CX_GYXX...
prompt Table is empty
prompt Loading CKTS_CX_XXXX...
prompt Table is empty
prompt Enabling triggers for CKTS_CX_GYXX...
alter table CKTS_CX_GYXX enable all triggers;
prompt Enabling triggers for CKTS_CX_XXXX...
alter table CKTS_CX_XXXX enable all triggers;
set feedback on
set define on
prompt Done.
