"""Functions for creating, transforming, and adding prefixes to strings."""


def add_prefix_un(word):
    return 'un' + word


def make_word_groups(vocab_words):
    prefix = vocab_words[0]
    x = 1
    while x < len(vocab_words):
        vocab_words[x] = prefix + vocab_words[x]
        x = x + 1
    return ' :: '.join(vocab_words)


def remove_suffix_ness(word):
    trimmed = word[0:-4]
    if (trimmed[-1] == 'i'):
        trimmed = trimmed[:-1] + "y"
    return trimmed

def adjective_to_verb(sentence, index):
    arr = sentence.split(' ')
    if arr[index][-1] == '.':
        arr[index] = arr[index][:-1]
    return arr[index] + "en"
