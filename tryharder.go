
var currentDays []int
var state2daysSoFar map[string]int
var workHours int
var lengthOfTasks int
var theTasks []int

func recursive(index int) int {
    
    if index >= lengthOfTasks {
        return 0
    }
    
    stateStr := ""
    for _, c := range currentDays {
        stateStr += string(c)
    }
    state := string(index) + stateStr
    
    if val, ok := state2daysSoFar[state]; ok {
        return val
    }
    
    currentDays = append(currentDays, theTasks[index])
    onePlusSolve := 1 + recursive(index + 1)
    currentDays = currentDays[:len(currentDays) - 1]
    
    for i, _ := range currentDays {
        if currentDays[i] + theTasks[index] <= workHours {
            currentDays[i] += theTasks[index]
            
            newVal := recursive(index + 1)
            if newVal < onePlusSolve {
                onePlusSolve = newVal
            }
            
            currentDays[i] -= theTasks[index]
        }
    } 
    
    
    state2daysSoFar[state] = onePlusSolve
    return onePlusSolve
}

func solution(workingHours int, tasks []int) int {
    state2daysSoFar = make(map[string]int)
    lengthOfTasks = len(tasks)
    workHours = workingHours
    theTasks = tasks
    for _, t := range tasks {
        if t > workingHours {
            return -1
        }
    }
    // position, tasks, sessiontime
    return recursive(0)    
}

