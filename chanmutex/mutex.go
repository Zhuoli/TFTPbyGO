package chanmutex

import (

)

type ChanLock chan bool
func NewChanLock() ChanLock { return make(chan bool, 1) }

func (l ChanLock) Lock() {
  l <- true
}
func (l ChanLock) Unlock() {
  <-l
}
func (l ChanLock) TryLock() bool {
  select {
  case l <- true:
    return true
  default:
  }
  return false
}

func (l ChanLock) TryUnLock() bool {
  select {
  case <-l:
    return true
  default:
  }
  return false
}