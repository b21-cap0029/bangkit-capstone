import os
import sys
import csv


def save_responses(response_name, responses):
    with open(response_name, 'w') as f:
        for r in responses:
            f.write(r + '\n')


if __name__ == '__main__':
    csv_name = input('csv name: ') + '.csv'
    annotator_id = input('annotator id (1/2/3): ')[0]
    response_name = '{}-{}.response'.format(csv_name[:-4], annotator_id)

    # load all prompts
    prompts = []
    if os.path.exists(csv_name):
        with open(csv_name) as cf:
            reader = csv.reader(cf, delimiter=',')
            prompts = [r[1] for r in reader][1:]
    else:
        print('file not found: {}'.format(csv_name))
        sys.exit()

    # load saved responses
    responses = []
    if os.path.exists(response_name):
        rf = open(response_name)
        responses = [r.strip() for r in rf.readlines()]
        rf.close()
    else:
        responses = ['-' for x in range(len(prompts))]
        with open(response_name, 'w') as rf:
            for r in responses:
                rf.write(r + '\n')

    # find first empty
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
            save_responses(response_name, responses)
            last_pos = (last_pos + 1) % len(prompts)
        elif next_c == 'a':
            last_pos = (last_pos - 1) % len(prompts)
        elif next_c == 'd':
            last_pos = (last_pos + 1) % len(prompts)

        c = next_c
