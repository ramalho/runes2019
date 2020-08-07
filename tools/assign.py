import sys
import collections
import random
from pprint import pprint
from operator import attrgetter

Person = collections.namedtuple('Person', 'monitor group name')

def distribute(rooms, people):
    people = list(people)
    random.shuffle(people)
    destination = 0
    while people:
        rooms[destination].append(people.pop())
        destination = (destination + 1) % len(rooms)

def main():
    try:
        infile = sys.argv[1]
    except IndexError:
        infile = 'groups.tsv'
    monitors = []
    learners = collections.defaultdict(list)
    with open(infile) as fp:
        for line in fp:
            p = Person(*line.rstrip().split('\t'))
            if p.monitor == 'y':
                monitors.append(p)
            else:
                learners[p.group].append(p)

    # assign monitors
    number_of_rooms = 3
    rooms = [[] for _ in range(number_of_rooms)]
    distribute(rooms, monitors)

    # assign learners
    for _, learner_group in learners.items():
        rooms.sort(key=len)
        distribute(rooms, learner_group)

    for r, room in enumerate(rooms, 1):
        print('_' * 40, 'Room', r)
        for n, p in enumerate(sorted(room, key=attrgetter('name')), 1):
            print(f'{n} {p.monitor:1s}\t{p.group}\t{p.name}')


if __name__ == '__main__':
    main()
