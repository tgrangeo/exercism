"""Functions for organizing and calculating student exam scores."""


def round_scores(student_scores):
    return [round(i) for i in student_scores]


def count_failed_students(student_scores):
    bad_student = [i for i in student_scores if i <= 40]
    return len(bad_student)


def above_threshold(student_scores, threshold):
     return [i for i in student_scores if i >= threshold]


def letter_grades(highest):
    step = round((highest - 41) /4)
    return [41,41 + step, 41 + step*2,41 + step*3]
    


def student_ranking(student_scores, student_names):
    res = []
    i = 0
    while i < len(student_names):
        res.append(str(i +1)+ ". " + student_names[i]+ ": "+ str(student_scores[i]))
        i = i +1
    return res


def perfect_score(student_info):
    for i in student_info:
        if i[1] == 100:
            return i
    return []
