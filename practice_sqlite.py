#tasks:
#1:create database users.dn
#2:create table users with id, first_name, second_name, email
#3:make three insertions
#4:query only emails of users




import sqlalchemy as db
#for creating table model
from sqlalchemy import Table, Column, String, MetaData, Integer
#select method 
from sqlalchemy.sql import select
import sqlite3
from uuid import UUID, uuid4


#1
#engine
engine = db.create_engine('sqlite:///users.db')

#2
metadata_obj = MetaData()
#table model
users = Table('users', metadata_obj,
    Column('id',Integer, primary_key = True),
    Column('first_name', String ),
    Column('second_name', String),
    Column('email', String),
)


#metadata_obj.create_all(engine)
#connecting engine
connection = engine.connect()
#3
#making insertions to table users in users.db
"""connection.execute(users.insert(),[
    {"id":1, "first_name":'Aslon', "second_name":'Khamidov',"email":'aslon@gmail.com'},
    {"id":2, "first_name":'ASad', "second_name":'Khamidov',"email":'Asad@gmail.com'},
    {"id":3, "first_name":'Aziza', "second_name":'Khamidov',"email":'Aziza@gmail.com'},
    {"id":4, "first_name":'Saodat', "second_name":'Khamidov',"email":'saodat@gmail.com'},
    {"id":5, "first_name":'kimdir', "second_name":'Khamidov',"email":'kimdir@gmail.com'}
] )"""

#4
selection = select(users)
results = connection.execute(selection)
#looping results 
for id, first_name, second_name, email in results:
    #printing emails of users
    print(email)







