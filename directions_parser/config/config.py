import logging
import os
from typing import Dict, IO

import yaml

from metaclasses import Singleton


class InvalidConfigFormat(Exception):
    pass


class DB(metaclass=Singleton):
    host: str
    port: str
    username: str
    password: str
    dbname: str
    sslmode: str

    def __init__(self, raw: Dict[str, str]):
        try:
            self.host = raw["host"]
            self.port = raw["port"]
            self.username = raw["username"]
            self.password = raw["password"]
            self.dbname = raw["dbname"]
            self.sslmode = raw["sslmode"]
        except KeyError as e:
            logging.error(e)
            raise InvalidConfigFormat


class Config(metaclass=Singleton):
    _db: DB

    @property
    def db(self):
        return self._db

    def __init__(self):
        config_file_path: str = os.getenv("CONFIG_FILE_PATH", "config/config.yaml")
        config_file: IO = open(config_file_path)
        data: Dict[str, str] = yaml.load(config_file, Loader=yaml.SafeLoader)

        self._db = DB(dict(data["db"]))
