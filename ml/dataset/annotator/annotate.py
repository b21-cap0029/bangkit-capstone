import os
import sys
import csv


def save_responses(file_name, responses):
    with open('{}.response'.format(file_name), 'w') as f:
        for r in responses:
            f.write(r + '\n')


if __name__ == '__main__':
    file_name = input('csv name: ')

    prompts = []
    if os.path.exists('{}.csv'.format(file_name)):
        with open('{}.csv'.format(file_name)) as cf:
            reader = csv.reader(cf, delimiter=',')
            prompts = [r[1] for r in reader][1:]
    else:
        print('file not found: {}.csv'.format(file_name))
        sys.exit()

    responses = []
    if os.path.exists('{}.response'.format(file_name)):
        rf = open('{}.response'.format(file_name))
        responses = [r.strip() for r in rf.readlines()]
        rf.close()
    else:
        responses = ['-' for x in range(len(prompts))]
        with open('{}.response'.format(file_name), 'w') as rf:
            for r in responses:
                rf.write(r + '\n')

    if '-' in responses:
        last_pos = responses.index('-') % len(prompts)
    else:
        last_pos = 0

    c = ' '
    while len(c) > 0 and c[0].lower() != 'q':
        tweet = prompts[last_pos]
        label = responses[last_pos]

        next_c = input('current position: {}\n'
                       'tweet: {}\n'
                       'label: {}\n'
                       'INPUT (0/1/q/a/d): '.format(last_pos + 1, tweet,
                                                    label)).lower()

        if next_c == '0' or next_c == '1':
            responses[last_pos] = next_c
            save_responses(file_name, responses)
            last_pos = (last_pos + 1) % len(prompts)
        elif next_c == 'a':
            last_pos = (last_pos - 1) % len(prompts)
        elif next_c == 'd':
            last_pos = (last_pos + 1) % len(prompts)

        c = next_c
