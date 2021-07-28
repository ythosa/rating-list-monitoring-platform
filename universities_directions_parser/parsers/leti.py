from typing import List

import requests
from bs4 import BeautifulSoup

from dto import Direction


def parse(directions_url: str) -> List[Direction]:
    r = requests.get(directions_url).content.decode('utf8')
    soup = BeautifulSoup(r, 'html.parser')

    url_base_path = 'https://etu.ru/'
    directions: List[Direction] = []
    for tr in soup.find_all('tr')[2:]:
        try:
            code = tr.contents[1].string
            name_fields = tr.contents[3].contents
            title = name_fields[0]

            if len(name_fields) > 2:
                title_substring = name_fields[2].get_text()
                name = f'{code} {title} {title_substring}'
            else:
                name = f'{code} {title}'

            url = url_base_path + tr.contents[5].contents[0].get('href')

            directions.append(Direction(name=name, url=url))
        except AttributeError:
            continue
        except IndexError:
            break

    return directions
