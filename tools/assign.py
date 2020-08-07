import sys
import collections
import random
from pprint import pprint
from operator import attrgetter
from functools import partial


Person = collections.namedtuple('Person', 'monitor group name')


def distribute(rooms, people):
    people = list(people)
    random.shuffle(people)
    destination = 0
    while people:
        rooms[destination].append(people.pop())
        destination = (destination + 1) % len(rooms)


def count_members(room, group):
    g = 0
    for p in room:
        if p.group == group:
            g += 1
    return len(room), g


def check_distribution(group, rooms):
    counts = [count_members(r, group)[1] for r in rooms]
    return max(counts) - min(counts) <= 1


def main():
    try:
        infile = sys.argv[1]
    except IndexError:
        infile = 'groups.tsv'
    monitors = []
    learners = collections.defaultdict(list)
    group_keys = set()
    with open(infile) as fp:
        for line in fp:
            p = Person(*line.rstrip().split('\t'))
            group_keys.add(p.group)
            if p.monitor == 'y':
                monitors.append(p)
            else:
                learners[p.group].append(p)

    number_of_rooms = 3
    check = False
    while not check:
        # assign monitors
        rooms = [[] for _ in range(number_of_rooms)]
        random.shuffle(monitors)
        distribute(rooms, monitors)

        # assign learners
        for group_key, learner_group in learners.items():
            rooms.sort(key=len)
            distribute(rooms, learner_group)
            key = len

        checks = []
        for group_key in group_keys:
            dist = check_distribution(group_key, rooms)
            checks.append(dist)
            if not dist:
                print(group_key, end=' ')
        check = all(checks)
        if not check:
            print('not well distributed. Retrying...')


    for r, room in enumerate(rooms, 1):
        print('_' * 40, 'Room', r)
        for n, p in enumerate(sorted(room, key=attrgetter('name')), 1):
            print(f'{n} {p.monitor:1s}\t{p.group}\t{p.name}')


if __name__ == '__main__':
    main()
