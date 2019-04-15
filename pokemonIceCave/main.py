# Write code for a function that finds the minimum number of steps
# required to solve a pokemon ice map

iceCaveMap = [
[0,0,0,0,0,0],
[0,0,0,1,0,0],
[2,0,0,0,0,2]
]

visited = {}

def findMinSteps(iceCave, x, y):
    print "x,y: " + str(x) +","+ str(y)
    if iceCave[y][x] == 1:
        return 0

    answers = []
    leftStep,rightStep,upStep, downStep = -1,-1,-1,-1
    if x-1 > -1:
        leftStep = findNextLocation(iceCave, x, y, "left")
        print "leftStep: " + str(leftStep)
    if x+1 < len(iceCave[y]):
        rightStep = findNextLocation(iceCave, x, y, "right")
        print "rightStep: " + str(rightStep)
    if y-1 > -1:
        upStep = findNextLocation(iceCave, x, y, "up")
        print "upStep: " + str(upStep)
    if y+1 < len(iceCave):
        downStep = findNextLocation(iceCave, x, y, "down")
        print "downStep: " + str(downStep)

    visited[y][x] = True

    if leftStep != -1 and not visited[y][leftStep]:
        answer = findMinSteps(iceCave, leftStep, y)
        if answer >= 0:
            answers.append(answer)
    if rightStep != -1 and not visited[y][rightStep]:
        answer = findMinSteps(iceCave, rightStep, y)
        if answer >= 0:
            answers.append(answer)
    if upStep != -1 and not visited[upStep][x]:
        answer = findMinSteps(iceCave, x, upStep)
        if answer >= 0:
            answers.append(answer)
    if downStep != -1 and not visited[downStep][x]:
        answer  = findMinSteps(iceCave,x, downStep)
        if answer >=0:
            answers.append(answer)
    if len(answers) == 0:
        return -1
    else:
        return 1 + min(answers)


def findNextLocation(iceCave, x, y, direction):
    sliding = True
    while(sliding):
        if direction == "right":
            if x == len(iceCave[y]):
                return x-1
            if iceCave[y][x] == 1:
                return x
            if iceCave[y][x] != 0:
                return x-1
            x = x+1
        if direction == "left":
            if x == -1:
                return 0
            if iceCave[y][x] == 1:
                return x
            if iceCave[y][x] != 0:
                return x+1
            x = x-1
        if direction == "up":
            if y == -1:
                return 0
            if iceCave[y][x] == 1:
                return y
            if iceCave[y][x] != 0:
                return y+1
            y = y-1
        if direction == "down":
            if y == len(iceCave):
                return y-1
            if iceCave[y][x] == 1:
                return y
            if iceCave[y][x] != 0:
                return y-1
            y = y + 1

for y in range(0,len(iceCaveMap)):
    visited[y] = []
    for x in range(0, len(iceCaveMap[y])):
        visited[y].append(False)

print findMinSteps(iceCaveMap, 0,0)
