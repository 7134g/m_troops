from sqlalchemy import Table, MetaData, Column, INTEGER, String

metadata = MetaData()

SeatPost = Table(
    'salecms_seatpost',
    metadata,
    Column('id', INTEGER, primary_key=True, autoincrement=True),
    Column('start', String(10), unique=True, nullable=False),
    Column('end', String(10), unique=True, nullable=False),
    Column('date', String(20), unique=True, nullable=False),
    Column('company', String(10), unique=True, nullable=False),
    Column('min_price', String(10), unique=True, nullable=False),
    Column('max_price', String(10), unique=True, nullable=False),
    Column('seatCount', INTEGER, unique=True, nullable=False),
    Column('ts', String(20), unique=True, nullable=False),
    Column('up', String(20), unique=True, nullable=True, default=0),)




CencelPost = Table(
    'salecms_cencelpost',
    metadata,
    Column('id', INTEGER, primary_key=True, autoincrement=True),
    Column('start', String(10), unique=True, nullable=False),
    Column('end', String(10), unique=True, nullable=False),
    Column('date', String(20), unique=True, nullable=False),
    Column('company', String(10), unique=True, nullable=False),
    Column('cencelCount', INTEGER, unique=True, nullable=False),
    Column('ts', String(20), unique=True, nullable=False),
    Column('up', String(20), unique=True, nullable=True),)




ViewData = Table(
    'salecms_chartview',
    metadata,
    Column('id', INTEGER, primary_key=True, autoincrement=True),
    Column('start', String(10), unique=True, nullable=False),
    Column('end', String(10), unique=True, nullable=False),
    Column('company', String(10), unique=True, nullable=False),
    Column('date', String(15), unique=True, nullable=False),
    Column('seatCount', INTEGER, unique=True, nullable=False),
    Column('seatPrice', INTEGER, unique=True, nullable=False),)


