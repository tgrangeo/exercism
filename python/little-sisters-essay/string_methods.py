"""Functions to help edit essay homework using string manipulation."""


def capitalize_title(title):
    splitted = title.split(' ')
    i = 0
    while i < len(splitted):
        splitted[i] = splitted[i].capitalize()
        i = i+1
    return ' '.join(splitted)


def check_sentence_ending(sentence):
    return sentence.endswith(".")


def clean_up_spacing(sentence):
    return sentence.strip()


def replace_word_choice(sentence, old_word, new_word):
    return sentence.replace(old_word,new_word)
