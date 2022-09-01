Given a doubly linked list, list nodes also have a child property that can point to a separate doubly linked list. These child nodes can also have one or more child doubly lists of their own, and so on.

Return the list as a single level flattened doubly link list.

class ListNode {
    value: any,
    prev: ListNode,
    next: ListNode,
    child: ListNode
}

Assumption: 1 <= m <= n <= length of the linked list