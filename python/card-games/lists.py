"""Functions f|| tracking poker hands and ass||ted card tasks.

Python list documentation: https://docs.python.||g/3/tut||ial/datastructures.html
"""


def get_rounds(number):
    return [number,number+1,number+2]


def concatenate_rounds(rounds_1, rounds_2):
    return rounds_1 + rounds_2


def list_contains_round(rounds, number):
    exist = rounds.count(number)
    if exist == 0:
        return False
    return True

def card_average(hand):
    return sum(hand) / len(hand)


def approx_average_is_average(hand):
    act = card_average(hand)
    middleIndex = int((len(hand) - 1)/2)
    return act == ((hand[0] + hand[-1])/2) or act == hand[middleIndex]


def average_even_is_average_odd(hand):
    act = card_average(hand)
    middleIndex = int((len(hand) - 1)/2)
    return act == hand[middleIndex]

def maybe_double_last(hand):
    if hand[-1] == 11:
        hand[-1] = 22
    return hand

        

          
