from sqlalchemy.orm import declarative_base
from sqlalchemy import table, Column, Integer, String, Float
import sqlalchemy as db
from sqlalchemy.orm import sessionmaker
import uuid
import csv

engine = db.create_engine("mysql+mysqlconnector://root:aslon2001@localhost/red30", echo = True)

Base = declarative_base()

class Sales(Base):
    __tablename__ = 'Sales'
    __table_args__ = {"schema":'red30'}

    order_num = Column(Integer, primary_key=True)
    order_type = Column(String(100))
    cust_name = Column(String(100))
    cust_state = Column(String(100))
    prod_category = Column(String(100))
    prod_number = Column(String(100))
    prod_name = Column(String(100))
    quantity = Column(Integer)
    price = Column(Float)
    discount = Column(Float)
    order_total = Column(Float)

    def __str__(self):
        return f'{self.order_num}'

#Base.metadata.create_all(engine)
Session = sessionmaker()
Session.configure(bind=engine)
session = Session()
"""with open('red30.csv', newline='') as file:
    reader = csv.DictReader(file)
    for row in reader:
        sales = Sales(order_num = int(row['order_num']), order_type = row['order_type'],
         cust_name = row['cust_name'], cust_state = row['cust_state'], 
         prod_category = row['prod_category'], prod_number = row['prod_number'],
          prod_name = row['prod_name'], quantity = int(row['quantity']), 
          price = float(row['price']), discount = float(row['discount']), 
          order_total = float(row['order_total']))
        print(sales)
        session.add(sales)"""


#session.commit()

big_sale = [0,0]
for sale in session.query(Sales).order_by(Sales.price):
    if sale.order_total > big_sale[0]:
        big_sale[0] = sale.order_total
        big_sale[1] = sale
    
print(big_sale[1].cust_name)
print(big_sale[0])





