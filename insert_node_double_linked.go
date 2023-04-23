package hackerrank

/*
 * Complete the 'sortedInsert' function below.
 *
 * The function is expected to return an INTEGER_DOUBLY_LINKED_LIST.
 * The function accepts following parameters:
 *  1. INTEGER_DOUBLY_LINKED_LIST llist
 *  2. INTEGER data
 */

/*
 * For your reference:
 *
 * DoublyLinkedListNode {
 *     data int32
 *     next *DoublyLinkedListNode
 *     prev *DoublyLinkedListNode
 * }
 *
 */

func sortedInsert(llist *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
    // Write your code here
    newLinkedList := llist

    for llist != nil {

        // check head
        if llist.prev == nil {
            if llist.data >= data {
                // insert value in the head

                node := &DoublyLinkedListNode{
                    next: llist,
                    prev: nil,
                    data: data,
                }
                llist.prev = node
                newLinkedList = node
                return newLinkedList
            } else {
                llist = llist.next
            }
        }
        // check tail
        if llist.next == nil {
            if llist.data <= data {

                node := &DoublyLinkedListNode{
                    next: nil,
                    prev: llist,
                    data: data,
                }

                llist.next = node
                return newLinkedList
            }
        }

        if llist.prev.data <= data && llist.data >= data {
            node := &DoublyLinkedListNode{
                next: llist,
                prev: llist.prev,
                data: data,
            }
            llist.prev.next = node
            llist.prev = node
            return newLinkedList
        }

        if llist.next.data >= data && llist.data <= data {
            node := &DoublyLinkedListNode{
                next: llist.next,
                prev: llist,
                data: data,
            }
            llist.next.prev = node
            llist.next = node
            return newLinkedList
        }

        llist = llist.next
    }

    return newLinkedList
}

