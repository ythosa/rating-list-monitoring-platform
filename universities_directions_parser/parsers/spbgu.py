from typing import List

import requests
from bs4 import BeautifulSoup

from dto import Direction


def parse(directions_url: str) -> List[Direction]:
    without_budget = (
        'СВ.5163.2021 Биология: биоинженерные технологии', 'СВ.5164.2021 Прикладные компьютерные технологии',
        'СВ.5188.2021 Международный менеджмент (с изучением европейских или восточных языков)',
        'СВ.5109.2021 Россиеведение', 'СВ.5193.2021 Иностранный язык (украинский и английский язык)',
        'СМ.5125.2021 Артист драматического театра и кино', 'СВ.5119.2021 Академическое пение',
        'СВ.5140.2021 Инструментальное исполнительство на органе, клавесине, карильоне',
        'СВ.5116.2021 Инструментальное исполнительство на скрипке', 'СМ.5124.2021 Станковая живопись',
    )

    r = requests.get(directions_url).content.decode('utf8')
    soup = BeautifulSoup(r, 'html.parser')

    names = [t.string for t in soup.find_all('b', {'style': 'font-size:12pt;'})]
    direction_names = []
    for name in names:
        if name not in without_budget:
            direction_names.append(name)

    urls = soup.find_all('a')
    direction_url_base_path = 'https://cabinet.spbu.ru/Lists/1k_EntryLists/'
    direction_urls = []
    for url in urls:
        if url.string == 'Госбюджетная':
            direction_urls.append(direction_url_base_path + url.get('href'))

    directions: List[Direction] = []
    for i in range(len(direction_names)):
        directions.append(Direction(name=direction_names[i], url=direction_urls[i]))

    return directions
