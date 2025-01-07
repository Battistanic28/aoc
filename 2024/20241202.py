# run `python3 input/day_2/script.py from root to generate formatted input`
from input.day_2.output import array_2d

# ~~~ Test Cases ~~~
edge_cases = [
    [7, 10, 8, 10, 11],
    [29, 28, 27, 25, 26, 25, 22, 20]
] # all should be True

mock_reports = [    
    [7,6,4,2,1], # pass
    [1,2,7,8,9], # fail
    [9,7,6,2,1], # fail
    [1,3,2,4,5], # fail
    [8,6,4,4,1], # fail
    [1,3,6,7,9], # pass
]

def get_direction(input1, input2):
    if input1 < input2:
        return 1
    if input1 > input2:
        return -1
    return 0

def is_safe(report):
        safety_threshold = 3
        initial_direction = get_direction(report[0], report[1])

        for i in range(len(report)-1):
            input1 = report[i]
            input2 = report[i+1]
            current_direction = get_direction(input1, input2)
            variance = abs(input1 - input2)
         
            if initial_direction != current_direction:
                return False
            if variance < 1 or variance > safety_threshold:
                return False 
            
        return True


def analyze_reports(reports):
    unsafe_reports = []
    safe_report_count = 0
    for report in reports:
        report_is_safe = is_safe(report)
        if report_is_safe == True:
            safe_report_count += 1
        else:
            unsafe_reports.append(report)

    return safe_report_count, unsafe_reports

# ~~~ Solve ~~~ 
# array_2d_int = [[int(item) for item in row] for row in array_2d]
# print(analyze_reports(array_2d_int))
# print(analyze_reports(array_2d_int))


# Part 2 
array_2d_int = [[int(item) for item in row] for row in array_2d]

safe_report_count, unsafe_reports = analyze_reports(array_2d_int)

# Check by 
def problem_dampener(unsafe_reports):
    safe_reports = 0
    for report in unsafe_reports:
        for i in range(len(report)):
            dampened_report = report[:i] + report[i + 1:]
            report_is_safe = is_safe(dampened_report)
            if report_is_safe:
                safe_reports +=1
                break
    return safe_reports

print(problem_dampener(unsafe_reports) + safe_report_count)