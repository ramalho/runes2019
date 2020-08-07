import sys
import collections
import random
from operator import attrgetter


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


def load(infile):
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
    return monitors, learners, group_keys


def main():
    try:
        infile = sys.argv[1]
    except IndexError:
        infile = 'groups.tsv'

    monitors, learners, group_keys = load(infile)

    number_of_rooms = 3
    while True:
        # assign monitors
        rooms = [[] for _ in range(number_of_rooms)]
        random.shuffle(monitors)
        distribute(rooms, monitors)

        # assign learners
        for group_key, learner_group in learners.items():
            rooms.sort(key=len)
            distribute(rooms, learner_group)
            key = len

        bad_distributions = []
        for group_key in group_keys:
            if not check_distribution(group_key, rooms):
                bad_distributions.append(group_key)
        if bad_distributions:
            print(', '.join(sorted(bad_distributions)),
                'not well distributed. Retrying...')
        else:
            print('Good distribution:')
            break  # done!

    for r, room in enumerate(rooms, 1):
        print('_' * 40, 'Room', r)
        for n, p in enumerate(sorted(room, key=attrgetter('name')), 1):
            print(f'{n} {p.monitor:1s}\t{p.group}\t{p.name}')


if __name__ == '__main__':
    main()
