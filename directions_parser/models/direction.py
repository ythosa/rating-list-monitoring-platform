from sqlalchemy import Column, Integer, String, ForeignKey

from models.base import Base


class Direction(Base):
    __tablename__ = "directions"

    id = Column(Integer, primary_key=True)
    name = Column(String, nullable=False)
    url = Column(String, nullable=False)
    university_id = Column(Integer, ForeignKey('universities.id'), primary_key=True)
