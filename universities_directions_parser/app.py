import time

import models
from config import Config
from parsers import parsing
from repository import Repository


def main():
    try:
        time.sleep(60)

        repository = Repository(Config().db)

        s = repository.get_session()

        for direction in s.query(models.Direction).all():
            s.delete(direction)

        for u in s.query(models.University).all():
            directions = map(
                lambda direction: models.Direction(name=direction.name, url=direction.url, university_id=u.id),
                parsing.get(u.name)(u.directions_page_url),
            )
            for d in directions:
                s.add(d)

        s.commit()
    except:
        return


if __name__ == '__main__':
    main()
