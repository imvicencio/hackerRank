

    // Complete the insertNodeAtHead function below.

    /*
     * For your reference:
     *
     * SinglyLinkedListNode {
     *     int data;
     *     SinglyLinkedListNode next;
     * }
     *
     */
    static SinglyLinkedListNode insertNodeAtHead(SinglyLinkedListNode llist, int data) {
           SinglyLinkedListNode temp = llist;
            if(temp == null){
                llist = new SinglyLinkedListNode(data);
                return llist;
            }
            
            SinglyLinkedListNode temp2 = new SinglyLinkedListNode(data);
            temp2.next = temp;
            return temp2;


    }

