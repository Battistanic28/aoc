from input.day_4.output import array_2d
# array_2d = [
#     ['M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'], 
#     ['M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'], 
#     ['A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'], 
#     ['M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'], 
#     ['X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'], 
#     ['X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'], 
#     ['S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'], 
#     ['S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'],
#     ['M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'],
#     ['M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X']
# ]

# Below is probably the worst code I've ever written
rows = len(array_2d)
cols = len(array_2d[0])

directions = [[0,1], [0,-1], [-1,0], [-1,-1], [-1,1], [1,-1], [1,0], [1,1]]
starting_locations = []
vectors = []

count = 0

for i in range(rows):
    for j in range(cols):
        if array_2d[i][j] == 'X':
            starting_locations.append([i,j])

for coords in starting_locations:
    for dir in directions:
        row = coords[0] + dir[0]
        col = coords[1] + dir[1]
        if 0 <= row <= rows-1 and 0 <= col <= cols-1:
            if array_2d[row][col] == 'M':
                vectors.append([[row,col], dir])

for vector in vectors:
    coords, dir = vector[0], vector[1]
    row, col = coords[0], coords[1]

    next_row, next_col = row+dir[0], col+dir[1]
    next_next_row, next_next_col = row+dir[0]+dir[0], col+dir[1]+dir[1]

    if 0 <= next_row <= rows-1 and 0 <= next_col <= cols-1:
        next_letter = array_2d[next_row][next_col]
        if next_letter == 'A':
            if 0 <= next_next_row <= rows-1 and 0 <= next_next_col <= cols-1:
                next_next_letter = array_2d[next_next_row][next_next_col]
                if next_next_letter == 'S':
                    count += 1

# print(count)

# Part 2
def get_starting_letters(grid): # A
    rows = len(grid)
    cols = len(grid[0])
    starting_locations = []
    for i in range(rows):
        for j in range(cols):
            if array_2d[i][j] == 'A':
                starting_locations.append([i,j])

    return starting_locations

def check_for_mas(center_letter_coords, neighbor_coords):
    letters = ['M','S']
    row = center_letter_coords[0]
    col = center_letter_coords[1]

    for coords in neighbor_coords:
        neighbor_row = row + coords[0]
        neighbor_col = col + coords[1]

        if 0 <= neighbor_row <= len(array_2d)-1 and 0 <= neighbor_col <= len(array_2d[0])-1:
            neighbor_letter = array_2d[neighbor_row][neighbor_col]

            if neighbor_letter in letters:
                letters.remove(neighbor_letter)

            if len(letters) == 0:
                return True
        
    return False

def search_x_pattern(starting_locations):
    count = 0
    direction_positive = [[-1,1],[1,-1]] # /
    direction_negative = [[-1,-1],[1,1]] # \
    
    for location in starting_locations:
        check_one = check_for_mas(location, direction_positive)
        check_two = check_for_mas(location, direction_negative)
        if check_one and check_two:
            count += 1

    return count

starting_locations = get_starting_letters(array_2d)
print(search_x_pattern(starting_locations))