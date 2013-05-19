package main

func iter_through_list(list LinkedList) {
	fmt.Println("Value: ", list.value)

	if list.next != nil {
		iter_through_list(*list.next)
	} else {
		fmt.Println("At end of list.")
	}
}
