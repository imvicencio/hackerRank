package hackerrank

/*
 * Complete the 'reverse' function below.
 *
 * The function is expected to return an INTEGER_SINGLY_LINKED_LIST.
 * The function accepts INTEGER_SINGLY_LINKED_LIST llist as parameter.
 */

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */

 type  SinglyLinkedListNode struct {
     data int32
      next *SinglyLinkedListNode
 }

func reverse(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
    
       // create a new link list
    var newLinkedList *SinglyLinkedListNode

    // order the link list in reverse order
    for llist != nil {
        newLinkedList = &SinglyLinkedListNode{
            data: llist.data,
            next: newLinkedList,
        }
        llist = llist.next
    }

    // return the new link list
    return newLinkedList
}

