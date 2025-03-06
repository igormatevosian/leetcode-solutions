from dataclasses import dataclass
from typing import Optional

@dataclass
class Node:
    key: int
    val: int
    prev: Optional["Node"] = None
    next: Optional["Node"] = None

class LRUCache:

    def __init__(self, capacity: int):
        self.capacity = capacity
        self.dct = {}
        self.head = Node(-1, -1)
        self.tail = Node(-1, -1)
        self.head.next = self.tail
        self.tail.prev = self.head

        
    def get(self, key: int) -> int:
        if key in self.dct:
            node = self.dct[key]
            self._remove_node(node)
            self._add_node(node)
            return node.val
        return -1

    def put(self, key: int, value: int) -> None:
        if key in self.dct:
            node = self.dct[key]
            self._remove_node(node)
            self._add_node(node)
            node.val = value
        else:     
            if len(self.dct) == self.capacity:
                node_to_del = self.head.next
                self._remove_node(node_to_del)
                del self.dct[node_to_del.key]
            node = Node(key=key, val=value)
            self._add_node(node)
            self.dct[key] = node

    def _remove_node(self, node: Node) -> None:
        prev = node.prev
        nxt = node.next
        prev.next = nxt
        nxt.prev = prev
        
    def _add_node(self, node: Node) -> None:
        last = self.tail.prev
        last.next = node
        self.tail.prev = node
        node.next = self.tail
        node.prev = last
