package hackerrank

/*
 * Complete the 'getNode' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_SINGLY_LINKED_LIST llist
 *  2. INTEGER positionFromTail
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

func getNode(llist *SinglyLinkedListNode, positionFromTail int32) int32 {
    // Write your code here
    mapPosition := make(map[int32]int32)
    var position int32 = 1
    
    for llist != nil {
        mapPosition[position] = llist.data
        llist = llist.next
        position++
    }

    return mapPosition[position - positionFromTail - 1]
}

