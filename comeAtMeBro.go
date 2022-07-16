// find first parens, find last parens
// reverse mid so on and so forth
// 

// hyupo.ueid(xyz)aonheu
//           ^   ^
// 
func reverse(str string, openParenthesesIndex int, closeParenthesesIndex int) string {
 
    var res string
    for i := closeParenthesesIndex - 1; i >= openParenthesesIndex + 1; i-- {
        fmt.Println("printing inside reverse, ")
        fmt.Println(string(str[i]))
        res = res + string(str[i])
    }
    return res
}

func solution(inputString string) string {

     var stack []int
     
     for i, c := range inputString {
         if c != ')' && c != '(' {
            continue   
         } else if c == '(' {
            stack = append(stack, i)    
         } else {
               // case of c being )
               lookBack := stack[len(stack)-1]
               stack = stack[:len(stack)-1]
               
               p1 := inputString[:lookBack+1]
               p2 := reverse(inputString, lookBack, i)
               p3 := inputString[i:]
                fmt.Println("p1 is " + p1)
                fmt.Println(p2)
                fmt.Println(p3)
               inputString =  p1 + p2 + p3
         }
     }
     
     fmt.Println(inputString)
     parensRemoved := ""
     for _, c := range inputString {
         if c != '(' && c != ')' {
             parensRemoved = parensRemoved + string(c)
         }
     }
     
     return parensRemoved
}

