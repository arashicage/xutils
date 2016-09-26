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
  is '������˰-��ѯ��Ҫ��Ϣ��';
comment on column CKTS_CX_GYXX.cx_bh
  is '��ѯ���';
comment on column CKTS_CX_GYXX.cx_bm
  is '��ѯ����';
comment on column CKTS_CX_GYXX.cx_zwm
  is '��ѯ��������';
comment on column CKTS_CX_GYXX.mxlj
  is '��ϸ����';
comment on column CKTS_CX_GYXX.wy_bz
  is '���Ψһ��־';
comment on column CKTS_CX_GYXX.cx_sm
  is '��ѯ˵��';
comment on column CKTS_CX_GYXX.cxdc_bz
  is '��ѯ������ʶ����ѯ�򵼳��ı�ʶ�ֶΣ�CΪ��ѯ��DΪ����';
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
  is '������˰-��ѯ��ϸ��Ϣ��';
comment on column CKTS_CX_XXXX.cx_bh
  is '��ѯ���';
comment on column CKTS_CX_XXXX.zd_mc
  is '�ֶ�����';
comment on column CKTS_CX_XXXX.zd_zwmc
  is '�ֶ���������';
comment on column CKTS_CX_XXXX.zdsj_lx
  is '�ֶ���������';
comment on column CKTS_CX_XXXX.dm_mc
  is '��������';
comment on column CKTS_CX_XXXX.jsdml
  is '���ʹ�����';
comment on column CKTS_CX_XXXX.tj_bz
  is '��Ϊ������־';
comment on column CKTS_CX_XXXX.mr_tj
  is 'Ĭ������';
comment on column CKTS_CX_XXXX.bx_tj
  is '��ѡ����';
comment on column CKTS_CX_XXXX.tj_xh
  is '�������';
comment on column CKTS_CX_XXXX.cx_xh
  is '��ѯ���';
comment on column CKTS_CX_XXXX.srzd_cd
  is '������󳤶�';
comment on column CKTS_CX_XXXX.fktj_bz
  is '�ǿ�������ʶ';
comment on column CKTS_CX_XXXX.tjmrz
  is '����Ĭ��ֵ';
comment on column CKTS_CX_XXXX.gltj_zd
  is '���������ֶ�';
comment on column CKTS_CX_XXXX.zdjshql
  is '�Զ�������ȡ��';
comment on column CKTS_CX_XXXX.jg_bz
  is '��Ϊ�����ʶ';
comment on column CKTS_CX_XXXX.mr_jg
  is 'Ĭ�Ͻ���ֶ�';
comment on column CKTS_CX_XXXX.bx_jg
  is '��ѡ����ֶ�';
comment on column CKTS_CX_XXXX.jg_xh
  is '������';
comment on column CKTS_CX_XXXX.mxlj
  is '��ϸ����';
comment on column CKTS_CX_XXXX.gljg_zd
  is '��������ֶ�';
comment on column CKTS_CX_XXXX.zdxskd
  is '�ֶ���ʾ���';
comment on column CKTS_CX_XXXX.dq_fs
  is '�ֶζ��뷽ʽ';
comment on column CKTS_CX_XXXX.yczd_bz
  is '�����ֶα�ʶ';
comment on column CKTS_CX_XXXX.px_bz
  is '��Ϊ�����ʶ';
comment on column CKTS_CX_XXXX.mr_px
  is 'Ĭ�������־';
comment on column CKTS_CX_XXXX.bx_px
  is '��ѡ�����־';
comment on column CKTS_CX_XXXX.px_fs
  is 'Ĭ������ʽ';
comment on column CKTS_CX_XXXX.px_xh
  is '�������';
comment on column CKTS_CX_XXXX.yx_bz
  is '�Ƿ���Ч';
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
