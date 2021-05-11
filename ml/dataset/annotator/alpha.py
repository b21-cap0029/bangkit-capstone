from krippendorff import alpha

def load_response(response_name):
    res = []
    with open('{}.response'.format(response_name)) as f:
        res = [int(x.strip()) for x in f.readlines()]
    return res

rnames = ['share-1', 'share-2', 'share-3']
A = list(map(load_response, rnames))
print("krippendorff's alpha: {0:.5g}".format(alpha(A)))
