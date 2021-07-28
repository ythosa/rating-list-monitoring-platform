from sqlalchemy import Column, Integer, String
from sqlalchemy.orm import relationship

from models.base import Base


class University(Base):
    __tablename__ = "universities"

    id = Column(Integer, primary_key=True)
    name = Column(String, nullable=False)
    directions_page_url = Column(String, nullable=False)
    children = relationship('Direction')
