import time
from typing import List

import models
from config import Config
from parsers import parsing
from repository import Repository


def main():
    time.sleep(40)  # time for starting other services :D

    repository = Repository(Config().db)
    s = repository.get_session()

    directions: List[models.Direction] = []
    for u in s.query(models.University).all():
        university_directions = map(
            lambda d: models.Direction(name=d.name, url=d.url, university_id=u.id),
            parsing.get(u.name)(u.directions_page_url),
        )
        directions.extend(university_directions)

    stored_directions: List[models.Direction] = s.query(models.Direction).all()
    if len(stored_directions) == len(directions):
        s.commit()
        return

    if len(stored_directions) > 0:
        s.execute(f'TRUNCATE TABLE {models.Direction.__tablename__}')

    for direction in directions:
        s.add(direction)

    s.commit()


if __name__ == '__main__':
    main()
