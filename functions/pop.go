package functions

// Pop the first element of the slice
//
// Example
//   type knownGreetings []string
//
//   greetings := knownGreetings{"ciao", "hello", "hola"}
//    for greeting := greetings.Pop(); greeting != nil; greeting = greetings.Pop() {
//  	  fmt.Println(*greeting)
//    }
//

func (ss *SliceType) Pop() (popped *ElementType) {

	if len(*ss) == 0 {
		return
	}

	popped = &(*ss)[0]
	*ss = (*ss)[1:]
	return
}
