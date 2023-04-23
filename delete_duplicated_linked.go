package hackerrank

/*
 * Complete the 'removeDuplicates' function below.
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

func removeDuplicates(llist *SinglyLinkedListNode) *SinglyLinkedListNode {
    // Write your code here
       var newLinkedList *SinglyLinkedListNode
    newLinkedList = llist

    for llist != nil {
        if llist.next != nil {
            nextValue := llist.next
            if llist.data == nextValue.data {
                llist.next = nextValue.next
                continue
            }
        }

        llist = llist.next
    }

    return newLinkedList
}

