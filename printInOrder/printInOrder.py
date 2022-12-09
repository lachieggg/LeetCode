#!/usr/bin/env python3

import threading

class Foo:
    def __init__(self):
        # Defines all locks
        self.firstlock = threading.Lock()
        self.secondlock = threading.Lock()
        self.thirdlock = threading.Lock()
        
        # Acquire necessary starting locks
        self.secondlock.acquire()
        self.thirdlock.acquire()

    def first(self, printFirst: 'Callable[[], None]') -> None:
        self.firstlock.acquire()
        printFirst()
        self.secondlock.release()

    def second(self, printSecond: 'Callable[[], None]') -> None:
        self.secondlock.acquire()
        printSecond()
        self.thirdlock.release()

    def third(self, printThird: 'Callable[[], None]') -> None:
        self.thirdlock.acquire()
        printThird()

def printFirst(): print("first")
def printSecond(): print("second")
def printThird(): print("third")

if(__name__ == "__main__"):
    f = Foo()
    t1 = threading.Thread(f.first(printFirst))
    t2 = threading.Thread(f.second(printSecond))
    t3 = threading.Thread(f.third(printThird))

    t3.start()
    t2.start()
    t1.start()
