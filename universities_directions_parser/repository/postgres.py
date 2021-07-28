from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

from metaclasses import Singleton


class Repository(metaclass=Singleton):
    def __init__(self, config):
        engine = create_engine(
            f'postgresql://{config.username}:{config.password}@'
            f'{config.host}:{config.port}/{config.dbname}')
        session = sessionmaker()
        session.configure(bind=engine)

        self._session = session

    def get_session(self):
        return self._session()
